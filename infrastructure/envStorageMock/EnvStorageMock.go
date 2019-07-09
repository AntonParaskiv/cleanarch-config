package envStorageMock

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

const ErrorSimulated = "simulated error"

type MockEnv struct {
	storage       map[string]string
	simulateError bool
}

func New() (mockEnv *MockEnv) {
	mockEnv = new(MockEnv)
	mockEnv.storage = make(map[string]string)
	return
}

func (mockEnv *MockEnv) Getenv(key string) (value string) {
	value = mockEnv.storage[key]
	return
}

func (mockEnv *MockEnv) Setenv(key, value string) (err error) {
	if mockEnv.simulateError {
		err = errors.Errorf(ErrorSimulated)
		return
	}
	mockEnv.storage[key] = value
	return
}

func (mockEnv *MockEnv) Unsetenv(key string) (err error) {
	if mockEnv.simulateError {
		err = errors.Errorf(ErrorSimulated)
		return
	}
	delete(mockEnv.storage, key)
	return
}

func (mockEnv *MockEnv) ExpandEnv(sIn string) (sOut string) {
	sOut = os.Expand(sIn, mockEnv.Getenv)
	return
}

func (mockEnv *MockEnv) LookupEnv(key string) (value string, isPresent bool) {
	value, isPresent = mockEnv.storage[key]
	return
}

func (mockEnv *MockEnv) Environ() (envStrings []string) {
	envStrings = make([]string, 0, len(mockEnv.storage))
	for key, value := range mockEnv.storage {
		envString := fmt.Sprintf("%s=%s", key, value)
		envStrings = append(envStrings, envString)
	}
	return
}

func (mockEnv *MockEnv) Clearenv() {
	mockEnv.storage = make(map[string]string)
}
