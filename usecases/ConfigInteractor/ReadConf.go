package ConfigInteractor

import "github.com/pkg/errors"

const (
	errReadConfFailed         = "read config failed: "
	errKeyMustBeFilled        = errReadConfFailed + "%s must be filled"
	errGetKeyFailedErr        = errReadConfFailed + "get %s failed: %s"
	errGetKeyFailedUnkownType = errReadConfFailed + "get %s failed: unknown type %s"
)

func (i *ConfigInteractor) ReadConf() (errs []error) {
	i.log.Debug("config reading started")

	var value interface{}
	var isPresent bool
	var err error

	// for all vars
	for key, confVar := range i.envMap {

		// get value from repo
		switch confVar.varType {
		case "string":
			value, isPresent = i.repo.LookupString(key)
			err = nil
		case "bool":
			value, isPresent, err = i.repo.LookupBool(key)
		case "int64":
			value, isPresent, err = i.repo.LookupInt64(key)
		case "uint64":
			value, isPresent, err = i.repo.LookupUint64(key)
		case "float64":
			value, isPresent, err = i.repo.LookupFloat64(key)
		default:
			err := errors.Errorf(errGetKeyFailedUnkownType, key, confVar.varType)
			errs = append(errs, err)
			continue
		}

		if err != nil {
			err := errors.Errorf(errGetKeyFailedErr, key, err.Error())
			errs = append(errs, err)
			continue
		}

		if !isPresent {
			// check required
			if confVar.isRequired {
				err := errors.Errorf(errKeyMustBeFilled, key)
				errs = append(errs, err)
				continue
			} else {
				// set default value
				confVar.Value = confVar.defaultValue
			}
		} else {
			confVar.Value = value
		}

	}

	// check errs
	if len(errs) == 0 {
		errs = nil
	}

	i.log.Debug("config reading finished")
	return
}
