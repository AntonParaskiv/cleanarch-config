package ConfigVar

type VarUint64 struct {
	configVar *Var
}

func (v *Var) Uint64() (vu *VarUint64) {
	v.varType = "uint64"
	vu = &VarUint64{configVar: v}
	return
}

func (v *VarUint64) Get() (value uint64) {
	return v.configVar.value.(uint64)
}

func (v *VarUint64) Default(defaultValue uint64) *VarUint64 {
	v.configVar.isRequired = false
	v.configVar.defaultValue = defaultValue
	return v
}
