package ConfigVar

type VarUint64 struct {
	confVar *Var
}

func (v *Var) Uint64() (vu *VarUint64) {
	v.varType = "uint64"
	vu = &VarUint64{confVar: v}
	return
}

func (v *VarUint64) Get() (value uint64) {
	return v.confVar.value.(uint64)
}

func (v *VarUint64) Default(defaultValue uint64) *VarUint64 {
	v.confVar.isRequired = false
	v.confVar.defaultValue = defaultValue
	return v
}
