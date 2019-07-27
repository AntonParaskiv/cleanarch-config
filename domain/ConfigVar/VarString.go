package ConfigVar

type VarString struct {
	confVar *Var
}

func (v *Var) String() (vs *VarString) {
	v.varType = "string"
	vs = &VarString{confVar: v}
	return
}

func (v *VarString) Get() (value string) {
	return v.confVar.value.(string)
}

func (v *VarString) Default(defaultValue string) *VarString {
	v.confVar.isRequired = false
	v.confVar.defaultValue = defaultValue
	return v
}
