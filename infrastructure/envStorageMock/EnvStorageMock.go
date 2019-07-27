package envStorageMock

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

const ErrorSimulated = "simulated error"

type MockEnv struct {
	Storage           map[string]string
	SimulateErrorFlag bool
}

func New() (mockEnv *MockEnv) {
	mockEnv = new(MockEnv)
	mockEnv.Storage = make(map[string]string)
	return
}

func (mockEnv *MockEnv) Getenv(key string) (value string) {
	value = mockEnv.Storage[key]
	return
}

func (mockEnv *MockEnv) Setenv(key, value string) (err error) {
	if mockEnv.SimulateErrorFlag {
		mockEnv.SimulateErrorFlag = false
		err = errors.Errorf(ErrorSimulated)
		return
	}

	mockEnv.Storage[key] = value
	return
}

func (mockEnv *MockEnv) Unsetenv(key string) (err error) {
	if mockEnv.SimulateErrorFlag {
		mockEnv.SimulateErrorFlag = false
		err = errors.Errorf(ErrorSimulated)
		return
	}
	delete(mockEnv.Storage, key)
	return
}

func (mockEnv *MockEnv) ExpandEnv(sIn string) (sOut string) {
	sOut = os.Expand(sIn, mockEnv.Getenv)
	return
}

func (mockEnv *MockEnv) LookupEnv(key string) (value string, isPresent bool) {
	value, isPresent = mockEnv.Storage[key]
	return
}

func (mockEnv *MockEnv) Environ() (envStrings []string) {
	envStrings = make([]string, 0, len(mockEnv.Storage))
	for key, value := range mockEnv.Storage {
		envString := fmt.Sprintf("%s=%s", key, value)
		envStrings = append(envStrings, envString)
	}
	return
}

func (mockEnv *MockEnv) Clearenv() {
	mockEnv.Storage = make(map[string]string)
}

func (mockEnv *MockEnv) SetStorageMap(storage map[string]string) *MockEnv {
	mockEnv.Storage = storage
	return mockEnv
}

func (mockEnv *MockEnv) SimulateError() *MockEnv {
	mockEnv.SimulateErrorFlag = true
	return mockEnv
}
