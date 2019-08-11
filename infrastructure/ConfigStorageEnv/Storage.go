package ConfigStorageEnv

type EnvStorage interface {
	Getenv(key string) (value string)
	Setenv(key, value string) (err error)
	Unsetenv(key string) (err error)
	ExpandEnv(sIn string) (sOut string)
	LookupEnv(key string) (value string, isPresent bool)
	Environ() (envStrings []string)
	Clearenv()
}

type ConfigStorage struct {
	storage EnvStorage
}

func New(storage EnvStorage) (cse *ConfigStorage) {
	cse = new(ConfigStorage)
	cse.storage = storage
	return
}

func (cse *ConfigStorage) Get(key string) (value string) {
	value = cse.storage.Getenv(key)
	return
}

func (cse *ConfigStorage) Set(key, value string) (err error) {
	err = cse.storage.Setenv(key, value)
	return
}

func (cse *ConfigStorage) UnSet(key string) (err error) {
	err = cse.storage.Unsetenv(key)
	return
}

func (cse *ConfigStorage) Expand(sIn string) (sOut string) {
	sOut = cse.storage.ExpandEnv(sIn)
	return
}

func (cse *ConfigStorage) Lookup(key string) (value string, isPresent bool) {
	value, isPresent = cse.storage.LookupEnv(key)
	return
}

func (cse *ConfigStorage) Vars() (vars []string) {
	vars = cse.storage.Environ()
	return
}

func (cse *ConfigStorage) ClearAll() {
	cse.storage.Clearenv()
}
