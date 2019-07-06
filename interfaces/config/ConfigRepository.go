package config

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type ConfigStorage interface {
	Vars() []string
	GetVar(key string) string
	SetVar(key, value string) error
	ClearAllVars()
}

type Logger interface {
	Debug(message string)
	Debugf(format string, a ...interface{})
	Info(message string)
	Infof(format string, a ...interface{})
	Warn(message string)
	Warnf(format string, a ...interface{})
	Error(message string)
	Errorf(format string, a ...interface{})
	Fatal(message string)
	Fatalf(format string, a ...interface{})
}

type ConfigRepository struct {
	storage ConfigStorage
	Log     Logger
}

func NewConfigRepository(storage ConfigStorage, log Logger) (cr *ConfigRepository) {
	cr = new(ConfigRepository)
	cr.storage = storage
	cr.Log = log
	return
}

func (cr *ConfigRepository) Vars() (Vars []string) {
	return cr.storage.Vars()
}

func (cr *ConfigRepository) GetString(key string) (value string) {
	value = cr.storage.GetVar(key)
	cr.Log.Debugf("%s=%s", key, value)
	return
}

func (cr *ConfigRepository) GetBool(key string) (value bool, filled bool, err error) {
	valueString := cr.storage.GetVar(key)
	cr.Log.Debugf("%s=%s", key, valueString)

	// value not filled
	if len(valueString) == 0 {
		return
	}
	filled = true

	value, err = strconv.ParseBool(valueString)
	if err != nil {
		err = errors.New("parseBool of " + key + " failed: " + err.Error())
		cr.Log.Errorf(err.Error())
	}

	return
}

func (cr *ConfigRepository) GetInt64(key string) (value int64, filled bool, err error) {
	valueString := cr.storage.GetVar(key)
	cr.Log.Debugf("%s=%s", key, valueString)

	// value not filled
	if len(valueString) == 0 {
		return
	}
	filled = true

	value, err = strconv.ParseInt(valueString, 10, 64)
	if err != nil {
		err = errors.New("parseInt of " + key + " failed: " + err.Error())
		cr.Log.Errorf(err.Error())
	}

	return
}

func (cr *ConfigRepository) GetUint64(key string) (value uint64, filled bool, err error) {
	valueString := cr.storage.GetVar(key)
	cr.Log.Debugf("%s=%s", key, valueString)

	// value not filled
	if len(valueString) == 0 {
		return
	}
	filled = true

	value, err = strconv.ParseUint(valueString, 10, 64)
	if err != nil {
		err = errors.New("parseUint of " + key + " failed: " + err.Error())
		cr.Log.Errorf(err.Error())
	}

	return
}

func (cr *ConfigRepository) GetFloat64(key string) (value float64, filled bool, err error) {
	valueString := cr.storage.GetVar(key)
	cr.Log.Debugf("%s=%s", key, valueString)

	// value not filled
	if len(valueString) == 0 {
		return
	}
	filled = true

	value, err = strconv.ParseFloat(valueString, 64)
	if err != nil {
		err = errors.New("parseFloat64 of " + key + " failed: " + err.Error())
		cr.Log.Errorf(err.Error())
	}

	return
}

func (cr *ConfigRepository) SetString(key, value string) (err error) {
	valueString := value
	if err := cr.storage.SetVar(key, valueString); err != nil {
		err = errors.New("SetVar failed: " + err.Error())
	}
	return
}

func (cr *ConfigRepository) SetBool(key string, value bool) (err error) {
	valueString := fmt.Sprintf("%t", value)
	if err := cr.storage.SetVar(key, valueString); err != nil {
		err = errors.New("SetVar failed: " + err.Error())
	}
	return
}

func (cr *ConfigRepository) SetInt64(key string, value int64) (err error) {
	valueString := fmt.Sprintf("%d", value)
	if err := cr.storage.SetVar(key, valueString); err != nil {
		err = errors.New("SetVar failed: " + err.Error())
	}
	return
}

func (cr *ConfigRepository) SetUint64(key string, value uint64) (err error) {
	valueString := fmt.Sprintf("%d", value)
	if err := cr.storage.SetVar(key, valueString); err != nil {
		err = errors.New("SetVar failed: " + err.Error())
	}
	return
}

func (cr *ConfigRepository) SetFloat64(key string, value float64) (err error) {
	valueString := fmt.Sprintf("%f", value)
	if err := cr.storage.SetVar(key, valueString); err != nil {
		err = errors.New("SetVar failed: " + err.Error())
	}
	return
}

func (cr *ConfigRepository) ClearAllVars() {
	cr.storage.ClearAllVars()
}
