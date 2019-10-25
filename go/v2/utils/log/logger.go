package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func (l *Logger) Print(args ...interface{}) {
	l.Info(args...)
}

func GetLogger() *Logger {
	return logger
}

type Config struct {
	Development bool
	Skip        bool
	Level       zapcore.Level
	OutputPaths map[string][]string
	ModuleName  string //系统名称namespace.service
}

//初始化日志对象
func (lf *Config) NewLogger() *Logger {
	return &Logger{
		lf.initLogger().Sugar(),
	}
}

var logger *Logger = (&Config{Development: true, Skip: true}).NewLogger()
var NoCall = (&Config{Development: true}).NewLogger()

func (lf *Config) SetLogger() {
	logger.SugaredLogger = lf.initLogger().Sugar()
}

func (lf *Config) SetNoCall() {
	lf.Skip = false
	NoCall.SugaredLogger = lf.initLogger().Sugar()
}

//构建日志对象基本信息
func (lf *Config) initLogger() *zap.Logger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    lf.ModuleName,
		CallerKey:  "caller",
		MessageKey: "msg",
		//StacktraceKey: "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if lf.Development {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	}

	var consoleEncoder, jsonEncoder zapcore.Encoder
	var cores []zapcore.Core

	if len(lf.OutputPaths["console"]) > 0 {
		consoleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
		sink, _, err := zap.Open(lf.OutputPaths["console"]...)
		if err != nil {
			log.Fatal(err)
		}
		cores = append(cores, zapcore.NewCore(consoleEncoder, sink, lf.Level))
	}

	if len(lf.OutputPaths["json"]) > 0 {
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		jsonEncoder = zapcore.NewJSONEncoder(encoderConfig)
		sink, _, err := zap.Open(lf.OutputPaths["json"]...)
		if err != nil {
			log.Fatal(err)
		}
		cores = append(cores, zapcore.NewCore(jsonEncoder, sink, lf.Level))
	}

	if len(cores) == 0 {
		consoleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), lf.Level))
	}
	core := zapcore.NewTee(cores...)

	logger := zap.New(core, lf.hook()..., )

	return logger
}

func (lf *Config) hook() []zap.Option {
	var hooks []zap.Option
	//系统名称
	if len(lf.OutputPaths["json"]) > 0 && lf.ModuleName != "" {
		hooks = append(hooks, zap.Fields(zap.Any("module", lf.ModuleName)))
	}

	if lf.Development {
		hooks = append(hooks, zap.Development())
	}
	if lf.Skip {
		hooks = append(hooks, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	return hooks
}

func Sync() {
	logger.Sync()
}

func Print(v ...interface{}) {

}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(format string, v ...interface{}) {
	logger.Warn(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Fatal(msg)
}

func Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Debug(msg)
}

func Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Info(msg)
}

func Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Warn(msg)
}

func Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Error(msg)
}
