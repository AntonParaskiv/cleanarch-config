package ConfigInteractor

import (
	"github.com/AntonParaskiv/cleanarch-config/domain/ConfigVar"
	"github.com/AntonParaskiv/cleanarch-config/mocks/LoggerMock"
)

type Repository interface {
	LookupString(key string) (value string, isPresent bool)
	LookupBool(key string) (value bool, isPresent bool, err error)
	LookupInt64(key string) (value int64, isPresent bool, err error)
	LookupUint64(key string) (value uint64, isPresent bool, err error)
	LookupFloat64(key string) (value float64, isPresent bool, err error)
	GetString(key string) (value string)
	GetBool(key string) (value bool, err error)
	GetInt64(key string) (value int64, err error)
	GetUint64(key string) (value uint64, err error)
	GetFloat64(key string) (value float64, err error)
	SetString(key, value string) (err error)
	SetBool(key string, value bool) (err error)
	SetInt64(key string, value int64) (err error)
	SetUint64(key string, value uint64) (err error)
	SetFloat64(key string, value float64) (err error)
	UnSet(key string) (err error)
	Expand(sIn string) (sOut string)
	Vars() (Vars []string)
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

type Interactor struct {
	repository Repository
	prefix     string
	envMap     map[string]*ConfigVar.Var
	log        Logger
}

func New(repository Repository) (i *Interactor) {
	i = new(Interactor)
	i.repository = repository
	i.envMap = make(map[string]*ConfigVar.Var, 0)
	i.log = LoggerMock.New()
	return
}

func (i *Interactor) SetPrefix(prefix string) *Interactor {
	i.prefix = prefix
	return i
}

func (i *Interactor) SetLogger(log Logger) *Interactor {
	if log == nil {
		log = LoggerMock.New()
	}
	i.log = log
	return i
}
