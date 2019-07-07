package config

import "github.com/pkg/errors"

type ConfVarBool struct {
	ConfVar *ConfVar
}

func (cv *ConfVarBool) Get() (value bool) {
	return cv.ConfVar.Value.(bool)
}

func (ci *ConfigInteractor) RegisterVarBool(key string, isRequired bool, defaultValue bool) (confVar *ConfVarBool) {
	if key == "" {
		err := errors.Errorf(errEmptyKey)
		ci.errs = append(ci.errs, err)
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarBool{
		ConfVar: &ConfVar{
			key:          key,
			varType:      "bool",
			isRequired:   isRequired,
			defaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
