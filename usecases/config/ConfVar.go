package config

type ConfVar struct {
	key          string
	Value        interface{}
	varType      string
	isRequired   bool
	defaultValue interface{}
	interactor   *ConfigInteractor
}

const (
	errRegisterConfigVarFailed = "register config var failed: "
	errEmptyKey                = errRegisterConfigVarFailed + "empty key"
)
