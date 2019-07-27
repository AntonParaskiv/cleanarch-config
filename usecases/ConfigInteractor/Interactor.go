package ConfigInteractor

type ConfigRepository interface {
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

type ConfigInteractor struct {
	repo   ConfigRepository
	prefix string
	envMap map[string]*ConfVar
	log    Logger
}

func New(repo ConfigRepository) (ci *ConfigInteractor) {
	ci = new(ConfigInteractor)
	ci.repo = repo
	ci.envMap = make(map[string]*ConfVar, 0)
	return
}

func (ci *ConfigInteractor) SetPrefix(prefix string) *ConfigInteractor {
	ci.prefix = prefix
	return ci
}

func (ci *ConfigInteractor) SetLogger(log Logger) *ConfigInteractor {
	ci.log = log
	return ci
}
