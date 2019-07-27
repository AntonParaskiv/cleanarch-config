package ConfigInteractor

type ConfVarFloat64 struct {
	confVar *ConfVar
}

func (cv *ConfVar) Float64() (cvb *ConfVarFloat64) {
	cv.varType = "float64"
	cvb = &ConfVarFloat64{confVar: cv}
	return
}

func (cv *ConfVarFloat64) Get() (value float64) {
	return cv.confVar.Value.(float64)
}

func (cv *ConfVarFloat64) Default(defaultValue float64) *ConfVarFloat64 {
	cv.confVar.isRequired = false
	cv.confVar.defaultValue = defaultValue
	return cv
}
