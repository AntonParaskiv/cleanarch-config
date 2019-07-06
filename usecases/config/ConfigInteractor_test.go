package config

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type ConfigRepositoryMock struct {
	String string
	Bool   bool
}

func (m *ConfigRepositoryMock) Vars() (Vars []string) { return }

func (m *ConfigRepositoryMock) GetString(key string) (Value string) {
	return m.String
}

func (m *ConfigRepositoryMock) GetBool(key string) (Value bool, filled bool, err error) {
	if m.Bool {
		return m.Bool, m.Bool, nil
	}
	return m.Bool, m.Bool, errors.New("false")
}

func (m *ConfigRepositoryMock) GetInt64(key string) (Value int64, filled bool, err error)     { return }
func (m *ConfigRepositoryMock) GetUint64(key string) (Value uint64, filled bool, err error)   { return }
func (m *ConfigRepositoryMock) GetFloat64(key string) (Value float64, filled bool, err error) { return }
func (m *ConfigRepositoryMock) SetString(key, Value string) (err error)                       { return }
func (m *ConfigRepositoryMock) SetBool(key string, Value bool) (err error)                    { return }
func (m *ConfigRepositoryMock) SetInt64(key string, Value int64) (err error)                  { return }
func (m *ConfigRepositoryMock) SetUint64(key string, Value uint64) (err error)                { return }
func (m *ConfigRepositoryMock) SetFloat64(key string, Value float64) (err error)              { return }
func (m *ConfigRepositoryMock) ClearAllVars()                                                 { return }

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
	configRepositoryMock = &ConfigRepositoryMock{}
	configInteractor     *ConfigInteractor
	loggerMock           = &LoggerMock{message: checkLoggerInited}
)

const (
	checkLoggerInited = "logger inited"
)

func TestNewConfigInteractor(t *testing.T) {
	configInteractor = NewConfigInteractor(configRepositoryMock, "prefix_", loggerMock)

	if configInteractor.repo != configRepositoryMock {
		t.Error("repository is not match")
		return
	}

	if configInteractor.prefix != "prefix_" {
		t.Error("prefix is not match")
		return
	}

	if configInteractor.Log != loggerMock {
		t.Error("logger is not match")
		return
	}

	if configInteractor.keyLogFileName != "prefix_"+"LOG_FILENAME" {
		t.Error("keyLogFileName is not match")
		return
	}

	if configInteractor.keyElkLogFileName != "prefix_"+"ELK_LOG_FILENAME" {
		t.Error("keyElkLogFileName is not match")
		return
	}

	if configInteractor.keyDebugFileName != "prefix_"+"DEBUG_FILENAME" {
		t.Error("keyDebugFileName is not match")
		return
	}

	if configInteractor.keyDebugMode != "prefix_"+"DEBUG_MODE" {
		t.Error("keyDebugMode is not match")
		return
	}

	if configInteractor.keyDbDriverName != "prefix_"+"DB_DRIVER_NAME" {
		t.Error("keyDbDriverName is not match")
		return
	}

	if configInteractor.keyDbSourceName != "prefix_"+"DB_SOURCE_NAME" {
		t.Error("keyDbSourceName is not match")
		return
	}

	if configInteractor.keyWebListen != "prefix_"+"WEB_LISTEN" {
		t.Error("keyWebListen is not match")
		return
	}
}

func TestConfigInteractor_readLogFileName(t *testing.T) {
	configRepositoryMock.String = "abrakadabra"
	if configInteractor.readLogFileName() != "abrakadabra" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readElkLogFileName(t *testing.T) {
	configRepositoryMock.String = "abrakadabra"
	if configInteractor.readElkLogFileName() != "abrakadabra" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readDebugFileName(t *testing.T) {
	configRepositoryMock.String = "abrakadabra"
	if configInteractor.readDebugFileName() != "abrakadabra" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readDebugMode_1(t *testing.T) {
	configRepositoryMock.Bool = true
	if configInteractor.readDebugMode() != true {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readDebugMode_2(t *testing.T) {
	configRepositoryMock.Bool = false
	if configInteractor.readDebugMode() != false {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readDbDriverName_1(t *testing.T) {
	configRepositoryMock.String = "abrakadabra"
	result, _ := configInteractor.readDbDriverName()

	if result != "abrakadabra" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readDbDriverName_2(t *testing.T) {
	configRepositoryMock.String = ""
	_, err := configInteractor.readDbDriverName()

	if err == nil {
		t.Error("should be error, but not reached")
	}
}

func TestConfigInteractor_readDbSourceName_1(t *testing.T) {
	configRepositoryMock.String = "abrakadabra"
	result, _ := configInteractor.readDbSourceName()

	if result != "abrakadabra" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readDbSourceName_2(t *testing.T) {
	configRepositoryMock.String = ""
	_, err := configInteractor.readDbSourceName()

	if err == nil {
		t.Error("should be error, but not reached")
	}
}

func TestConfigInteractor_readWebListen_1(t *testing.T) {
	configRepositoryMock.String = "abrakadabra"
	if configInteractor.readWebListen() != "abrakadabra" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_readWebListen_2(t *testing.T) {
	configRepositoryMock.String = ""
	if configInteractor.readWebListen() == "" {
		t.Error("value should be changed, but not reached")
	}
}

func TestConfigInteractor_ReadConf(t *testing.T) {
	configRepositoryMock.String = "String"
	configRepositoryMock.Bool = true

	configInteractor.ReadConf()

	if configInteractor.logFileName != "String" {
		t.Error("logFileName is not match")
	}

	if configInteractor.elkLogFileName != "String" {
		t.Error("elkLogFileName is not match")
	}

	if configInteractor.debugFileName != "String" {
		t.Error("debugFileName is not match")
	}

	if configInteractor.debugMode != true {
		t.Error("debugMode is not match")
	}

	if configInteractor.webListen != "String" {
		t.Error("webListen is not match")
	}

}

func TestConfigInteractor_GetLogFileName(t *testing.T) {
	if configInteractor.GetLogFileName() != "String" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_GetElkLogFileName(t *testing.T) {
	if configInteractor.GetElkLogFileName() != "String" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_GetDebugFileName(t *testing.T) {
	if configInteractor.GetDebugFileName() != "String" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_GetDebugMode(t *testing.T) {
	if configInteractor.GetDebugMode() != true {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_GetDbDriverName(t *testing.T) {
	if configInteractor.GetDbDriverName() != "" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_GetDbSourceName(t *testing.T) {
	if configInteractor.GetDbSourceName() != "" {
		t.Error("value is not match")
	}
}

func TestConfigInteractor_GetWebListen(t *testing.T) {
	if configInteractor.GetWebListen() != "String" {
		t.Error("value is not match")
	}
}
