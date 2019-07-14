package configStorageEnv

type EnvStorage interface {
	Getenv(key string) (value string)
	Setenv(key, value string) (err error)
	Unsetenv(key string) (err error)
	ExpandEnv(sIn string) (sOut string)
	LookupEnv(key string) (value string, isPresent bool)
	Environ() (envStrings []string)
	Clearenv()
}

type ConfigStorageEnv struct {
	storage EnvStorage
}

func New(storage EnvStorage) (cse *ConfigStorageEnv) {
	cse = new(ConfigStorageEnv)
	cse.storage = storage
	return
}

func (cse *ConfigStorageEnv) Get(key string) (value string) {
	value = cse.storage.Getenv(key)
	return
}

func (cse *ConfigStorageEnv) Set(key, value string) (err error) {
	err = cse.storage.Setenv(key, value)
	return
}

func (cse *ConfigStorageEnv) UnSet(key string) (err error) {
	err = cse.storage.Unsetenv(key)
	return
}

func (cse *ConfigStorageEnv) Expand(sIn string) (sOut string) {
	sOut = cse.storage.ExpandEnv(sIn)
	return
}

func (cse *ConfigStorageEnv) Lookup(key string) (value string, isPresent bool) {
	value, isPresent = cse.storage.LookupEnv(key)
	return
}

func (cse *ConfigStorageEnv) Vars() (vars []string) {
	vars = cse.storage.Environ()
	return
}

func (cse *ConfigStorageEnv) ClearAll() {
	cse.storage.Clearenv()
}
