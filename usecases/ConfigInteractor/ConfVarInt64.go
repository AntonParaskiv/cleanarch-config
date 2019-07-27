package ConfigInteractor

type ConfVarInt64 struct {
	confVar *ConfVar
}

func (cv *ConfVar) Int64() (cvb *ConfVarInt64) {
	cv.varType = "int64"
	cvb = &ConfVarInt64{confVar: cv}
	return
}

func (cv *ConfVarInt64) Get() (value int64) {
	return cv.confVar.Value.(int64)
}

func (cv *ConfVarInt64) Default(defaultValue int64) *ConfVarInt64 {
	cv.confVar.isRequired = false
	cv.confVar.defaultValue = defaultValue
	return cv
}
