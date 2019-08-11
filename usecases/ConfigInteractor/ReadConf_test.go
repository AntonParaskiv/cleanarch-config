package ConfigInteractor

import (
	"github.com/AntonParaskiv/cleanarch-config/infrastructure/ConfigStorageEnv"
	"github.com/AntonParaskiv/cleanarch-config/infrastructure/EnvStorageMock"
	"github.com/AntonParaskiv/cleanarch-config/interfaces/ConfigRepository"
	"github.com/AntonParaskiv/cleanarch-config/mocks/LoggerMock"
	"reflect"
	"testing"

	"github.com/AntonParaskiv/cleanarch-config/domain/ConfigVar"
)

func TestInteractor_ReadConf(t *testing.T) {
	envVarString := ConfigVar.New("testVarString")
	envVarString.String()
	wantEnvVarString := ConfigVar.New("testVarString")
	wantEnvVarString.String()
	wantEnvVarString.SetValue("testValue")

	envVarBool := ConfigVar.New("testVarBool")
	envVarBool.Bool()
	wantEnvVarBool := ConfigVar.New("testVarBool")
	wantEnvVarBool.Bool()
	wantEnvVarBool.SetValue(true)

	envVarInt64 := ConfigVar.New("testVarInt64")
	envVarInt64.Int64()
	wantEnvVarInt64 := ConfigVar.New("testVarInt64")
	wantEnvVarInt64.Int64()
	wantEnvVarInt64.SetValue(int64(-314))

	envVarUint64 := ConfigVar.New("testVarUint64")
	envVarUint64.Uint64()
	wantEnvVarUint64 := ConfigVar.New("testVarUint64")
	wantEnvVarUint64.Uint64()
	wantEnvVarUint64.SetValue(uint64(314))

	envVarFloat64 := ConfigVar.New("testVarFloat64")
	envVarFloat64.Float64()
	wantEnvVarFloat64 := ConfigVar.New("testVarFloat64")
	wantEnvVarFloat64.Float64()
	wantEnvVarFloat64.SetValue(-3.14)

	envVarStringDefaultValue := ConfigVar.New("testVarStringDefaultValue")
	envVarStringDefaultValue.String().Default("testValue")
	wantEnvVarStringDefaultValue := ConfigVar.New("testVarStringDefaultValue")
	wantEnvVarStringDefaultValue.SetValue("testValue")
	wantEnvVarStringDefaultValue.String().Default("testValue")

	// Errors
	envVarStringNoDefaultValue := ConfigVar.New("testVarStringNoDefaultValue")
	envVarStringNoDefaultValue.String()

	envVarBoolLookupError := ConfigVar.New("testVarBoolLookupError")
	envVarBoolLookupError.Bool()

	envVarNoType := ConfigVar.New("testVarNoType")

	type fields struct {
		repository Repository
		prefix     string
		envMap     map[string]*ConfigVar.Var
		log        Logger
	}
	tests := []struct {
		name         string
		fields       fields
		wantErrsLen  int
		wantVarValue interface{}
		wantEnvMap   map[string]*ConfigVar.Var
	}{
		{
			name: "SuccessString",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarString": "testValue",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarString": envVarString,
				},
				log: LoggerMock.New(),
			},

			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarString": wantEnvVarString,
			},
		},
		{
			name: "SuccessBool",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarBool": "true",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarBool": envVarBool,
				},
				log: LoggerMock.New(),
			},

			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarBool": wantEnvVarBool,
			},
		},
		{
			name: "SuccessInt64",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarInt64": "-314",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarInt64": envVarInt64,
				},
				log: LoggerMock.New(),
			},

			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarInt64": wantEnvVarInt64,
			},
		},
		{
			name: "SuccessUint64",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarUint64": "314",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarUint64": envVarUint64,
				},
				log: LoggerMock.New(),
			},

			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarUint64": wantEnvVarUint64,
			},
		},
		{
			name: "SuccessFloat64",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarFloat64": "-3.14",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarFloat64": envVarFloat64,
				},
				log: LoggerMock.New(),
			},

			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarFloat64": wantEnvVarFloat64,
			},
		},
		{
			name: "SuccessStringNotPresentDefaultValue",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarStringDefaultValue": envVarStringDefaultValue,
				},
				log: LoggerMock.New(),
			},

			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarStringDefaultValue": wantEnvVarStringDefaultValue,
			},
		},
		{
			name: "ErrorStringNotPresentNoDefaultValue",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarStringNotPresent": envVarStringNoDefaultValue,
				},
				log: LoggerMock.New(),
			},
			wantErrsLen: 1,
			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarStringNotPresent": envVarStringNoDefaultValue,
			},
		},
		{
			name: "ErrorLookup",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarBoolLookupError": "string to bool convert failed",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarBoolLookupError": envVarBoolLookupError,
				},
				log: LoggerMock.New(),
			},
			wantErrsLen: 1,
			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarBoolLookupError": envVarBoolLookupError,
			},
		},
		{
			name: "ErrorLookup",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarBoolLookupError": "string to bool convert failed",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarBoolLookupError": envVarBoolLookupError,
				},
				log: LoggerMock.New(),
			},
			wantErrsLen: 1,
			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarBoolLookupError": envVarBoolLookupError,
			},
		},
		{
			name: "ErrorLookup",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarBoolLookupError": "string to bool convert failed",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarBoolLookupError": envVarBoolLookupError,
				},
				log: LoggerMock.New(),
			},
			wantErrsLen: 1,
			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarBoolLookupError": envVarBoolLookupError,
			},
		},
		{
			name: "ErrorNoType",
			fields: fields{
				repository: ConfigRepository.New(ConfigStorageEnv.New(&EnvStorageMock.MockEnv{
					Storage: map[string]string{
						"testVarNoType": "no type",
					},
				})),
				envMap: map[string]*ConfigVar.Var{
					"testVarNoType": envVarNoType,
				},
				log: LoggerMock.New(),
			},
			wantErrsLen: 1,
			wantEnvMap: map[string]*ConfigVar.Var{
				"testVarNoType": envVarNoType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repository: tt.fields.repository,
				prefix:     tt.fields.prefix,
				envMap:     tt.fields.envMap,
				log:        tt.fields.log,
			}
			if gotErrs := i.ReadConf(); len(gotErrs) != tt.wantErrsLen {
				t.Errorf("Interactor.ReadConf().ErrsLen = %v, want %v", len(gotErrs), tt.wantErrsLen)
			}
			if !reflect.DeepEqual(i.envMap, tt.wantEnvMap) {
				t.Errorf("Interactor.envMap %v, want %v", i.envMap, tt.wantEnvMap)
			}
		})
	}
}
