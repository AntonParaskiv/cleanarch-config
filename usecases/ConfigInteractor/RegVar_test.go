package ConfigInteractor

import (
	"reflect"
	"testing"

	"github.com/AntonParaskiv/cleanarch-config/domain/ConfigVar"
)

func TestInteractor_RegVar(t *testing.T) {
	type fields struct {
		repository Repository
		prefix     string
		envMap     map[string]*ConfigVar.Var
		log        Logger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantConfigVar  *ConfigVar.Var
		wantInteractor *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				envMap: map[string]*ConfigVar.Var{},
			},
			args: args{
				key: "testKey",
			},
			wantConfigVar: ConfigVar.New("testKey"),
			wantInteractor: &Interactor{
				envMap: map[string]*ConfigVar.Var{
					"testKey": ConfigVar.New("testKey"),
				},
			},
		},
		{
			name: "SuccessWithPrefix",
			fields: fields{
				prefix: "prefix_",
				envMap: map[string]*ConfigVar.Var{},
			},
			args: args{
				key: "testKey",
			},
			wantConfigVar: ConfigVar.New("prefix_testKey"),
			wantInteractor: &Interactor{
				prefix: "prefix_",
				envMap: map[string]*ConfigVar.Var{
					"prefix_testKey": ConfigVar.New("prefix_testKey"),
				},
			},
		},
		//{ // panic
		//	name: "Error",
		//	fields: fields{
		//		prefix: "prefix_",
		//		envMap: map[string]*ConfigVar.Var{},
		//	},
		//	args: args{
		//		key: "",
		//	},
		//	wantConfigVar: ConfigVar.New("prefix_testKey"),
		//	wantInteractor: &Interactor{
		//		prefix: "prefix_",
		//		envMap: map[string]*ConfigVar.Var{},
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repository: tt.fields.repository,
				prefix:     tt.fields.prefix,
				envMap:     tt.fields.envMap,
				log:        tt.fields.log,
			}
			if gotConfigVar := i.RegVar(tt.args.key); !reflect.DeepEqual(gotConfigVar, tt.wantConfigVar) {
				t.Errorf("Interactor.RegVar() = %v, want %v", gotConfigVar, tt.wantConfigVar)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
