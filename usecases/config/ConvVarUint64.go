package config

import "github.com/pkg/errors"

type ConfVarUint64 struct {
	ConfVar *ConfVar
}

func (cv *ConfVarUint64) Get() (value uint64) {
	return cv.ConfVar.Value.(uint64)
}

func (ci *ConfigInteractor) RegisterVarUint64(key string, isRequired bool, defaultValue uint64) (confVar *ConfVarUint64) {
	if key == "" {
		err := errors.Errorf(errEmptyKey)
		ci.errs = append(ci.errs, err)
		return
	}
	key = ci.prefix + key

	confVar = &ConfVarUint64{
		ConfVar: &ConfVar{
			key:          key,
			varType:      "uint64",
			isRequired:   isRequired,
			defaultValue: defaultValue,
			interactor:   ci,
		},
	}

	ci.envMap[key] = confVar.ConfVar
	return
}
