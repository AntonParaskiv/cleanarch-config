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
	testKey1   = "key1"
	testValue1 = "value1"
	testKey2   = "key2"
	testValue2 = "value2"
	testKey3   = "key3"
	testValue3 = "value3"

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

	// set extra vars for next tests
	if err := varStorage.Setenv(testKey1, testValue1); err != nil {
		t.Error(ErrorStorageSetFailed, testKey1)
		return
	}
	if err := varStorage.Setenv(testKey2, testValue2); err != nil {
		t.Error(ErrorStorageSetFailed, testKey2)
		return
	}
	if err := varStorage.Setenv(testKey3, testValue3); err != nil {
		t.Error(ErrorStorageSetFailed, testKey3)
		return
	}

}

//func TestConfigRepository_lookupString(t *testing.T) {
//
//}

//
// String
//
//
//func checkSetString(key string, value string) (err error) {
//	// repository set
//	err = configRepository.SetString(key, value)
//	if err != nil {
//		err = errors.New("SetString failed: " + err.Error())
//		return
//	}
//
//	// mock check
//	mockValue := storage.Getenv(key)
//	if mockValue != value {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetString(key, valueExpect string) (err error) {
//	repoVar := configRepository.GetString(key)
//	if repoVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//	return
//}
//
//// empty string
//func TestConfigRepository_SetString1(t *testing.T) {
//	if err := checkSetString(envVarName1, ""); err != nil {
//		t.Error(err)
//	}
//}
//
//// filled string
//func TestConfigRepository_SetString2(t *testing.T) {
//	if err := checkSetString(envVarName2, "TEST_VALUE_2"); err != nil {
//		t.Error(err)
//	}
//}
//
//// empty string
//func TestConfigRepository_GetString1(t *testing.T) {
//	if err := checkGetString(envVarName1, ""); err != nil {
//		t.Error(err)
//	}
//}
//
//// filled string
//func TestConfigRepository_GetString2(t *testing.T) {
//	if err := checkGetString(envVarName2, "TEST_VALUE_2"); err != nil {
//		t.Error(err)
//	}
//}
//
////
//// Bool
////
//
//func checkSetBool(key string, value bool, valueExpect string) (err error) {
//	// repository set
//	err = configRepository.SetBool(key, value)
//	if err != nil {
//		err = errors.New("SetBool failed: " + err.Error())
//		return
//	}
//
//	// mock check
//	mockVar := storage.Getenv(key)
//	if mockVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetBool(key string, valueExpect bool) (err error) {
//	// repo get
//	repoVar, ok, err := configRepository.GetBool(key)
//	if err != nil {
//		err = errors.New("GetBool failed: " + err.Error())
//		return
//	}
//
//	// check filled
//	if !ok {
//		err = errors.New("GetBool failed: value not filled")
//		return
//	}
//
//	// check err
//	if repoVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetBoolNotFilled(key string) (err error) {
//	// repo get
//	_, ok, err := configRepository.GetBool(key)
//	if err != nil {
//		err = errors.New("GetBool failed: " + err.Error())
//		return
//	}
//
//	// check filled
//	if ok {
//		err = errors.New("GetBool failed: should be not filled, but it filled")
//		return
//	}
//
//	return
//}
//
//func checkErrGetBool(key string) (err error) {
//	_, _, err = configRepository.GetBool(key)
//	if err == nil {
//		err = errors.New("should be error, but not reached")
//		return
//	}
//	err = nil
//	return
//}
//
//// true
//func TestConfigRepository_SetBool1(t *testing.T) {
//	if err := checkSetBool(envVarName3, true, "true"); err != nil {
//		t.Error(err)
//	}
//}
//
//// false
//func TestConfigRepository_SetBool2(t *testing.T) {
//	if err := checkSetBool(envVarName4, false, "false"); err != nil {
//		t.Error(err)
//	}
//}
//
//// true
//func TestConfigRepository_GetBool1(t *testing.T) {
//	if err := checkGetBool(envVarName3, true); err != nil {
//		t.Error(err)
//	}
//}
//
//// false
//func TestConfigRepository_GetBool2(t *testing.T) {
//	if err := checkGetBool(envVarName4, false); err != nil {
//		t.Error(err)
//	}
//}
//
//// broken
//func TestConfigRepository_GetBool3(t *testing.T) {
//	if err := checkErrGetBool(envVarName2); err != nil {
//		t.Error(err)
//	}
//}
//
//// not filled
//func TestConfigRepository_GetBool4(t *testing.T) {
//	if err := checkGetBoolNotFilled(envVarName1); err != nil {
//		t.Error(err)
//	}
//}
//
////
//// Int64
////
//
//func checkSetInt64(key string, value int64, valueExpect string) (err error) {
//	// repository set
//	err = configRepository.SetInt64(key, value)
//	if err != nil {
//		err = errors.New("SetInt64 failed: " + err.Error())
//		return
//	}
//
//	// mock check
//	mockVar := storage.Getenv(key)
//	if mockVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetInt64(key string, valueExpect int64) (err error) {
//	// repo get
//	repoValue, ok, err := configRepository.GetInt64(key)
//	if err != nil {
//		err = errors.New("GetInt64 failed: " + err.Error())
//		return
//	}
//
//	// check filled
//	if !ok {
//		err = errors.New("GetInt64 failed: value not filled")
//		return
//	}
//
//	// check err
//	if repoValue != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetInt64NotFilled(key string) (err error) {
//	// repo get
//	_, ok, err := configRepository.GetInt64(key)
//	if err != nil {
//		err = errors.New("GetInt64 failed: " + err.Error())
//		return
//	}
//
//	// check filled
//	if ok {
//		err = errors.New("GetInt64 failed: should be not filled, but it filled")
//		return
//	}
//
//	return
//}
//func checkErrGetInt64(key string) (err error) {
//	_, _, err = configRepository.GetInt64(key)
//	if err == nil {
//		err = errors.New("should be error, but not reached")
//		return
//	}
//	err = nil
//	return
//}
//
//// min int64
//func TestConfigRepository_SetInt641(t *testing.T) {
//	if err := checkSetInt64(envVarName5, -9223372036854775808, "-9223372036854775808"); err != nil {
//		t.Error(err)
//	}
//}
//
//// max int64
//func TestConfigRepository_SetInt642(t *testing.T) {
//	if err := checkSetInt64(envVarName6, 9223372036854775807, "9223372036854775807"); err != nil {
//		t.Error(err)
//	}
//}
//
//// min int64
//func TestConfigRepository_GetInt641(t *testing.T) {
//	if err := checkGetInt64(envVarName5, -9223372036854775808); err != nil {
//		t.Error(err)
//	}
//}
//
//// max int64
//func TestConfigRepository_GetInt642(t *testing.T) {
//	if err := checkGetInt64(envVarName6, 9223372036854775807); err != nil {
//		t.Error(err)
//	}
//}
//
//// broken
//func TestConfigRepository_GetInt643(t *testing.T) {
//	if err := checkErrGetInt64(envVarName2); err != nil {
//		t.Error(err)
//	}
//}
//
//// not filled
//func TestConfigRepository_GetInt644(t *testing.T) {
//	if err := checkGetInt64NotFilled(envVarName1); err != nil {
//		t.Error(err)
//	}
//}
//
////
//// Uint64
////
//
//func checkSetUint64(key string, value uint64, valueExpect string) (err error) {
//	// repository set
//	err = configRepository.SetUint64(key, value)
//	if err != nil {
//		err = errors.New("SetUint64 failed: " + err.Error())
//		return
//	}
//
//	// mock check
//	mockVar := storage.Getenv(key)
//	if mockVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetUint64(key string, valueExpect uint64) (err error) {
//	// repo get
//	repoVar, ok, err := configRepository.GetUint64(key)
//	if err != nil {
//		err = errors.New("GetUint64 failed: " + err.Error())
//		return
//	}
//
//	// check filled
//	if !ok {
//		err = errors.New("GetUint64 failed: value not filled")
//		return
//	}
//
//	// check err
//	if repoVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetUint64NotFilled(key string) (err error) {
//	// repo get
//	_, ok, err := configRepository.GetUint64(key)
//	if err != nil {
//		err = errors.New("GetUint64 failed: " + err.Error())
//		return
//	}
//
//	// check filled
//	if ok {
//		err = errors.New("GetUint64 failed: should be not filled, but it filled")
//		return
//	}
//
//	return
//}
//
//func checkErrGetUint64(key string) (err error) {
//	_, _, err = configRepository.GetUint64(key)
//	if err == nil {
//		err = errors.New("should be error, but not reached")
//		return
//	}
//	err = nil
//	return
//}
//
//// min uint64
//func TestConfigRepository_SetUint641(t *testing.T) {
//	if err := checkSetUint64(envVarName7, 0, "0"); err != nil {
//		t.Error(err)
//	}
//}
//
//// max uint64
//func TestConfigRepository_SetUint642(t *testing.T) {
//	if err := checkSetUint64(envVarName8, 18446744073709551615, "18446744073709551615"); err != nil {
//		t.Error(err)
//	}
//}
//
//// min int64
//func TestConfigRepository_GetUint641(t *testing.T) {
//	if err := checkGetUint64(envVarName7, 0); err != nil {
//		t.Error(err)
//	}
//}
//
//// max int64
//func TestConfigRepository_GetUint642(t *testing.T) {
//	if err := checkGetUint64(envVarName8, 18446744073709551615); err != nil {
//		t.Error(err)
//	}
//}
//
//// broken
//func TestConfigRepository_GetUint643(t *testing.T) {
//	if err := checkErrGetUint64(envVarName2); err != nil {
//		t.Error(err)
//	}
//}
//
//// not filled
//func TestConfigRepository_GetUint644(t *testing.T) {
//	if err := checkGetUint64NotFilled(envVarName1); err != nil {
//		t.Error(err)
//	}
//}
//
////
//// Float64
////
//
//func checkSetFloat64(key string, value float64, valueExpect string) (err error) {
//	// repository set
//	err = configRepository.SetFloat64(key, value)
//	if err != nil {
//		err = errors.New("SetFloat64 failed: " + err.Error())
//		return
//	}
//
//	// check mock
//	mockVar := storage.Getenv(key)
//	if mockVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//
//	return
//}
//
//func checkGetFloat64(key string, valueExpect float64) (err error) {
//	// repo get
//	repoVar, ok, err := configRepository.GetFloat64(key)
//	if err != nil {
//		err = errors.New("GetFloat64 failed: " + err.Error())
//		return
//	}
//	if !ok {
//		err = errors.New("GetFloat64 failed: value not filled")
//		return
//	}
//	if repoVar != valueExpect {
//		err = errors.New("result is not equal to expect")
//		return
//	}
//	return
//}
//
//func checkGetFloat64NotFilled(key string) (err error) {
//	// repo get
//	_, ok, err := configRepository.GetFloat64(key)
//	if err != nil {
//		err = errors.New("GetFloat64 failed: " + err.Error())
//		return
//	}
//	if ok {
//		err = errors.New("GetFloat64 failed: should be not filled, but it filled")
//		return
//	}
//	return
//}
//
//func checkErrGetFloat64(key string) (err error) {
//	_, _, err = configRepository.GetFloat64(key)
//	if err == nil {
//		err = errors.New("should be error, but not reached")
//		return
//	}
//	err = nil
//
//	return
//}
//
//// zero
//func TestConfigRepository_SetFloat641(t *testing.T) {
//	if err := checkSetFloat64(envVarName9, 0, "0.000000"); err != nil {
//		t.Error(err)
//	}
//}
//
//// pi + round
//func TestConfigRepository_SetFloat642(t *testing.T) {
//	if err := checkSetFloat64(envVarName10, 3.141592653589793, "3.141593"); err != nil {
//		t.Error(err)
//	}
//}
//
//// negative pi + round
//func TestConfigRepository_SetFloat643(t *testing.T) {
//	if err := checkSetFloat64(envVarName11, -3.141592653589793, "-3.141593"); err != nil {
//		t.Error(err)
//	}
//}
//
//// zero
//func TestConfigRepository_GetFloat641(t *testing.T) {
//	if err := checkGetFloat64(envVarName9, 0); err != nil {
//		t.Error(err)
//	}
//}
//
//// pi
//func TestConfigRepository_GetFloat642(t *testing.T) {
//	if err := checkGetFloat64(envVarName10, 3.141593); err != nil {
//		t.Error(err)
//	}
//}
//
//// negative pi
//func TestConfigRepository_GetFloat643(t *testing.T) {
//	if err := checkGetFloat64(envVarName11, -3.141593); err != nil {
//		t.Error(err)
//	}
//}
//
//// broken
//func TestConfigRepository_GetFloat644(t *testing.T) {
//	if err := checkErrGetFloat64(envVarName2); err != nil {
//		t.Error(err)
//	}
//}
//
//// negative pi
//func TestConfigRepository_GetFloat645(t *testing.T) {
//	if err := checkGetFloat64NotFilled(envVarName1); err != nil {
//		t.Error(err)
//	}
//}
//
////
//// ClearAll
////
//
//func TestConfigRepository_ClearAllVars(t *testing.T) {
//	// mock set
//	if err := storage.Setenv(envVarName2, "TEST_VALUE_2"); err != nil {
//		t.Error("storage setenv failed: " + err.Error())
//		return
//	}
//
//	// clear vars
//	configRepository.ClearAll()
//
//	// check mock
//	mockVars := storage.Environ()
//	if len(mockVars) > 0 {
//		t.Error(errors.New("result is not equal to expect"))
//		return
//	}
//}
//
////
//// Vars
////
//
//func TestConfigRepository_Vars(t *testing.T) {
//	// mock set
//	if err := storage.Setenv("TEST_VAR_1", "TEST_VALUE_1"); err != nil {
//		t.Error("storage setenv failed: " + err.Error())
//		return
//	}
//	if err := storage.Setenv("TEST_VAR_2", "TEST_VALUE_2"); err != nil {
//		t.Error("storage setenv failed: " + err.Error())
//		return
//	}
//	if err := storage.Setenv("TEST_VAR_3", "TEST_VALUE_3"); err != nil {
//		t.Error("storage setenv failed: " + err.Error())
//		return
//	}
//	expectedVars := []string{
//		"TEST_VAR_1=TEST_VALUE_1",
//		"TEST_VAR_2=TEST_VALUE_2",
//		"TEST_VAR_3=TEST_VALUE_3",
//	}
//
//	// check repo
//	repoVars := configRepository.Vars()
//
//	if !sameStringSlice(repoVars, expectedVars) {
//		t.Error(errors.New("result is not equal to expect"))
//		return
//	}
//}
//
//func sameStringSlice(x, y []string) bool {
//	if len(x) != len(y) {
//		return false
//	}
//	// create a map of string -> int
//	diff := make(map[string]int, len(x))
//	for _, _x := range x {
//		// 0 value for int is 0, so just increment a counter for the string
//		diff[_x]++
//	}
//	for _, _y := range y {
//		// If the string _y is not in diff bail out early
//		if _, ok := diff[_y]; !ok {
//			return false
//		}
//		diff[_y] -= 1
//		if diff[_y] == 0 {
//			delete(diff, _y)
//		}
//	}
//	if len(diff) == 0 {
//		return true
//	}
//	return false
//}
