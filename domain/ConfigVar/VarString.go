package ConfigVar

type VarString struct {
	configVar *Var
}

func (v *Var) String() (vs *VarString) {
	v.varType = "string"
	vs = &VarString{configVar: v}
	return
}

func (v *VarString) Get() (value string) {
	return v.configVar.value.(string)
}

func (v *VarString) Default(defaultValue string) *VarString {
	v.configVar.isRequired = false
	v.configVar.defaultValue = defaultValue
	return v
}
