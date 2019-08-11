package ConfigVar

import (
	"reflect"
	"testing"
)

func TestVar_Uint64(t *testing.T) {
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
		wantVu *VarUint64
	}{
		{
			name: "Success",
			wantVu: &VarUint64{
				configVar: &Var{
					varType: "uint64",
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
			if gotVu := v.Uint64(); !reflect.DeepEqual(gotVu, tt.wantVu) {
				t.Errorf("Var.Uint64() = %v, want %v", gotVu, tt.wantVu)
			}
		})
	}
}

func TestVarUint64_Get(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue uint64
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					value: uint64(314),
				},
			},
			wantValue: uint64(314),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarUint64{
				configVar: tt.fields.configVar,
			}
			if gotValue := v.Get(); gotValue != tt.wantValue {
				t.Errorf("VarUint64.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestVarUint64_Default(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	type args struct {
		defaultValue uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VarUint64
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					isRequired: true,
				},
			},
			args: args{
				defaultValue: uint64(314),
			},
			want: &VarUint64{
				configVar: &Var{
					isRequired:   false,
					defaultValue: uint64(314),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarUint64{
				configVar: tt.fields.configVar,
			}
			if got := v.Default(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VarUint64.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
