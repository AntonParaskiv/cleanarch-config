package config

import "github.com/pkg/errors"

type ConfVarFloat64 struct {
	ConfVar *ConfVar
}

func (cv *ConfVarFloat64) Get() (value float64) {
	return cv.ConfVar.Value.(float64)
}

func (ci *ConfigInteractor) RegisterVarFloat64(key string, isRequired bool, defaultValue float64) (confVar *ConfVarFloat64, err error) {
	if key == "" {
		err = errors.Errorf("empty key")
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarFloat64{
		ConfVar: &ConfVar{
			Key:          key,
			Type:         "float64",
			IsRequired:   isRequired,
			DefaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
