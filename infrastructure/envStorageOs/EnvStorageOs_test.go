package envStorageOs

import (
	"github.com/pkg/errors"
	"os"
	"reflect"
	"testing"
)

const (
	ErrorResultIsNotEqualToExpect   = "result is not equal to expect"
	ErrorShouldBeErrorButNotReached = "should be error, but not reached"
	ErrorSetEnvFailed               = "setenv failed: "
)

var (
	osEnv *OsEnv
)

// New
func TestNew(t *testing.T) {
	osEnv = New()
	osEnvExpect := &OsEnv{}

	if !reflect.DeepEqual(osEnv, osEnvExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}

	// set extra vars for next tests
	if err := os.Setenv("key", "value"); err != nil {
		t.Error(ErrorSetEnvFailed, err.Error())
		return
	}
	if err := os.Setenv("lorem", "ipsum"); err != nil {
		t.Error(ErrorSetEnvFailed, err.Error())
		return
	}
	if err := os.Setenv("ilike", "gophers"); err != nil {
		t.Error(ErrorSetEnvFailed, err.Error())
		return
	}
}

// Getenv
func TestOsEnv_Getenv(t *testing.T) {
	envKey := "key"
	envValueExpect := "value"

	// get var
	resultValue := osEnv.Getenv(envKey)
	if resultValue != envValueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// Setenv success
func TestOsEnv_Setenv(t *testing.T) {
	envKey := "foo"
	envValue := "bar"

	// set var
	err := osEnv.Setenv(envKey, envValue)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}

	// check var
	if resultValue := os.Getenv(envKey); resultValue != envValue {
		t.Error("result is not equal to expect")
		return
	}
}

// Setenv failed
func TestOsEnv_Setenv2(t *testing.T) {
	envKey := ""
	envValue := "bar"

	// set var
	err := osEnv.Setenv(envKey, envValue)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
		return
	}
}

// Unset success
func TestOsEnv_Unsetenv(t *testing.T) {
	envKey := "foo"

	// unset var
	err := osEnv.Unsetenv(envKey)
	if err != nil {
		t.Error(errors.Errorf("unsetenv failed: %s", err.Error()))
		return
	}

	// check var
	if _, ok := os.LookupEnv(envKey); ok {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}
}

// ExpandEnv exist
func TestOsEnv_ExpandEnv(t *testing.T) {
	sIn := "${ilike} are my favorite animals"
	sOutExpect := "gophers are my favorite animals"

	sOut := osEnv.ExpandEnv(sIn)
	if sOut != sOutExpect {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
		return
	}
}

// ExpandEnv not exist
func TestOsEnv_ExpandEnv2(t *testing.T) {
	sIn := "I hate ${nobody}"
	sOutExpect := "I hate "

	sOut := osEnv.ExpandEnv(sIn)
	if sOut != sOutExpect {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
		return
	}
}

// LookupEnv exist
func TestOsEnv_LookupEnv(t *testing.T) {
	envKey := "ilike"
	envValueExpect := "gophers"

	envValue, ok := osEnv.LookupEnv(envKey)
	if !ok {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
		return
	}
	if envValue != envValueExpect {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
		return
	}
}

// LookupEnv not exist
func TestOsEnv_LookupEnv2(t *testing.T) {
	envKey := "ihate"

	_, ok := osEnv.LookupEnv(envKey)
	if ok {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
		return
	}
}

// Clearenv
func TestMockEnv_Clearenv(t *testing.T) {
	environExpecct := []string{}

	osEnv.Clearenv()
	if !reflect.DeepEqual(os.Environ(), environExpecct) {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
		return
	}
}

// Environ
func TestOsEnv_Environ(t *testing.T) {
	if err := os.Setenv("key", "value"); err != nil {
		t.Error(ErrorSetEnvFailed, err.Error())
		return
	}
	if err := os.Setenv("lorem", "ipsum"); err != nil {
		t.Error(ErrorSetEnvFailed, err.Error())
		return
	}
	if err := os.Setenv("ilike", "gophers"); err != nil {
		t.Error(ErrorSetEnvFailed, err.Error())
		return
	}

	varsExpect := []string{
		"key=value",
		"lorem=ipsum",
		"ilike=gophers",
	}

	vars := osEnv.Environ()
	if !sameStringSlice(vars, varsExpect) {
		t.Error(errors.Errorf(ErrorResultIsNotEqualToExpect))
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
	osEnv = New()

	if err := os.Setenv("key", "value"); err != nil {
		b.Error(ErrorSetEnvFailed, err.Error())
		return
	}
	if err := os.Setenv("lorem", "ipsum"); err != nil {
		b.Error(ErrorSetEnvFailed, err.Error())
		return
	}
	if err := os.Setenv("ilike", "gophers"); err != nil {
		b.Error(ErrorSetEnvFailed, err.Error())
		return
	}

	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkOsEnv_Getenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.Getenv("key")
	}
}

func BenchmarkOsEnv_Setenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.Setenv("foo", "bar")
	}
}

func BenchmarkOsEnv_Setenv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.Setenv("", "bar")
	}
}

func BenchmarkOsEnv_Unsetenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.Unsetenv("foo")
	}
}

func BenchmarkOsEnv_ExpandEnv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.ExpandEnv("${ilike} are my favorite animals")
	}
}

func BenchmarkOsEnv_ExpandEnv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.ExpandEnv("I hate ${nobody}")
	}
}

func BenchmarkOsEnv_LookupEnv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.LookupEnv("ilike")
	}
}

func BenchmarkOsEnv_LookupEnv2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.LookupEnv("ihate")
	}
}

func BenchmarkOsEnv_Environ(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.Environ()
	}
}

func BenchmarkOsEnv_Clearenv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		osEnv.Clearenv()
	}
}
