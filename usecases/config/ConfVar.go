package config

type ConfVar struct {
	Key          string
	Value        interface{}
	Type         string
	IsRequired   bool
	DefaultValue interface{}
	interactor   *ConfigInteractor
}
