package ConfigInteractor

//import (
//	"fmt"
//	"github.com/pkg/errors"
//	"testing"
//)
//
//type ConfigRepositoryMock struct {
//	String string
//	Bool   bool
//}
//
//func (m *ConfigRepositoryMock) Vars() (Vars []string) { return }
//
//func (m *ConfigRepositoryMock) GetString(key string) (value string) {
//	return m.String
//}
//
//func (m *ConfigRepositoryMock) GetBool(key string) (value bool, filled bool, err error) {
//	if m.Bool {
//		return m.Bool, m.Bool, nil
//	}
//	return m.Bool, m.Bool, errors.New("false")
//}
//
//func (m *ConfigRepositoryMock) GetInt64(key string) (value int64, filled bool, err error)     { return }
//func (m *ConfigRepositoryMock) GetUint64(key string) (value uint64, filled bool, err error)   { return }
//func (m *ConfigRepositoryMock) GetFloat64(key string) (value float64, filled bool, err error) { return }
//func (m *ConfigRepositoryMock) SetString(key, value string) (err error)                       { return }
//func (m *ConfigRepositoryMock) SetBool(key string, value bool) (err error)                    { return }
//func (m *ConfigRepositoryMock) SetInt64(key string, value int64) (err error)                  { return }
//func (m *ConfigRepositoryMock) SetUint64(key string, value uint64) (err error)                { return }
//func (m *ConfigRepositoryMock) SetFloat64(key string, value float64) (err error)              { return }
//func (m *ConfigRepositoryMock) ClearAllVars()                                                 { return }
//
//type LoggerMock struct {
//	message string
//}
//
//func (lm *LoggerMock) Debug(message string) {
//	lm.message = message
//}
//
//func (lm *LoggerMock) Debugf(format string, a ...interface{}) {
//	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
//}
//
//func (lm *LoggerMock) Info(message string) {
//	lm.message = message
//}
//
//func (lm *LoggerMock) Infof(format string, a ...interface{}) {
//	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
//}
//
//func (lm *LoggerMock) Warn(message string) {
//	lm.message = message
//}
//
//func (lm *LoggerMock) Warnf(format string, a ...interface{}) {
//	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
//}
//
//func (lm *LoggerMock) Error(message string) {
//	lm.message = message
//}
//
//func (lm *LoggerMock) Errorf(format string, a ...interface{}) {
//	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
//}
//
//func (lm *LoggerMock) Fatal(message string) {
//	lm.message = message
//}
//
//func (lm *LoggerMock) Fatalf(format string, a ...interface{}) {
//	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
//}
//
//var (
//	configRepositoryMock = &ConfigRepositoryMock{}
//	ConfigInteractor     *ConfigInteractor
//	loggerMock           = &LoggerMock{message: checkLoggerInited}
//)
//
//const (
//	checkLoggerInited = "logger inited"
//)
//
//func TestNewConfigInteractor(t *testing.T) {
//	ConfigInteractor = New(configRepositoryMock, "prefix_", loggerMock)
//
//	if ConfigInteractor.repo != configRepositoryMock {
//		t.Error("repository is not match")
//		return
//	}
//
//	if ConfigInteractor.prefix != "prefix_" {
//		t.Error("prefix is not match")
//		return
//	}
//
//	if ConfigInteractor.log != loggerMock {
//		t.Error("logger is not match")
//		return
//	}
//
//	if ConfigInteractor.keyLogFileName != "prefix_"+"LOG_FILENAME" {
//		t.Error("keyLogFileName is not match")
//		return
//	}
//
//	if ConfigInteractor.keyElkLogFileName != "prefix_"+"ELK_LOG_FILENAME" {
//		t.Error("keyElkLogFileName is not match")
//		return
//	}
//
//	if ConfigInteractor.keyDebugFileName != "prefix_"+"DEBUG_FILENAME" {
//		t.Error("keyDebugFileName is not match")
//		return
//	}
//
//	if ConfigInteractor.keyDebugMode != "prefix_"+"DEBUG_MODE" {
//		t.Error("keyDebugMode is not match")
//		return
//	}
//
//	if ConfigInteractor.keyDbDriverName != "prefix_"+"DB_DRIVER_NAME" {
//		t.Error("keyDbDriverName is not match")
//		return
//	}
//
//	if ConfigInteractor.keyDbSourceName != "prefix_"+"DB_SOURCE_NAME" {
//		t.Error("keyDbSourceName is not match")
//		return
//	}
//
//	if ConfigInteractor.keyWebListen != "prefix_"+"WEB_LISTEN" {
//		t.Error("keyWebListen is not match")
//		return
//	}
//}
//
//func TestConfigInteractor_readLogFileName(t *testing.T) {
//	configRepositoryMock.String = "abrakadabra"
//	if ConfigInteractor.readLogFileName() != "abrakadabra" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readElkLogFileName(t *testing.T) {
//	configRepositoryMock.String = "abrakadabra"
//	if ConfigInteractor.readElkLogFileName() != "abrakadabra" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readDebugFileName(t *testing.T) {
//	configRepositoryMock.String = "abrakadabra"
//	if ConfigInteractor.readDebugFileName() != "abrakadabra" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readDebugMode_1(t *testing.T) {
//	configRepositoryMock.Bool = true
//	if ConfigInteractor.readDebugMode() != true {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readDebugMode_2(t *testing.T) {
//	configRepositoryMock.Bool = false
//	if ConfigInteractor.readDebugMode() != false {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readDbDriverName_1(t *testing.T) {
//	configRepositoryMock.String = "abrakadabra"
//	result, _ := ConfigInteractor.readDbDriverName()
//
//	if result != "abrakadabra" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readDbDriverName_2(t *testing.T) {
//	configRepositoryMock.String = ""
//	_, err := ConfigInteractor.readDbDriverName()
//
//	if err == nil {
//		t.Error("should be error, but not reached")
//	}
//}
//
//func TestConfigInteractor_readDbSourceName_1(t *testing.T) {
//	configRepositoryMock.String = "abrakadabra"
//	result, _ := ConfigInteractor.readDbSourceName()
//
//	if result != "abrakadabra" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readDbSourceName_2(t *testing.T) {
//	configRepositoryMock.String = ""
//	_, err := ConfigInteractor.readDbSourceName()
//
//	if err == nil {
//		t.Error("should be error, but not reached")
//	}
//}
//
//func TestConfigInteractor_readWebListen_1(t *testing.T) {
//	configRepositoryMock.String = "abrakadabra"
//	if ConfigInteractor.readWebListen() != "abrakadabra" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_readWebListen_2(t *testing.T) {
//	configRepositoryMock.String = ""
//	if ConfigInteractor.readWebListen() == "" {
//		t.Error("value should be changed, but not reached")
//	}
//}
//
//func TestConfigInteractor_ReadConf(t *testing.T) {
//	configRepositoryMock.String = "String"
//	configRepositoryMock.Bool = true
//
//	ConfigInteractor.ReadConf()
//
//	if ConfigInteractor.logFileName != "String" {
//		t.Error("logFileName is not match")
//	}
//
//	if ConfigInteractor.elkLogFileName != "String" {
//		t.Error("elkLogFileName is not match")
//	}
//
//	if ConfigInteractor.debugFileName != "String" {
//		t.Error("debugFileName is not match")
//	}
//
//	if ConfigInteractor.debugMode != true {
//		t.Error("debugMode is not match")
//	}
//
//	if ConfigInteractor.webListen != "String" {
//		t.Error("webListen is not match")
//	}
//
//}
//
//func TestConfigInteractor_GetLogFileName(t *testing.T) {
//	if ConfigInteractor.GetLogFileName() != "String" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_GetElkLogFileName(t *testing.T) {
//	if ConfigInteractor.GetElkLogFileName() != "String" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_GetDebugFileName(t *testing.T) {
//	if ConfigInteractor.GetDebugFileName() != "String" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_GetDebugMode(t *testing.T) {
//	if ConfigInteractor.GetDebugMode() != true {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_GetDbDriverName(t *testing.T) {
//	if ConfigInteractor.GetDbDriverName() != "" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_GetDbSourceName(t *testing.T) {
//	if ConfigInteractor.GetDbSourceName() != "" {
//		t.Error("value is not match")
//	}
//}
//
//func TestConfigInteractor_GetWebListen(t *testing.T) {
//	if ConfigInteractor.GetWebListen() != "String" {
//		t.Error("value is not match")
//	}
//}
