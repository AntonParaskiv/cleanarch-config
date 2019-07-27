package ConfigInteractor

type ConfVarUint64 struct {
	confVar *ConfVar
}

func (cv *ConfVar) Uint64() (cvb *ConfVarUint64) {
	cv.varType = "uint64"
	cvb = &ConfVarUint64{confVar: cv}
	return
}

func (cv *ConfVarUint64) Get() (value uint64) {
	return cv.confVar.Value.(uint64)
}

func (cv *ConfVarUint64) Default(defaultValue uint64) *ConfVarUint64 {
	cv.confVar.isRequired = false
	cv.confVar.defaultValue = defaultValue
	return cv
}
