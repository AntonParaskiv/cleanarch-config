package config

import "github.com/pkg/errors"

type ConfVarInt64 struct {
	ConfVar *ConfVar
}

func (cv *ConfVarInt64) Get() (value int64) {
	return cv.ConfVar.Value.(int64)
}

func (ci *ConfigInteractor) RegisterVarInt64(key string, isRequired bool, defaultValue int64) (confVar *ConfVarInt64) {
	if key == "" {
		err := errors.Errorf(errEmptyKey)
		ci.errs = append(ci.errs, err)
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarInt64{
		ConfVar: &ConfVar{
			key:          key,
			varType:      "int64",
			isRequired:   isRequired,
			defaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
