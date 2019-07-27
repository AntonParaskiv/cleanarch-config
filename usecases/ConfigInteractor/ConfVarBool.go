package ConfigInteractor

type ConfVarBool struct {
	confVar *ConfVar
}

func (cv *ConfVar) Bool() (cvb *ConfVarBool) {
	cv.varType = "bool"
	cvb = &ConfVarBool{confVar: cv}
	return
}

func (cv *ConfVarBool) Get() (value bool) {
	return cv.confVar.Value.(bool)
}

func (cv *ConfVarBool) Default(defaultValue bool) *ConfVarBool {
	cv.confVar.isRequired = false
	cv.confVar.defaultValue = defaultValue
	return cv
}
