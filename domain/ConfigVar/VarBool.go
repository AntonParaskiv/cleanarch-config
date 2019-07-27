package ConfigVar

type VarBool struct {
	configVar *Var
}

func (v *Var) Bool() (vb *VarBool) {
	v.varType = "bool"
	vb = &VarBool{configVar: v}
	return
}

func (v *VarBool) Get() (value bool) {
	return v.configVar.value.(bool)
}

func (v *VarBool) Default(defaultValue bool) *VarBool {
	v.configVar.isRequired = false
	v.configVar.defaultValue = defaultValue
	return v
}
