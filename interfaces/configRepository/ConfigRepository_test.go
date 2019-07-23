package configRepository

import (
	"fmt"
	ConfigInfrastructure "github.com/AntonParaskiv/cleanarch-config/infrastructure/configStorageEnv"
	"github.com/AntonParaskiv/cleanarch-config/infrastructure/envStorageMock"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

type LoggerMock struct {
	message string
}

func (lm *LoggerMock) Debug(message string) {
	lm.message = message
}

func (lm *LoggerMock) Debugf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Info(message string) {
	lm.message = message
}

func (lm *LoggerMock) Infof(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Warn(message string) {
	lm.message = message
}

func (lm *LoggerMock) Warnf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Error(message string) {
	lm.message = message
}

func (lm *LoggerMock) Errorf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (lm *LoggerMock) Fatal(message string) {
	lm.message = message
}

func (lm *LoggerMock) Fatalf(format string, a ...interface{}) {
	lm.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func TestNew(t *testing.T) {
	type args struct {
		storage ConfigStorage
		log     Logger
	}
	tests := []struct {
		name   string
		args   args
		wantCr *ConfigRepository
	}{
		{
			name: "Success",
			args: args{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				log:     new(LoggerMock),
			},
			wantCr: &ConfigRepository{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCr := New(tt.args.storage, tt.args.log); !reflect.DeepEqual(gotCr, tt.wantCr) {
				t.Errorf("New() = %v, want %v", gotCr, tt.wantCr)
			}
		})
	}
}

func TestConfigRepository_lookupString(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantValue     string
		wantIsPresent bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     "value",
			wantIsPresent: true,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     "",
			wantIsPresent: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, gotIsPresent := r.lookupString(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.lookupString() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsPresent != tt.wantIsPresent {
				t.Errorf("ConfigRepository.lookupString() gotIsPresent = %v, want %v", gotIsPresent, tt.wantIsPresent)
			}
		})
	}
}

func TestConfigRepository_LookupString(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantValue     string
		wantIsPresent bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     "value",
			wantIsPresent: true,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     "",
			wantIsPresent: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, gotIsPresent := r.LookupString(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.LookupString() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsPresent != tt.wantIsPresent {
				t.Errorf("ConfigRepository.LookupString() gotIsPresent = %v, want %v", gotIsPresent, tt.wantIsPresent)
			}
		})
	}
}

func TestConfigRepository_LookupBool(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantValue     bool
		wantIsPresent bool
		wantErr       bool
	}{
		{
			name: "IsPresentTrue",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "true"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     true,
			wantIsPresent: true,
		},
		{
			name: "IsPresentFalse",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "false"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     false,
			wantIsPresent: true,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     false,
			wantIsPresent: false,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantIsPresent: true,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, gotIsPresent, err := r.LookupBool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.LookupBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.LookupBool() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsPresent != tt.wantIsPresent {
				t.Errorf("ConfigRepository.LookupBool() gotIsPresent = %v, want %v", gotIsPresent, tt.wantIsPresent)
			}
		})
	}
}

func TestConfigRepository_LookupInt64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantValue     int64
		wantIsPresent bool
		wantErr       bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "-1000"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     -1000,
			wantIsPresent: true,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     0,
			wantIsPresent: false,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantIsPresent: true,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, gotIsPresent, err := r.LookupInt64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.LookupInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.LookupInt64() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsPresent != tt.wantIsPresent {
				t.Errorf("ConfigRepository.LookupInt64() gotIsPresent = %v, want %v", gotIsPresent, tt.wantIsPresent)
			}
		})
	}
}

func TestConfigRepository_LookupUint64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantValue     uint64
		wantIsPresent bool
		wantErr       bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "1000"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     1000,
			wantIsPresent: true,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     0,
			wantIsPresent: false,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantIsPresent: true,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, gotIsPresent, err := r.LookupUint64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.LookupUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.LookupUint64() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsPresent != tt.wantIsPresent {
				t.Errorf("ConfigRepository.LookupUint64() gotIsPresent = %v, want %v", gotIsPresent, tt.wantIsPresent)
			}
		})
	}
}

func TestConfigRepository_LookupFloat64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantValue     float64
		wantIsPresent bool
		wantErr       bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "3.14"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     3.14,
			wantIsPresent: true,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue:     0,
			wantIsPresent: false,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantIsPresent: true,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, gotIsPresent, err := r.LookupFloat64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.LookupFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.LookupFloat64() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsPresent != tt.wantIsPresent {
				t.Errorf("ConfigRepository.LookupFloat64() gotIsPresent = %v, want %v", gotIsPresent, tt.wantIsPresent)
			}
		})
	}
}

