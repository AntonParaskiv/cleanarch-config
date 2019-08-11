package ConfigVar

type VarInt64 struct {
	configVar *Var
}

func (v *Var) Int64() (vi *VarInt64) {
	v.varType = "int64"
	vi = &VarInt64{configVar: v}
	return
}

func (v *VarInt64) Get() (value int64) {
	return v.configVar.value.(int64)
}

func (v *VarInt64) Default(defaultValue int64) *VarInt64 {
	v.configVar.isRequired = false
	v.configVar.defaultValue = defaultValue
	return v
}
