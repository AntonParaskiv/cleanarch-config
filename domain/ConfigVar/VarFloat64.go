package ConfigVar

type VarFloat64 struct {
	confVar *Var
}

func (v *Var) Float64() (vf *VarFloat64) {
	v.varType = "float64"
	vf = &VarFloat64{confVar: v}
	return
}

func (v *VarFloat64) Get() (value float64) {
	return v.confVar.value.(float64)
}

func (v *VarFloat64) Default(defaultValue float64) *VarFloat64 {
	v.confVar.isRequired = false
	v.confVar.defaultValue = defaultValue
	return v
}
