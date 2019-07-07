package config

import "github.com/pkg/errors"

const (
	errReadConfFailed  = "read config failed: "
	errKeyMustBeFilled = errReadConfFailed + "%s must be filled"
	errGetKeyFailedErr = errReadConfFailed + "get %s failed: %s"
)

func (ci *ConfigInteractor) ReadConf() (errReached bool) {
	ci.Log.Debug("config reading started")

	// for all vars
	for key, confVar := range ci.envMap {

		switch confVar.varType {

		case "string":
			// get value from repo
			value := ci.repo.GetString(key)

			if value == "" {
				// check require
				if confVar.isRequired {
					err := errors.Errorf(errKeyMustBeFilled, key)
					ci.errs = append(ci.errs, err)
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.defaultValue
				}
			} else {
				confVar.Value = value
			}

		case "bool":
			// get value from repo
			value, filled, err := ci.repo.GetBool(key)
			if err != nil {
				err := errors.Errorf(errGetKeyFailedErr, key, err.Error())
				ci.errs = append(ci.errs, err)
				errReached = true
				continue
			}
			if !filled {
				if confVar.isRequired {
					err := errors.Errorf(errKeyMustBeFilled, key)
					ci.errs = append(ci.errs, err)
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.defaultValue
				}
			} else {
				confVar.Value = value
			}

		case "int64":
			// get value from repo
			value, filled, err := ci.repo.GetInt64(key)
			if err != nil {
				err := errors.Errorf(errGetKeyFailedErr, key, err.Error())
				ci.errs = append(ci.errs, err)
				errReached = true
				continue
			}
			if !filled {
				if confVar.isRequired {
					err := errors.Errorf(errKeyMustBeFilled, key)
					ci.errs = append(ci.errs, err)
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.defaultValue
				}
			} else {
				confVar.Value = value
			}

		case "uint64":
			// get value from repo
			value, filled, err := ci.repo.GetUint64(key)
			if err != nil {
				err := errors.Errorf(errGetKeyFailedErr, key, err.Error())
				ci.errs = append(ci.errs, err)
				errReached = true
				continue
			}
			if !filled {
				if confVar.isRequired {
					err := errors.Errorf(errKeyMustBeFilled, key)
					ci.errs = append(ci.errs, err)
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.defaultValue
				}
			} else {
				confVar.Value = value
			}

		case "float64":
			// get value from repo
			value, filled, err := ci.repo.GetFloat64(key)
			if err != nil {
				err := errors.Errorf(errGetKeyFailedErr, key, err.Error())
				ci.errs = append(ci.errs, err)
				errReached = true
				continue
			}
			if !filled {
				if confVar.isRequired {
					err := errors.Errorf(errKeyMustBeFilled, key)
					ci.errs = append(ci.errs, err)
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.defaultValue
				}
			} else {
				confVar.Value = value
			}
		}
	}

	ci.Log.Debug("config reading finished")
	return
}
