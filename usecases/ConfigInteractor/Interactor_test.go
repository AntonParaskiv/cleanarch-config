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

func TestNew(t *testing.T) {
	type args struct {
		repository Repository
	}
	tests := []struct {
		name  string
		args  args
		wantI *Interactor
	}{
		{
			name: "Success",
			args: args{
				repository: ConfigRepository.New(ConfigStorageEnv.New(EnvStorageMock.New())),
			},
			wantI: &Interactor{
				repository: ConfigRepository.New(ConfigStorageEnv.New(EnvStorageMock.New())),
				envMap:     make(map[string]*ConfigVar.Var, 0),
				log:        LoggerMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(tt.args.repository); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInteractor_SetPrefix(t *testing.T) {
	type fields struct {
		repository Repository
		prefix     string
		envMap     map[string]*ConfigVar.Var
		log        Logger
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name: "Success",
			args: args{
				prefix: "testPrefix",
			},
			want: &Interactor{
				prefix: "testPrefix",
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
			if got := i.SetPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetLogger(t *testing.T) {
	type fields struct {
		repository Repository
		prefix     string
		envMap     map[string]*ConfigVar.Var
		log        Logger
	}
	type args struct {
		log Logger
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name: "Success",
			args: args{
				log: LoggerMock.New().SetMessage("testMessage"),
			},
			want: &Interactor{
				log: LoggerMock.New().SetMessage("testMessage"),
			},
		},
		{
			name: "SuccessNil",
			args: args{
				log: nil,
			},
			want: &Interactor{
				log: LoggerMock.New(),
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
			if got := i.SetLogger(tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
