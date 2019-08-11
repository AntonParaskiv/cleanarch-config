package ConfigVar

import (
	"reflect"
	"testing"
)

func TestVar_Int64(t *testing.T) {
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
		wantVi *VarInt64
	}{
		{
			name: "Success",
			wantVi: &VarInt64{
				configVar: &Var{
					varType: "int64",
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
			if gotVi := v.Int64(); !reflect.DeepEqual(gotVi, tt.wantVi) {
				t.Errorf("Var.Int64() = %v, want %v", gotVi, tt.wantVi)
			}
		})
	}
}

func TestVarInt64_Get(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue int64
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					value: int64(-314),
				},
			},
			wantValue: int64(-314),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarInt64{
				configVar: tt.fields.configVar,
			}
			if gotValue := v.Get(); gotValue != tt.wantValue {
				t.Errorf("VarInt64.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestVarInt64_Default(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	type args struct {
		defaultValue int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VarInt64
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					isRequired: true,
				},
			},
			args: args{
				defaultValue: int64(-314),
			},
			want: &VarInt64{
				configVar: &Var{
					isRequired:   false,
					defaultValue: int64(-314),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarInt64{
				configVar: tt.fields.configVar,
			}
			if got := v.Default(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VarInt64.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
