package ConfigVar

type VarInt64 struct {
	confVar *Var
}

func (v *Var) Int64() (vi *VarInt64) {
	v.varType = "int64"
	vi = &VarInt64{confVar: v}
	return
}

func (v *VarInt64) Get() (value int64) {
	return v.confVar.value.(int64)
}

func (v *VarInt64) Default(defaultValue int64) *VarInt64 {
	v.confVar.isRequired = false
	v.confVar.defaultValue = defaultValue
	return v
}
