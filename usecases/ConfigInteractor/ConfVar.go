package ConfigInteractor

import "github.com/pkg/errors"

const (
	errRegisterConfigVarFailed = "register config var failed: "
	errEmptyKey                = errRegisterConfigVarFailed + "empty key"
)

type ConfVar struct {
	key          string
	Value        interface{}
	varType      string
	isRequired   bool
	defaultValue interface{}
	interactor   *ConfigInteractor
}

func (i *ConfigInteractor) RegVar(key string) (confVar *ConfVar) {
	if key == "" {
		err := errors.Errorf(errEmptyKey)
		panic(err)
		return
	}
	key = i.prefix + key

	confVar = &ConfVar{
		key:          key,
		varType:      "",
		isRequired:   true,
		defaultValue: nil,
		interactor:   i,
	}

	i.envMap[key] = confVar
	return
}
