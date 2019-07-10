package configStorageEnv

import (
	"github.com/AntonParaskiv/cleanarch-config/infrastructure/envStorageMock"
	"reflect"
	"testing"
)

const (
	ErrorResultIsNotEqualToExpect   = "result is not equal to expect"
	ErrorShouldBeErrorButNotReached = "should be error, but not reached"
	ErrorSetVarFailed               = "set var failed: "
	ErrorUnSetVarFailed             = "unset var failed: "
)

var (
	envStorage       *envStorageMock.MockEnv
	configStorageEnv *ConfigStorageEnv
)

// New
func TestConfigStorageEnv_New(t *testing.T) {
	envStorage = envStorageMock.New()
	envStorage.Storage["key"] = "value"
	envStorage.Storage["lorem"] = "ipsum"
	envStorage.Storage["ilike"] = "gophers"

	configStorageEnvExpect := &ConfigStorageEnv{
		storage: envStorage,
	}

	configStorageEnv = New(envStorage)
	if !reflect.DeepEqual(configStorageEnv, configStorageEnvExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// get
func TestConfigStorageEnv_GetVar(t *testing.T) {
	envKey := "key"
	envValueExpect := "value"

	// get var
	resultValue := configStorageEnv.GetVar(envKey)
	if resultValue != envValueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Set success
func TestConfigStorageEnv_SetVar(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	if err := configStorageEnv.SetVar(envKey, envValue); err != nil {
		t.Error(ErrorSetVarFailed + err.Error())
		return
	}

	// mock check var
	if envStorage.Storage[envKey] != envValue {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Set failed
func TestConfigStorageEnv_SetVar2(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	envStorage.SimulateError = true
	if err := configStorageEnv.SetVar(envKey, envValue); err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
	envStorage.SimulateError = false
}

// Unset failed
func TestConfigStorageEnv_UnSetVar(t *testing.T) {
	envKey := "foo"

	// unset var
	envStorage.SimulateError = true
	err := configStorageEnv.UnSetVar(envKey)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
	envStorage.SimulateError = false
}

// Unset success
func TestConfigStorageEnv_UnSetVar2(t *testing.T) {
	envKey := "foo"

	// unset var
	err := configStorageEnv.UnSetVar(envKey)
	if err != nil {
		t.Error(ErrorUnSetVarFailed, err.Error())
		return
	}

	// check var
	if _, ok := envStorage.Storage[envKey]; ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Expand exist
func TestConfigStorageEnv_ExpandVar(t *testing.T) {
	sIn := "${ilike} are my favorite animals"
	sOutExpect := "gophers are my favorite animals"

	sOut := configStorageEnv.ExpandVar(sIn)
	if sOut != sOutExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Expand not exist
func TestConfigStorageEnv_ExpandVar2(t *testing.T) {
	sIn := "I hate ${nobody}"
	sOutExpect := "I hate "

	sOut := configStorageEnv.ExpandVar(sIn)
	if sOut != sOutExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Lookup exist
func TestConfigStorageEnv_LookupVar(t *testing.T) {
	envKey := "ilike"
	envValueExpect := "gophers"

	envValue, ok := configStorageEnv.LookupVar(envKey)
	if !ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
	if envValue != envValueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Lookup not exist
func TestConfigStorageEnv_LookupVar2(t *testing.T) {
	envKey := "ihate"

	_, ok := configStorageEnv.LookupVar(envKey)
	if ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Clearenv
func TestConfigStorageEnv_ClearAllVars(t *testing.T) {
	storageExpect := map[string]string{}

	configStorageEnv.ClearAllVars()
	if !reflect.DeepEqual(envStorage.Storage, storageExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Environ
func TestMockEnv_Environ(t *testing.T) {
	envStorage.Storage["key"] = "value"
	envStorage.Storage["lorem"] = "ipsum"
	envStorage.Storage["ilike"] = "gophers"

	varsExpect := []string{
		"key=value",
		"lorem=ipsum",
		"ilike=gophers",
	}

	vars := configStorageEnv.Vars()
	if !sameStringSlice(vars, varsExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

func BenchmarkNewConfigStorageEnv(b *testing.B) {
	b.ReportAllocs()

	envStorage = envStorageMock.New()
	envStorage.Storage["key"] = "value"
	envStorage.Storage["lorem"] = "ipsum"
	envStorage.Storage["ilike"] = "gophers"

	configStorageEnv = New(envStorage)

	for i := 0; i < b.N; i++ {
		New(envStorage)
	}
}

func BenchmarkConfigStorageEnv_GetVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.GetVar("key")
	}
}

func BenchmarkConfigStorageEnv_SetVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.SetVar("foo", "bar")
	}
}

func BenchmarkConfigStorageEnv_SetVar2(b *testing.B) {
	b.ReportAllocs()
	envStorage.SimulateError = true
	for i := 0; i < b.N; i++ {
		configStorageEnv.SetVar("foo", "bar")
	}
	envStorage.SimulateError = false
}

func BenchmarkConfigStorageEnv_UnSetVar(b *testing.B) {
	b.ReportAllocs()
	envStorage.SimulateError = true
	for i := 0; i < b.N; i++ {
		configStorageEnv.UnSetVar("foo")
	}
	envStorage.SimulateError = false
}

func BenchmarkConfigStorageEnv_UnSetVar2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.UnSetVar("foo")
	}
}

func BenchmarkConfigStorageEnv_ExpandVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.ExpandVar("${ilike} are my favorite animals")
	}
}

func BenchmarkConfigStorageEnv_ExpandVar2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.ExpandVar("I hate ${nobody}")
	}
}

func BenchmarkConfigStorageEnv_LookupVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.LookupVar("ilike")
	}
}

func BenchmarkConfigStorageEnv_LookupVar2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.LookupVar("ihate")
	}
}

func BenchmarkConfigStorageEnv_EnvironVars(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Vars()
	}
}

func BenchmarkConfigStorageEnv_ClearAllVars(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.ClearAllVars()
	}
}
