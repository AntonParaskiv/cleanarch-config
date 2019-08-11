package ConfigVar

type Var struct {
	key          string
	value        interface{}
	varType      string
	isRequired   bool
	defaultValue interface{}
}

func New(key string) (v *Var) {
	v = new(Var)
	v.key = key
	v.isRequired = true
	return
}

func (v *Var) SetValue(value interface{}) {
	v.value = value
}

func (v *Var) Type() string {
	return v.varType
}

func (v *Var) IsRequired() bool {
	return v.isRequired
}

func (v *Var) DefaultValue() interface{} {
	return v.defaultValue
}
