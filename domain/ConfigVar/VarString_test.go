package ConfigVar

import (
	"reflect"
	"testing"
)

func TestVar_String(t *testing.T) {
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
		wantVs *VarString
	}{
		{
			name: "Success",
			wantVs: &VarString{
				configVar: &Var{
					varType: "string",
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
			if gotVs := v.String(); !reflect.DeepEqual(gotVs, tt.wantVs) {
				t.Errorf("Var.String() = %v, want %v", gotVs, tt.wantVs)
			}
		})
	}
}

func TestVarString_Get(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue string
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					value: "text",
				},
			},
			wantValue: "text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarString{
				configVar: tt.fields.configVar,
			}
			if gotValue := v.Get(); gotValue != tt.wantValue {
				t.Errorf("VarString.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestVarString_Default(t *testing.T) {
	type fields struct {
		configVar *Var
	}
	type args struct {
		defaultValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VarString
	}{
		{
			name: "Success",
			fields: fields{
				configVar: &Var{
					isRequired: true,
				},
			},
			args: args{
				defaultValue: "text",
			},
			want: &VarString{
				configVar: &Var{
					isRequired:   false,
					defaultValue: "text",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VarString{
				configVar: tt.fields.configVar,
			}
			if got := v.Default(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VarString.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
