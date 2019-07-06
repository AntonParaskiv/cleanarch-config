package config

import "github.com/pkg/errors"

type ConfVarBool struct {
	ConfVar *ConfVar
}

func (cv *ConfVarBool) Get() (value bool) {
	return cv.ConfVar.Value.(bool)
}

func (ci *ConfigInteractor) RegisterVarBool(key string, isRequired bool, defaultValue bool) (confVar *ConfVarBool, err error) {
	if key == "" {
		err = errors.Errorf("empty key")
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarBool{
		ConfVar: &ConfVar{
			Key:          key,
			Type:         "bool",
			IsRequired:   isRequired,
			DefaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
