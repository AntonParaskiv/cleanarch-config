package ConfigRepository

import (
	"fmt"
	"github.com/AntonParaskiv/cleanarch-config/mocks/LoggerMock"
	"github.com/pkg/errors"
	"strconv"
)

const (
	ErrorParseFailed = "parse %s %s failed: %s"
	ErrorSetFailed   = "set failed: %s"
)

type Storage interface {
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

type Repository struct {
	storage Storage
	log     Logger
}

func New(storage Storage) (r *Repository) {
	r = new(Repository)
	r.storage = storage
	r.log = LoggerMock.New()
	return
}

func (r *Repository) SetLogger(log Logger) *Repository {
	if log == nil {
		log = LoggerMock.New()
	}
	r.log = log
	return r
}

func (r *Repository) lookupString(key string) (value string, isPresent bool) {
	value, isPresent = r.storage.Lookup(key)
	if isPresent {
		r.log.Debugf("%s=%s", key, value)
	} else {
		r.log.Debugf("%s is not present", key)
	}
	return
}

func (r *Repository) LookupString(key string) (value string, isPresent bool) {
	return r.lookupString(key)
}

func (r *Repository) LookupBool(key string) (value bool, isPresent bool, err error) {
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

func (r *Repository) LookupInt64(key string) (value int64, isPresent bool, err error) {
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

func (r *Repository) LookupUint64(key string) (value uint64, isPresent bool, err error) {
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

func (r *Repository) LookupFloat64(key string) (value float64, isPresent bool, err error) {
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

func (r *Repository) GetString(key string) (value string) {
	value, _ = r.lookupString(key)
	return
}

func (r *Repository) GetBool(key string) (value bool, err error) {
	value, _, err = r.LookupBool(key)
	return
}

func (r *Repository) GetInt64(key string) (value int64, err error) {
	value, _, err = r.LookupInt64(key)
	return
}

func (r *Repository) GetUint64(key string) (value uint64, err error) {
	value, _, err = r.LookupUint64(key)
	return
}

func (r *Repository) GetFloat64(key string) (value float64, err error) {
	value, _, err = r.LookupFloat64(key)
	return
}

func (r *Repository) SetString(key, value string) (err error) {
	valueString := value
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *Repository) SetBool(key string, value bool) (err error) {
	valueString := fmt.Sprintf("%t", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *Repository) SetInt64(key string, value int64) (err error) {
	valueString := fmt.Sprintf("%d", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *Repository) SetUint64(key string, value uint64) (err error) {
	valueString := fmt.Sprintf("%d", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *Repository) SetFloat64(key string, value float64) (err error) {
	valueString := fmt.Sprintf("%f", value)
	if err = r.storage.Set(key, valueString); err != nil {
		err = errors.Errorf(ErrorSetFailed, err.Error())
	}
	return
}

func (r *Repository) UnSet(key string) (err error) {
	return r.storage.UnSet(key)
}

func (r *Repository) Expand(sIn string) (sOut string) {
	return r.storage.Expand(sIn)
}

func (r *Repository) Vars() (Vars []string) {
	return r.storage.Vars()
}

func (r *Repository) ClearAll() {
	r.storage.ClearAll()
}
