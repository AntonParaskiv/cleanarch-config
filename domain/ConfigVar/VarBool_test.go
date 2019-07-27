package ConfigVar

import (
	"reflect"
	"testing"
)

func TestVar_Bool(t *testing.T) {
	type fields struct {
		key          string
		value        interface{}
		varType      string
		isRequired   bool
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		wantVb *VarBool
	}{
		{
			name: "Success",
			wantVb: &VarBool{
				configVar: &Var{
					varType: "bool",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Var{
				key:          tt.fields.key,
				value:        tt.fields.value,
				varType:      tt.fields.varType,
				isRequired:   tt.fields.isRequired,
				defaultValue: tt.fields.defaultValue,
			}
			if gotVb := v.Bool(); !reflect.DeepEqual(gotVb, tt.wantVb) {
				t.Errorf("Var.Bool() = %v, want %v", gotVb, tt.wantVb)
			}
		})
	}
}

func TestVarBool_Get(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue bool
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					value: true,
				},
			},
			wantValue: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarBool{
				configVar: tt.fields.configVar,
			}
			if gotValue := v.Get(); gotValue != tt.wantValue {
				t.Errorf("VarBool.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestVarBool_Default(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	type args struct {
		defaultValue bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VarBool
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					isRequired: true,
				},
			},
			args: args{
				defaultValue: true,
			},
			want: &VarBool{
				configVar: &Var{
					isRequired:   false,
					defaultValue: true,
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarBool{
				configVar: tt.fields.configVar,
			}
			if got := v.Default(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VarBool.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
