package ConfigRepository

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

const (
	ErrorParseFailed = "parse %s %s failed: %s"
	ErrorSetFailed   = "set failed: %s"
)

type ConfigStorage interface {
	Get(key string) (value string)
	Set(key, value string) (err error)
	UnSet(key string) (err error)
	Expand(sIn string) (sOut string)
	Lookup(key string) (value string, isPresent bool)
	Vars() (vars []string)
	ClearAll()
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
	log     Logger
}

func New(storage ConfigStorage) (cr *ConfigRepository) {
	cr = new(ConfigRepository)
	cr.storage = storage
	return
}

func (cr *ConfigRepository) SetLogger(log Logger) *ConfigRepository {
	cr.log = log
	return cr
}

func (r *ConfigRepository) lookupString(key string) (value string, isPresent bool) {
	value, isPresent = r.storage.Lookup(key)
	if isPresent {
		r.log.Debugf("%s=%s", key, value)
	} else {
		r.log.Debugf("%s is not present", key)
	}
	return
}

func (r *ConfigRepository) LookupString(key string) (value string, isPresent bool) {
	return r.lookupString(key)
}

func (r *ConfigRepository) LookupBool(key string) (value bool, isPresent bool, err error) {
	valueString, isPresent := r.lookupString(key)
	if isPresent {
		value, err = strconv.ParseBool(valueString)
		if err != nil {
			err = errors.Errorf(ErrorParseFailed, "bool", key, err.Error())
			r.log.Errorf(err.Error())
		}
	}
	return
}

func (r *ConfigRepository) LookupInt64(key string) (value int64, isPresent bool, err error) {
	valueString, isPresent := r.lookupString(key)
	if isPresent {
		value, err = strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			err = errors.Errorf(ErrorParseFailed, "int64", key, err.Error())
			r.log.Errorf(err.Error())
		}
	}
	return
}

func (r *ConfigRepository) LookupUint64(key string) (value uint64, isPresent bool, err error) {
	valueString, isPresent := r.lookupString(key)
	if isPresent {
		value, err = strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			err = errors.Errorf(ErrorParseFailed, "uint64", key, err.Error())
			r.log.Errorf(err.Error())
		}
	}
	return
}

func (r *ConfigRepository) LookupFloat64(key string) (value float64, isPresent bool, err error) {
	valueString, isPresent := r.lookupString(key)
	if isPresent {
		value, err = strconv.ParseFloat(valueString, 64)
		if err != nil {
			err = errors.Errorf(ErrorParseFailed, "float64", key, err.Error())
			r.log.Errorf(err.Error())
		}
	}
	return
}

func (r *ConfigRepository) GetString(key string) (value string) {
	value, _ = r.lookupString(key)
	return
}

func (r *ConfigRepository) GetBool(key string) (value bool, err error) {
	value, _, err = r.LookupBool(key)
	return
}

func (r *ConfigRepository) GetInt64(key string) (value int64, err error) {
	value, _, err = r.LookupInt64(key)
	return
}

func (r *ConfigRepository) GetUint64(key string) (value uint64, err error) {
	value, _, err = r.LookupUint64(key)
	return
}

func (r *ConfigRepository) GetFloat64(key string) (value float64, err error) {
	value, _, err = r.LookupFloat64(key)
	return
}

func (r *ConfigRepository) SetString(key, value string) (err error) {
	valueString := value
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *ConfigRepository) SetBool(key string, value bool) (err error) {
	valueString := fmt.Sprintf("%t", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *ConfigRepository) SetInt64(key string, value int64) (err error) {
	valueString := fmt.Sprintf("%d", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *ConfigRepository) SetUint64(key string, value uint64) (err error) {
	valueString := fmt.Sprintf("%d", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *ConfigRepository) SetFloat64(key string, value float64) (err error) {
	valueString := fmt.Sprintf("%f", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *ConfigRepository) UnSet(key string) (err error) {
	return r.storage.UnSet(key)
}

func (r *ConfigRepository) Expand(sIn string) (sOut string) {
	return r.storage.Expand(sIn)
}

func (r *ConfigRepository) Vars() (Vars []string) {
	return r.storage.Vars()
}

func (r *ConfigRepository) ClearAll() {
	r.storage.ClearAll()
}