func TestConfigRepository_GetString(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: "value",
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if gotValue := r.GetString(tt.args.key); gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.GetString() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestConfigRepository_GetBool(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue bool
		wantErr   bool
	}{
		{
			name: "IsPresentTrue",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "true"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: true,
		},
		{
			name: "IsPresentFalse",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "false"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: false,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: false,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, err := r.GetBool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.GetBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.GetBool() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestConfigRepository_GetInt64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int64
		wantErr   bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "-1000"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: -1000,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: 0,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, err := r.GetInt64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.GetInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.GetInt64() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestConfigRepository_GetUint64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue uint64
		wantErr   bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "1000"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: 1000,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: 0,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, err := r.GetUint64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.GetUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.GetUint64() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestConfigRepository_GetFloat64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue float64
		wantErr   bool
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "3.14"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: 3.14,
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantValue: 0,
		},
		{
			name: "ParseErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
				Log:     new(LoggerMock),
			},
			args: args{
				key: "key",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			gotValue, err := r.GetFloat64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.GetFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ConfigRepository.GetFloat64() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestConfigRepository_SetString(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: "value",
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "value"})),
		},
		{
			name: "StorageErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SimulateError()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: "value",
			},
			wantErr:     true,
			wantStorage: ConfigInfrastructure.New(envStorageMock.New()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if err := r.SetString(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.SetString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
			}
		})
	}
}

func TestConfigRepository_SetBool(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key   string
		value bool
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: true,
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "true"})),
		},
		{
			name: "StorageErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SimulateError()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: true,
			},
			wantErr:     true,
			wantStorage: ConfigInfrastructure.New(envStorageMock.New()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if err := r.SetBool(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.SetBool() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
			}
		})
	}
}

func TestConfigRepository_SetInt64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: -1000,
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "-1000"})),
		},
		{
			name: "StorageErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SimulateError()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: -1000,
			},
			wantErr:     true,
			wantStorage: ConfigInfrastructure.New(envStorageMock.New()),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if err := r.SetInt64(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.SetInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
			}
		})
	}
}

func TestConfigRepository_SetUint64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key   string
		value uint64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: 1000,
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "1000"})),
		},
		{
			name: "StorageErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SimulateError()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: 1000,
			},
			wantErr:     true,
			wantStorage: ConfigInfrastructure.New(envStorageMock.New()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if err := r.SetUint64(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.SetUint64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
			}
		})
	}
}

func TestConfigRepository_SetFloat64(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: 3.14,
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{"key": "3.140000"})),
		},
		{
			name: "StorageErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SimulateError()),
				Log:     new(LoggerMock),
			},
			args: args{
				key:   "key",
				value: 3.14,
			},
			wantErr:     true,
			wantStorage: ConfigInfrastructure.New(envStorageMock.New()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if err := r.SetFloat64(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.SetFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
				spew.Dump(r.storage)
			}
		})
	}
}

func TestConfigRepository_UnSet(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				})),
				Log: new(LoggerMock),
			},
			args: args{
				key: "key2",
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{
				"key1": "value1",
				"key3": "value3",
			})),
		},
		{
			name: "StorageErr",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SimulateError().SetStorageMap(map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				})),
				Log: new(LoggerMock),
			},
			args: args{
				key: "key2",
			},
			wantErr: true,
			wantStorage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if err := r.UnSet(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepository.UnSet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
				spew.Dump(r.storage)
			}
		})
	}
}

func TestConfigRepository_Expand(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	type args struct {
		sIn string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSOut string
	}{
		{
			name: "IsPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{
					"key1": "value1",
				})),
				Log: new(LoggerMock),
			},
			args: args{
				sIn: "key1 is $key1",
			},
			wantSOut: "key1 is value1",
		},
		{
			name: "IsNotPresent",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New()),
				Log:     new(LoggerMock),
			},
			args: args{
				sIn: "key1 is $key1",
			},
			wantSOut: "key1 is ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if gotSOut := r.Expand(tt.args.sIn); gotSOut != tt.wantSOut {
				t.Errorf("ConfigRepository.Expand() = %v, want %v", gotSOut, tt.wantSOut)
			}
		})
	}
}

func TestConfigRepository_Vars(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	tests := []struct {
		name     string
		fields   fields
		wantVars []string
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				})),
				Log: new(LoggerMock),
			},
			wantVars: []string{
				"key1=value1",
				"key2=value2",
				"key3=value3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			if gotVars := r.Vars(); !sameStringSlice(gotVars, tt.wantVars) {
				t.Errorf("ConfigRepository.Vars() = %v, want %v", gotVars, tt.wantVars)
			}
		})
	}
}

func TestConfigRepository_ClearAll(t *testing.T) {
	type fields struct {
		storage ConfigStorage
		Log     Logger
	}
	tests := []struct {
		name        string
		fields      fields
		wantStorage ConfigStorage
	}{
		{
			name: "Success",
			fields: fields{
				storage: ConfigInfrastructure.New(envStorageMock.New().SetStorageMap(map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				})),
				Log: new(LoggerMock),
			},
			wantStorage: ConfigInfrastructure.New(envStorageMock.New()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConfigRepository{
				storage: tt.fields.storage,
				Log:     tt.fields.Log,
			}
			r.ClearAll()

			if !reflect.DeepEqual(tt.wantStorage, r.storage) {
				t.Errorf("ConfigRepository.SetString() storage = %v, wantStorage %v", r.storage, tt.wantStorage)
				spew.Dump(r.storage)
			}
		})
	}
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
