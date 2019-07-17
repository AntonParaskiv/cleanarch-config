package configRepository

import (
	"fmt"
	ConfigInfrastructure "github.com/AntonParaskiv/cleanarch-config/infrastructure/configStorageEnv"
	"github.com/AntonParaskiv/cleanarch-config/infrastructure/envStorageMock"
	"reflect"
	"testing"
)

const (
	ErrorResultIsNotEqualToExpect   = "result is not equal to expect"
	ErrorShouldBeErrorButNotReached = "should be error, but not reached"
	ErrorStorageSetFailed           = "storage set failed: "
	//ErrorUnSetFailed                = "unset failed: "
)

type LoggerMock struct {
	message string
}

func (lm *LoggerMock) Debug(message string) {
	lm.message = message
}

func (lm *LoggerMock) Debugf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Info(message string) {
	lm.message = message
}

func (lm *LoggerMock) Infof(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Warn(message string) {
	lm.message = message
}

func (lm *LoggerMock) Warnf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Error(message string) {
	lm.message = message
}

func (lm *LoggerMock) Errorf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Fatal(message string) {
	lm.message = message
}

func (lm *LoggerMock) Fatalf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

var (
	varStorage       ConfigInfrastructure.EnvStorage
	configStorage    ConfigStorage
	configRepository *ConfigRepository
	loggerMock       *LoggerMock
)

const (
//testKey1   = "key1"
//testValue1 = "string"
//testKey2   = "key2"
//testValue2 = "true"
//testKey3   = "key3"
//testValue3 = "value3"

//envVarName1  = "TEST_VAR_1"
//	envVarName2  = "TEST_VAR_2"
//	envVarName3  = "TEST_VAR_3"
//	envVarName4  = "TEST_VAR_4"
//	envVarName5  = "TEST_VAR_5"
//	envVarName6  = "TEST_VAR_6"
//	envVarName7  = "TEST_VAR_7"
//	envVarName8  = "TEST_VAR_8"
//	envVarName9  = "TEST_VAR_9"
//	envVarName10 = "TEST_VAR_10"
//	envVarName11 = "TEST_VAR_11"
)

func TestConfigRepository_New(t *testing.T) {
	varStorage = envStorageMock.New()
	configStorage = ConfigInfrastructure.New(varStorage)
	loggerMock = &LoggerMock{}

	configRepositoryExpect := &ConfigRepository{
		storage: configStorage,
		Log:     loggerMock,
	}

	// new repository
	configRepository = New(configStorage, loggerMock)
	if !reflect.DeepEqual(configRepository, configRepositoryExpect) {
		t.Error(ErrorResultIsNotEqualToExpect)
		return
	}

	//// set extra vars for next tests
	//if err := varStorage.Setenv(testKey1, testValue1); err != nil {
	//	t.Error(ErrorStorageSetFailed, testKey1)
	//	return
	//}
	//if err := varStorage.Setenv(testKey2, testValue2); err != nil {
	//	t.Error(ErrorStorageSetFailed, testKey2)
	//	return
	//}
	//if err := varStorage.Setenv(testKey3, testValue3); err != nil {
	//	t.Error(ErrorStorageSetFailed, testKey3)
	//	return
	//}

}

// lookup string is present
func TestConfigRepository_LookupString(t *testing.T) {
	keyStorage := "key10"
	valueStorage := "value10"
	valueExpect := valueStorage
	isPresentExpect := true

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	value, isPresent := configRepository.LookupString(keyStorage)
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup string is not present
func TestConfigRepository_LookupString2(t *testing.T) {
	keyStorage := "key20"
	valueExpect := ""
	isPresentExpect := false

	// get & check result
	value, isPresent := configRepository.LookupString(keyStorage)
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup bool is present
func TestConfigRepository_LookupBool(t *testing.T) {
	keyStorage := "key30"
	valueStorage := "true"
	valueExpect := true
	isPresentExpect := true

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	value, isPresent, err := configRepository.LookupBool(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup bool is not present
func TestConfigRepository_LookupBool2(t *testing.T) {
	keyStorage := "key40"
	valueExpect := false
	isPresentExpect := false

	// get & check result
	value, isPresent, err := configRepository.LookupBool(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup bool err parse
func TestConfigRepository_LookupBool3(t *testing.T) {
	keyStorage := "key50"
	valueStorage := "tru"

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	_, _, err := configRepository.LookupBool(keyStorage)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
	}
	return
}

// lookup int64 is present
func TestConfigRepository_LookupInt64(t *testing.T) {
	keyStorage := "key60"
	valueStorage := "-128"
	valueExpect := int64(-128)
	isPresentExpect := true

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	value, isPresent, err := configRepository.LookupInt64(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup int64 is not present
func TestConfigRepository_LookupInt642(t *testing.T) {
	keyStorage := "key70"
	valueExpect := int64(0)
	isPresentExpect := false

	// get & check result
	value, isPresent, err := configRepository.LookupInt64(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup int64 err parse
func TestConfigRepository_LookupInt643(t *testing.T) {
	keyStorage := "key80"
	valueStorage := "value80"

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	_, _, err := configRepository.LookupInt64(keyStorage)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
	}
	return
}

// lookup uint64 is present
func TestConfigRepository_LookupUint64(t *testing.T) {
	keyStorage := "key90"
	valueStorage := "128"
	valueExpect := uint64(128)
	isPresentExpect := true

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	value, isPresent, err := configRepository.LookupUint64(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup uint64 is not present
func TestConfigRepository_LookupUint642(t *testing.T) {
	keyStorage := "key100"
	valueExpect := uint64(0)
	isPresentExpect := false

	// get & check result
	value, isPresent, err := configRepository.LookupUint64(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup uint64 err parse
func TestConfigRepository_LookupUint643(t *testing.T) {
	keyStorage := "key110"
	valueStorage := "value110"

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	_, _, err := configRepository.LookupUint64(keyStorage)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
	}
	return
}

// lookup float64 is present
func TestConfigRepository_LookupFloat64(t *testing.T) {
	keyStorage := "key120"
	valueStorage := "120.12345"
	valueExpect := float64(120.12345)
	isPresentExpect := true

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	value, isPresent, err := configRepository.LookupFloat64(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup float64 is not present
func TestConfigRepository_LookupFloat642(t *testing.T) {
	keyStorage := "key130"
	valueExpect := float64(0)
	isPresentExpect := false

	// get & check result
	value, isPresent, err := configRepository.LookupFloat64(keyStorage)
	if err != nil {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	if isPresent != isPresentExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}

// lookup floar64 err parse
func TestConfigRepository_LookupFloat643(t *testing.T) {
	keyStorage := "key140"
	valueStorage := "value140"

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	_, _, err := configRepository.LookupFloat64(keyStorage)
	if err == nil {
		t.Error(ErrorShouldBeErrorButNotReached)
	}
	return
}

//
//
//

// get string
func TestConfigRepository_GetString(t *testing.T) {
	keyStorage := "key10"
	valueStorage := "value10"
	valueExpect := valueStorage

	// storage set
	if err := configStorage.Set(keyStorage, valueStorage); err != nil {
		t.Error(ErrorStorageSetFailed, err.Error())
		return
	}

	// get & check result
	value := configRepository.GetString(keyStorage)
	if value != valueExpect {
		t.Error(ErrorResultIsNotEqualToExpect)
	}
	return
}
