package config

import "github.com/pkg/errors"

type ConfVarUint64 struct {
	ConfVar *ConfVar
}

func (cv *ConfVarUint64) Get() (value uint64) {
	return cv.ConfVar.Value.(uint64)
}

func (ci *ConfigInteractor) RegisterVarUint64(key string, isRequired bool, defaultValue uint64) (confVar *ConfVarUint64, err error) {
	if key == "" {
		err = errors.Errorf("empty key")
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarUint64{
		ConfVar: &ConfVar{
			Key:          key,
			Type:         "uint64",
			IsRequired:   isRequired,
			DefaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
