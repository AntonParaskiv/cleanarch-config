package config

import "github.com/pkg/errors"

type ConfVarString struct {
	confVar *ConfVar
}

func (cv *ConfVarString) Get() (value string) {
	return cv.confVar.Value.(string)
}

func (ci *ConfigInteractor) RegisterVarString(key string, isRequired bool, defaultValue string) (confVar *ConfVarString) {
	if key == "" {
		err := errors.Errorf(errEmptyKey)
		ci.errs = append(ci.errs, err)
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarString{
		confVar: &ConfVar{
			key:          key,
			varType:      "string",
			isRequired:   isRequired,
			defaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.confVar
	return
}
