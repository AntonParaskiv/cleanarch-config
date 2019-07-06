package config

import "github.com/pkg/errors"

type ConfVarString struct {
	ConfVar *ConfVar
}

func (cv *ConfVarString) Get() (value string) {
	return cv.ConfVar.Value.(string)
}

func (ci *ConfigInteractor) RegisterVarString(key string, isRequired bool, defaultValue string) (confVar *ConfVarString, err error) {
	if key == "" {
		err = errors.Errorf("empty key")
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarString{
		ConfVar: &ConfVar{
			Key:          key,
			Type:         "string",
			IsRequired:   isRequired,
			DefaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
