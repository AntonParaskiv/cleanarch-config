package ConfigVar

import (
	"reflect"
	"testing"
)

func TestVar_Float64(t *testing.T) {
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
		wantVf *VarFloat64
	}{
		{
			name: "Success",
			wantVf: &VarFloat64{
				configVar: &Var{
					varType: "float64",
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
			if gotVf := v.Float64(); !reflect.DeepEqual(gotVf, tt.wantVf) {
				t.Errorf("Var.Float64() = %v, want %v", gotVf, tt.wantVf)
			}
		})
	}
}

func TestVarFloat64_Get(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue float64
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					value: -3.14,
				},
			},
			wantValue: -3.14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarFloat64{
				configVar: tt.fields.configVar,
			}
			if gotValue := v.Get(); gotValue != tt.wantValue {
				t.Errorf("VarFloat64.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestVarFloat64_Default(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	type args struct {
		defaultValue float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VarFloat64
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					isRequired: true,
				},
			},
			args: args{
				defaultValue: -3.14,
			},
			want: &VarFloat64{
				configVar: &Var{
					isRequired:   false,
					defaultValue: -3.14,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarFloat64{
				configVar: tt.fields.configVar,
			}
			if got := v.Default(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VarFloat64.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
