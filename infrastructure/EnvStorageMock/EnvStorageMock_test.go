package EnvStorageMock

import (
	"reflect"
	"testing"
)

const (
	ErrorResultIsNotEqualToExpect   = "result is not equal to expect"
	ErrorShouldBeErrorButNotReached = "should be error, but not reached"
	ErrorSetEnvFailed               = "setenv failed: "
	ErrorUnSetEnvFailed             = "unsetenv failed: "
)

var (
	mockEnv *MockEnv
)

// New
func TestMockEnv_New(t *testing.T) {
	mockEnv = New()
	mockEnvExpect := &MockEnv{
		Storage:           map[string]string{},
		SimulateErrorFlag: false,
	}

	if !reflect.DeepEqual(mockEnv, mockEnvExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}

	// set extra vars for next tests
	mockEnv.Storage["key"] = "value"
	mockEnv.Storage["lorem"] = "ipsum"
	mockEnv.Storage["ilike"] = "gophers"
}

// Getenv
func TestMockEnv_Getenv(t *testing.T) {
	envKey := "key"
	envValueExpect := "value"

	// get var
	resultValue := mockEnv.Getenv(envKey)
	if resultValue != envValueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Setenv success
func TestMockEnv_Setenv(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	err := mockEnv.Setenv(envKey, envValue)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}

	// check var
	if resultValue := mockEnv.Storage[envKey]; resultValue != envValue {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Setenv failed
func TestMockEnv_Setenv2(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	mockEnv.SimulateError()
	err := mockEnv.Setenv(envKey, envValue)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
}

// Unset failed
func TestMockEnv_Unsetenv(t *testing.T) {
	envKey := "foo"

	// unset var
	mockEnv.SimulateError()
	err := mockEnv.Unsetenv(envKey)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
}

// Unset success
func TestMockEnv_Unsetenv2(t *testing.T) {
	envKey := "foo"

	// unset var
	err := mockEnv.Unsetenv(envKey)
	if err != nil {
		t.Error(ErrorUnSetEnvFailed, err.Error())
		return
	}

	// check var
	if _, ok := mockEnv.Storage[envKey]; ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// ExpandEnv exist
func TestMockEnv_ExpandEnv(t *testing.T) {
	sIn := "${ilike} are my favorite animals"
	sOutExpect := "gophers are my favorite animals"

	sOut := mockEnv.ExpandEnv(sIn)
	if sOut != sOutExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// ExpandEnv not exist
func TestMockEnv_ExpandEnv2(t *testing.T) {
	sIn := "I hate ${nobody}"
	sOutExpect := "I hate "

	sOut := mockEnv.ExpandEnv(sIn)
	if sOut != sOutExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// LookupEnv exist
func TestMockEnv_LookupEnv(t *testing.T) {
	envKey := "ilike"
	envValueExpect := "gophers"

	envValue, ok := mockEnv.LookupEnv(envKey)
	if !ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
	if envValue != envValueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// LookupEnv not exist
func TestMockEnv_LookupEnv2(t *testing.T) {
	envKey := "ihate"

	_, ok := mockEnv.LookupEnv(envKey)
	if ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Clearenv
func TestMockEnv_Clearenv(t *testing.T) {
	storageExpect := map[string]string{}

	mockEnv.Clearenv()
	if !reflect.DeepEqual(mockEnv.Storage, storageExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Environ
func TestMockEnv_Environ(t *testing.T) {
	mockEnv.Storage["key"] = "value"
	mockEnv.Storage["lorem"] = "ipsum"
	mockEnv.Storage["ilike"] = "gophers"

	varsExpect := []string{
		"key=value",
		"lorem=ipsum",
		"ilike=gophers",
	}

	vars := mockEnv.Environ()
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

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	mockEnv = New()
	mockEnv.Storage["key"] = "value"
	mockEnv.Storage["lorem"] = "ipsum"
	mockEnv.Storage["ilike"] = "gophers"

	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkMockEnv_Getenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.Getenv("key")
	}
}

func BenchmarkMockEnv_Setenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.Setenv("foo", "bar")
	}
}

func BenchmarkMockEnv_Setenv2(b *testing.B) {
	b.ReportAllocs()
	mockEnv.SimulateError()
	for i := 0; i < b.N; i++ {
		mockEnv.Setenv("foo", "bar")
	}
}

func BenchmarkMockEnv_Unsetenv(b *testing.B) {
	b.ReportAllocs()
	mockEnv.SimulateError()
	for i := 0; i < b.N; i++ {
		mockEnv.Unsetenv("foo")
	}
}

func BenchmarkMockEnv_Unsetenv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.Unsetenv("foo")
	}
}

func BenchmarkMockEnv_ExpandEnv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.ExpandEnv("${ilike} are my favorite animals")
	}
}

func BenchmarkMockEnv_ExpandEnv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.ExpandEnv("I hate ${nobody}")
	}
}

func BenchmarkMockEnv_LookupEnv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.LookupEnv("ilike")
	}
}

func BenchmarkMockEnv_LookupEnv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.LookupEnv("ihate")
	}
}

func BenchmarkMockEnv_Environ(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.Environ()
	}
}

func BenchmarkMockEnv_Clearenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mockEnv.Clearenv()
	}
}
