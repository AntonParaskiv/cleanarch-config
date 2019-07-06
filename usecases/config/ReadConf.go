package config

import "github.com/pkg/errors"

func (ci *ConfigInteractor) ReadConf() (errReached bool) {
	ci.Log.Debug("config reading started")

	// for all vars
	for key, confVar := range ci.envMap {

		switch confVar.Type {

		case "string":
			// get value from repo
			value := ci.repo.GetString(key)

			if value == "" {
				// check require
				if confVar.IsRequired {
					err := errors.Errorf("%s must be filled", key)
					ci.Log.Error(err.Error())
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.DefaultValue
				}
			} else {
				confVar.Value = value
			}

		case "bool":
			// get value from repo
			value, filled, err := ci.repo.GetBool(key)
			if err != nil {
				err := errors.Errorf("get %s failed: %s", key, err.Error())
				ci.Log.Error(err.Error())
				errReached = true
				continue
			}
			if !filled {
				if confVar.IsRequired {
					err := errors.Errorf("%s must be filled", key)
					ci.Log.Error(err.Error())
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.DefaultValue
				}
			} else {
				confVar.Value = value
			}

		case "int64":
			// get value from repo
			value, filled, err := ci.repo.GetInt64(key)
			if err != nil {
				err := errors.Errorf("get %s failed: %s", key, err.Error())
				ci.Log.Error(err.Error())
				errReached = true
				continue
			}
			if !filled {
				if confVar.IsRequired {
					err := errors.Errorf("%s must be filled", key)
					ci.Log.Error(err.Error())
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.DefaultValue
				}
			} else {
				confVar.Value = value
			}

		case "uint64":
			// get value from repo
			value, filled, err := ci.repo.GetUint64(key)
			if err != nil {
				err := errors.Errorf("get %s failed: %s", key, err.Error())
				ci.Log.Error(err.Error())
				errReached = true
				continue
			}
			if !filled {
				if confVar.IsRequired {
					err := errors.Errorf("%s must be filled", key)
					ci.Log.Error(err.Error())
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.DefaultValue
				}
			} else {
				confVar.Value = value
			}

		case "float64":
			// get value from repo
			value, filled, err := ci.repo.GetFloat64(key)
			if err != nil {
				err := errors.Errorf("get %s failed: %s", key, err.Error())
				ci.Log.Error(err.Error())
				errReached = true
				continue
			}
			if !filled {
				if confVar.IsRequired {
					err := errors.Errorf("%s must be filled", key)
					ci.Log.Error(err.Error())
					errReached = true
					continue
				} else {
					// set default value
					confVar.Value = confVar.DefaultValue
				}
			} else {
				confVar.Value = value
			}
		}
	}

	ci.Log.Debug("config reading finished")
	return
}
