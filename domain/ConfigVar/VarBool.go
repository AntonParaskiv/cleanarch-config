package ConfigVar

type VarBool struct {
	confVar *Var
}

func (v *Var) Bool() (vb *VarBool) {
	v.varType = "bool"
	vb = &VarBool{confVar: v}
	return
}

func (v *VarBool) Get() (value bool) {
	return v.confVar.value.(bool)
}

func (v *VarBool) Default(defaultValue bool) *VarBool {
	v.confVar.isRequired = false
	v.confVar.defaultValue = defaultValue
	return v
}
