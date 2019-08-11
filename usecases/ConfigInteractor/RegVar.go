package ConfigInteractor

import (
	"github.com/AntonParaskiv/cleanarch-config/domain/ConfigVar"
	"github.com/pkg/errors"
)

const (
	errRegisterConfigVarFailed = "register config var failed: "
	errEmptyKey                = errRegisterConfigVarFailed + "empty key"
)

func (i *Interactor) RegVar(key string) (configVar *ConfigVar.Var) {
	if key == "" {
		err := errors.Errorf(errEmptyKey)
		panic(err)
		return
	}
	key = i.prefix + key

	configVar = ConfigVar.New(key)
	i.envMap[key] = configVar
	return
}
