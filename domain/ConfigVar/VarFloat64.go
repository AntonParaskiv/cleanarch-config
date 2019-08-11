package ConfigVar

type VarFloat64 struct {
	configVar *Var
}

func (v *Var) Float64() (vf *VarFloat64) {
	v.varType = "float64"
	vf = &VarFloat64{configVar: v}
	return
}

func (v *VarFloat64) Get() (value float64) {
	return v.configVar.value.(float64)
}

func (v *VarFloat64) Default(defaultValue float64) *VarFloat64 {
	v.configVar.isRequired = false
	v.configVar.defaultValue = defaultValue
	return v
}
