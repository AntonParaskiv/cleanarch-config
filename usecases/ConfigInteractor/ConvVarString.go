package ConfigInteractor

type ConfVarString struct {
	confVar *ConfVar
}

func (cv *ConfVar) String() (cvb *ConfVarString) {
	cv.varType = "string"
	cvb = &ConfVarString{confVar: cv}
	return
}

func (cv *ConfVarString) Get() (value string) {
	return cv.confVar.Value.(string)
}

func (cv *ConfVarString) Default(defaultValue string) *ConfVarString {
	cv.confVar.isRequired = false
	cv.confVar.defaultValue = defaultValue
	return cv
}
