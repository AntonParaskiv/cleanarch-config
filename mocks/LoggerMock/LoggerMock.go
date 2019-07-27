package LoggerMock

import "fmt"

type LoggerMock struct {
	message string
}

func New() (m *LoggerMock) {
	m = new(LoggerMock)
	return
}

func (m *LoggerMock) Message() string {
	return m.message
}

func (m *LoggerMock) Debug(message string) {
	m.message = message
}

func (m *LoggerMock) Debugf(format string, a ...interface{}) {
	m.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (m *LoggerMock) Info(message string) {
	m.message = message
}

func (m *LoggerMock) Infof(format string, a ...interface{}) {
	m.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (m *LoggerMock) Warn(message string) {
	m.message = message
}

func (m *LoggerMock) Warnf(format string, a ...interface{}) {
	m.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (m *LoggerMock) Error(message string) {
	m.message = message
}

func (m *LoggerMock) Errorf(format string, a ...interface{}) {
	m.message = fmt.Sprintf(format, fmt.Sprint(a...))
}

func (m *LoggerMock) Fatal(message string) {
	m.message = message
}

func (m *LoggerMock) Fatalf(format string, a ...interface{}) {
	m.message = fmt.Sprintf(format, fmt.Sprint(a...))
}
