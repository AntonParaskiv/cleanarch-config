package configStorageEnv

import (
	"github.com/AntonParaskiv/cleanarch-config/infrastructure/envStorageMock"
	"reflect"
	"testing"
)

const (
	ErrorResultIsNotEqualToExpect   = "result is not equal to expect"
	ErrorShouldBeErrorButNotReached = "should be error, but not reached"
	ErrorSetFailed                  = "set failed: "
	ErrorUnSetFailed                = "unset failed: "
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
func TestConfigStorageEnv_Get(t *testing.T) {
	envKey := "key"
	envValueExpect := "value"

	// get var
	resultValue := configStorageEnv.Get(envKey)
	if resultValue != envValueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Set success
func TestConfigStorageEnv_Set(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	if err := configStorageEnv.Set(envKey, envValue); err != nil {
		t.Error(ErrorSetFailed + err.Error())
		return
	}

	// mock check var
	if envStorage.Storage[envKey] != envValue {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Set failed
func TestConfigStorageEnv_Set2(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	envStorage.SimulateError = true
	if err := configStorageEnv.Set(envKey, envValue); err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
	envStorage.SimulateError = false
}

// Unset failed
func TestConfigStorageEnv_UnSet(t *testing.T) {
	envKey := "foo"

	// unset var
	envStorage.SimulateError = true
	err := configStorageEnv.UnSet(envKey)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
	envStorage.SimulateError = false
}

// Unset success
func TestConfigStorageEnv_UnSet2(t *testing.T) {
	envKey := "foo"

	// unset var
	err := configStorageEnv.UnSet(envKey)
	if err != nil {
		t.Error(ErrorUnSetFailed, err.Error())
		return
	}

	// check var
	if _, ok := envStorage.Storage[envKey]; ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Expand exist
func TestConfigStorageEnv_Expand(t *testing.T) {
	sIn := "${ilike} are my favorite animals"
	sOutExpect := "gophers are my favorite animals"

	sOut := configStorageEnv.Expand(sIn)
	if sOut != sOutExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Expand not exist
func TestConfigStorageEnv_Expand2(t *testing.T) {
	sIn := "I hate ${nobody}"
	sOutExpect := "I hate "

	sOut := configStorageEnv.Expand(sIn)
	if sOut != sOutExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Lookup exist
func TestConfigStorageEnv_Lookup(t *testing.T) {
	envKey := "ilike"
	envValueExpect := "gophers"

	envValue, ok := configStorageEnv.Lookup(envKey)
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
func TestConfigStorageEnv_Lookup2(t *testing.T) {
	envKey := "ihate"

	_, ok := configStorageEnv.Lookup(envKey)
	if ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Clearenv
func TestConfigStorageEnv_ClearAll(t *testing.T) {
	storageExpect := map[string]string{}

	configStorageEnv.ClearAll()
	if !reflect.DeepEqual(envStorage.Storage, storageExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Vars
func TestMockEnv_Vars(t *testing.T) {
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

func BenchmarkConfigStorageEnv_Get(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Get("key")
	}
}

func BenchmarkConfigStorageEnv_Set(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Set("foo", "bar")
	}
}

func BenchmarkConfigStorageEnv_Set2(b *testing.B) {
	b.ReportAllocs()
	envStorage.SimulateError = true
	for i := 0; i < b.N; i++ {
		configStorageEnv.Set("foo", "bar")
	}
	envStorage.SimulateError = false
}

func BenchmarkConfigStorageEnv_UnSet(b *testing.B) {
	b.ReportAllocs()
	envStorage.SimulateError = true
	for i := 0; i < b.N; i++ {
		configStorageEnv.UnSet("foo")
	}
	envStorage.SimulateError = false
}

func BenchmarkConfigStorageEnv_UnSet2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.UnSet("foo")
	}
}

func BenchmarkConfigStorageEnv_Expand(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Expand("${ilike} are my favorite animals")
	}
}

func BenchmarkConfigStorageEnv_Expand2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Expand("I hate ${nobody}")
	}
}

func BenchmarkConfigStorageEnv_Lookup(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Lookup("ilike")
	}
}

func BenchmarkConfigStorageEnv_Lookup2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Lookup("ihate")
	}
}

func BenchmarkConfigStorageEnv_Vars(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.Vars()
	}
}

func BenchmarkConfigStorageEnv_ClearAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		configStorageEnv.ClearAll()
	}
}
