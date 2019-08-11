package LoggerMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantM *LoggerMock
	}{
		{
			name:  "Success",
			wantM: &LoggerMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := New(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("New() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Message(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success",
			fields: fields{
				message: "TestMessage",
			},
			want: "TestMessage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			if got := m.Message(); got != tt.want {
				t.Errorf("LoggerMock.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerMock_Debug(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				message: "testMessage",
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Debug(tt.args.message)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Debugf(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				format: "test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Debugf(tt.args.format, tt.args.a...)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Info(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				message: "testMessage",
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Info(tt.args.message)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Infof(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				format: "test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Infof(tt.args.format, tt.args.a...)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Warn(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				message: "testMessage",
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Warn(tt.args.message)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Warnf(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				format: "test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Warnf(tt.args.format, tt.args.a...)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Error(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				message: "testMessage",
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Error(tt.args.message)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Errorf(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				format: "test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Errorf(tt.args.format, tt.args.a...)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Fatal(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				message: "testMessage",
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Fatal(tt.args.message)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_Fatalf(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				format: "test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantM: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			m.Fatalf(tt.args.format, tt.args.a...)
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("LoggerMock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestLoggerMock_SetMessage(t *testing.T) {
	type fields struct {
		message string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LoggerMock
	}{
		{
			name: "Success",
			args: args{
				message: "testMessage",
			},
			want: &LoggerMock{
				message: "testMessage",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LoggerMock{
				message: tt.fields.message,
			}
			if got := m.SetMessage(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerMock.SetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
