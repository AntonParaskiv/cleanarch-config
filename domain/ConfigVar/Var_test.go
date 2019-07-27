package ConfigVar

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		wantV *Var
	}{
		{
			name: "Success",
			args: args{
				key: "TestKey",
			},
			wantV: &Var{
				key:        "TestKey",
				isRequired: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := New(tt.args.key); !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("New() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestVar_SetValue(t *testing.T) {
	type fields struct {
		key          string
		value        interface{}
		varType      string
		isRequired   bool
		defaultValue interface{}
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantV  *Var
	}{
		{
			name: "Success",
			args: args{
				value: "TestValue",
			},
			wantV: &Var{
				value: "TestValue",
			},
		},
		// TODO: Add test cases.
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
			v.SetValue(tt.args.value)
			if !reflect.DeepEqual(v, tt.wantV) {
				t.Errorf("Var = %v, want %v", v, tt.wantV)
			}
		})
	}
}

func TestVar_Type(t *testing.T) {
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
		want   string
	}{
		{
			name: "Success",
			fields: fields{
				varType: "string",
			},
			want: "string",
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
			if got := v.Type(); got != tt.want {
				t.Errorf("Var.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVar_IsRequired(t *testing.T) {
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
		want   bool
	}{
		{
			name: "Success",
			fields: fields{
				isRequired: true,
			},
			want: true,
		},
		// TODO: Add test cases.
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
			if got := v.IsRequired(); got != tt.want {
				t.Errorf("Var.IsRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVar_DefaultValue(t *testing.T) {
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
		want   interface{}
	}{
		{
			name: "Success",
			fields: fields{
				defaultValue: "DefaultValue",
			},
			want: "DefaultValue",
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
			if got := v.DefaultValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Var.DefaultValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
