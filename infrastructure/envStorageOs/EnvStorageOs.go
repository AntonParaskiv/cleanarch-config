package envStorageOs

import (
	"os"
)

type OsEnv struct {
}

func New() (osEnv *OsEnv) {
	osEnv = new(OsEnv)
	return
}

func (osEnv *OsEnv) Getenv(key string) (value string) {
	value = os.Getenv(key)
	return
}

func (osEnv *OsEnv) Setenv(key, value string) (err error) {
	err = os.Setenv(key, value)
	return
}

func (osEnv *OsEnv) Unsetenv(key string) (err error) {
	err = os.Unsetenv(key)
	return
}

func (osEnv *OsEnv) ExpandEnv(sIn string) (sOut string) {
	sOut = os.ExpandEnv(sIn)
	return
}

func (osEnv *OsEnv) LookupEnv(key string) (value string, isPresent bool) {
	value, isPresent = os.LookupEnv(key)
	return
}

func (osEnv *OsEnv) Environ() (envStrings []string) {
	envStrings = os.Environ()
	return
}

func (osEnv *OsEnv) Clearenv() {
	os.Clearenv()
}
