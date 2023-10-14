package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"go.uber.org/fx/fxevent"
	"io"
	"os"
	"reflect"
	"runtime"
)

type (
	StackInfo struct {
		file     string
		line     int
		funcName string
	}
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() (*Logger, error) {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.TraceLevel)

	return &Logger{
		logger: l,
	}, nil
}

func (l *Logger) LogEvent(event fxevent.Event) {
	eventDetail, err := StructToMap(event)
	if err != nil {
		panic(err)
	}
	eventName := reflect.TypeOf(event).Elem().Name()

	if eventDetail["Err"] == nil {
		l.logger.
			WithFields(map[string]interface{}{"event": eventDetail}).
			Trace(fmt.Sprintf("[Fx Event] %s", eventName))
	} else {
		l.logger.
			WithFields(map[string]interface{}{"event": eventDetail}).
			Error(fmt.Sprintf("[Fx Event] %s: %+v", eventName, eventDetail["Err"]))
	}
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (l *Logger) SimpleInfoF(format string, args ...interface{}) {
	stackInfo := l.makeStackInfo(runtime.Caller(1))
	l.logger.
		WithFields(l.makeCommonFields(stackInfo, nil)).
		Infof(format, args...)
}

func (l *Logger) SimpleFatal(e error, params map[string]interface{}) {
	stackInfo := l.makeStackInfo(runtime.Caller(1))
	l.logger.
		WithFields(l.makeCommonFields(stackInfo, params)).
		Fatalf("%+v", e)
}

func (l *Logger) makeCommonFields(stackInfo *StackInfo, params map[string]interface{}) map[string]interface{} {
	var function *string
	var file *string
	var line *int
	if stackInfo != nil {
		function = &stackInfo.funcName
		file = &stackInfo.file
		line = &stackInfo.line
	}

	hostname, _ := os.Hostname()

	return map[string]interface{}{
		"params":   params,
		"function": function,
		"file":     file,
		"line":     line,
		"host":     hostname,
	}
}

func (l *Logger) makeStackInfo(pc uintptr, file string, line int, ok bool) *StackInfo {
	if !ok {
		return nil
	}

	funcName := runtime.FuncForPC(pc).Name()
	return &StackInfo{
		file:     file,
		line:     line,
		funcName: funcName,
	}
}

func (l *Logger) SetupForTest() *test.Hook {
	l.logger.SetLevel(logrus.TraceLevel)
	l.logger.SetOutput(io.Discard)
	return test.NewLocal(l.logger)
}
