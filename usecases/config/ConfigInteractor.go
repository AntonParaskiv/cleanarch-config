package config

type ConfigRepository interface {
	Vars() (Vars []string)
	GetString(key string) (Value string)
	GetBool(key string) (Value bool, filled bool, err error)
	GetInt64(key string) (Value int64, filled bool, err error)
	GetUint64(key string) (Value uint64, filled bool, err error)
	GetFloat64(key string) (Value float64, filled bool, err error)
	SetString(key, Value string) (err error)
	SetBool(key string, Value bool) (err error)
	SetInt64(key string, Value int64) (err error)
	SetUint64(key string, Value uint64) (err error)
	SetFloat64(key string, Value float64) (err error)
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

type ConfigInteractor struct {
	repo   ConfigRepository
	prefix string
	envMap map[string]*ConfVar
	Log    Logger
}

func NewConfigInteractor(repo ConfigRepository, prefix string, log Logger) (ci *ConfigInteractor) {
	ci = new(ConfigInteractor)
	ci.repo = repo
	ci.prefix = prefix
	ci.Log = log
	ci.envMap = make(map[string]*ConfVar, 0)
	return
}
