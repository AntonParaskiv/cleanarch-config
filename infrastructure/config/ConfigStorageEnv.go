package config

import "github.com/adammck/venv"

type ConfigStorageEnv struct {
	storage venv.Env
}

func (c *ConfigStorageEnv) Vars() (Vars []string) {
	return c.storage.Environ()
}

func (c *ConfigStorageEnv) GetVar(key string) (Value string) {
	return c.storage.Getenv(key)
}

func (c *ConfigStorageEnv) SetVar(key, value string) (err error) {
	return c.storage.Setenv(key, value)
}

func (c *ConfigStorageEnv) ClearAllVars() {
	c.storage.Clearenv()
}

func NewConfigStorageEnv(storage venv.Env) (cse *ConfigStorageEnv) {
	cse = new(ConfigStorageEnv)
	cse.storage = storage
	return
}
