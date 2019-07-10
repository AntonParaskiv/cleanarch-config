package configStorageEnv

type EnvStorageInterface interface {
	Getenv(key string) (value string)
	Setenv(key, value string) (err error)
	Unsetenv(key string) (err error)
	ExpandEnv(sIn string) (sOut string)
	LookupEnv(key string) (value string, isPresent bool)
	Environ() (envStrings []string)
	Clearenv()
}

type ConfigStorageEnv struct {
	storage EnvStorageInterface
}

func New(storage EnvStorageInterface) (cse *ConfigStorageEnv) {
	cse = new(ConfigStorageEnv)
	cse.storage = storage
	return
}

func (cse *ConfigStorageEnv) GetVar(key string) (value string) {
	value = cse.storage.Getenv(key)
	return
}

func (cse *ConfigStorageEnv) SetVar(key, value string) (err error) {
	err = cse.storage.Setenv(key, value)
	return
}

func (cse *ConfigStorageEnv) UnSetVar(key string) (err error) {
	err = cse.storage.Unsetenv(key)
	return
}

func (cse *ConfigStorageEnv) ExpandVar(sIn string) (sOut string) {
	sOut = cse.storage.ExpandEnv(sIn)
	return
}

func (cse *ConfigStorageEnv) LookupVar(key string) (value string, isPresent bool) {
	value, isPresent = cse.storage.LookupEnv(key)
	return
}

func (cse *ConfigStorageEnv) Vars() (vars []string) {
	vars = cse.storage.Environ()
	return
}

func (cse *ConfigStorageEnv) ClearAllVars() {
	cse.storage.Clearenv()
}
