package config

import (
	"github.com/adammck/venv"
	"reflect"
	"testing"
)

var (
	storage venv.Env // mock
	config  *ConfigStorageEnv
)

const (
	envVarName1  = "TEST_VAR_1"
	envVarValue1 = "TEST_VALUE_1"
	envVarName2  = "TEST_VAR_2"
	envVarValue2 = "TEST_VALUE_2"
	envVarName3  = "TEST_VAR_3"
	envVarValue3 = "TEST_VALUE_3"
)

func TestNewConfigStorageEnv(t *testing.T) {
	storage = venv.Mock()
	config = NewConfigStorageEnv(storage)

	// mock set vars
	if err := storage.Setenv(envVarName1, envVarValue1); err != nil {
		t.Error("mock set var failed: " + err.Error())
		return
	}
	if err := storage.Setenv(envVarName2, envVarValue2); err != nil {
		t.Error("mock set var failed: " + err.Error())
		return
	}
}

func TestConfigStorageEnv_SetVar(t *testing.T) {
	// set var
	if err := config.SetVar(envVarName3, envVarValue3); err != nil {
		t.Error("setVar failed: " + err.Error())
		return
	}

	// mock check var
	mockVar := storage.Getenv(envVarName3)
	if mockVar != envVarValue3 {
		t.Error("result is not equal to expect")
	}
}

func TestConfigStorageEnv_GetVar(t *testing.T) {
	// get var
	Var := config.GetVar(envVarName3)
	if Var != envVarValue3 {
		t.Error("result is not equal to expect")
	}
}

func TestConfigStorageEnv_Vars(t *testing.T) {
	VarsExpect := []string{
		"TEST_VAR_2=TEST_VALUE_2",
		"TEST_VAR_3=TEST_VALUE_3",
		"TEST_VAR_1=TEST_VALUE_1",
	}

	// get and compare vars
	Vars := config.Vars()
	if !sameStringSlice(Vars, VarsExpect) {
		t.Error("result is not equal to expect")
	}
}

func TestConfigStorageEnv_ClearAllVars(t *testing.T) {
	VarsExpect := []string{}

	// clear vars
	config.ClearAllVars()

	// mock check vars
	mockVars := storage.Environ()
	if !reflect.DeepEqual(mockVars, VarsExpect) {
		t.Error("result is not equal to expect")
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
	storage = venv.Mock()
	config = NewConfigStorageEnv(storage)

	for i := 0; i < b.N; i++ {
		NewConfigStorageEnv(storage)
	}
}

func BenchmarkConfigStorageEnv_SetVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		config.SetVar(envVarName3, envVarValue3)
	}
}

func BenchmarkConfigStorageEnv_GetVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		config.GetVar(envVarName3)
	}
}

func BenchmarkConfigStorageEnv_Vars(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		config.Vars()
	}
}

func BenchmarkConfigStorageEnv_ClearAllVars(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		config.ClearAllVars()
	}
}
