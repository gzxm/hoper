package log

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/liov/hoper/go/v2/utils/log/output"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
	Logger *zap.Logger
}

func GetLogger() *Logger {
	return Default
}

type Config struct {
	Development bool
	Caller      bool
	Level       zapcore.Level
	OutputPaths map[string][]string
	ModuleName  string //系统名称namespace.service
}

//初始化日志对象
func (lf *Config) NewLogger() *Logger {
	logger := lf.initLogger().
		With(
			zap.String("source", lf.ModuleName),
		)
	return &Logger{logger.Sugar(), logger}
}

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	l.Logger = l.Logger.WithOptions(opts...)
	l.SugaredLogger = l.Logger.Sugar()
	return l
}

var Default = (&Config{Development: true, Caller: true, Level: -1}).NewLogger()
var NoCall = (&Config{Development: true}).NewLogger()
var CallTwo = Default.WithOptions(zap.AddCallerSkip(2))

func init() {
	output.RegisterSink()
}

func (lf *Config) SetLogger() {
	Default.SugaredLogger = lf.initLogger().Sugar()
}

func (lf *Config) SetNoCall() {
	lf.Caller = false
	NoCall.SugaredLogger = lf.initLogger().Sugar()
}

//构建日志对象基本信息
func (lf *Config) initLogger() *zap.Logger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       lf.ModuleName,
		CallerKey:     "caller",
		FunctionKey:   "func",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006/01/02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(runtime.FuncForPC(caller.PC).Name() + ` ` + caller.TrimmedPath())
		},
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
	//如果没有设置输出，默认控制台
	if len(cores) == 0 {
		consoleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), lf.Level))
	}
	core := zapcore.NewTee(cores...)

	logger := zap.New(core, lf.hook()...)

	return logger
}

func (lf *Config) hook() []zap.Option {
	var hooks []zap.Option
	//系统名称
	if len(lf.OutputPaths["json"]) > 0 && lf.ModuleName != "" {
		hooks = append(hooks, zap.Fields(zap.Any("module", lf.ModuleName)))
	}

	if lf.Development {
		hooks = append(hooks, zap.Development(), zap.AddStacktrace(zapcore.DPanicLevel))
	}
	if lf.Caller {
		hooks = append(hooks, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	return hooks
}

func Sync() {
	Default.Sync()
}

func Print(v ...interface{}) {
	Default.Print(v...)
}

func Debug(v ...interface{}) {
	Default.Debug(v...)
}

func Info(v ...interface{}) {
	Default.Info(v...)
}

func Warn(format string, v ...interface{}) {
	Default.Warn(v...)
}

func Error(v ...interface{}) {
	Default.Error(v...)
}

func Panic(v ...interface{}) {
	Default.Panic(v...)
}

func Fatal(v ...interface{}) {
	Default.Fatal(v...)
}

func Printf(format string, v ...interface{}) {
	Default.Printf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	Default.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	Default.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	Default.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	Default.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	Default.Panicf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Default.Fatalf(format, v...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues)|0 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	Default.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues)|0 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	Default.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues)|0 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	Default.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues)|0 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	Default.Errorw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues)|0 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	Default.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues)|0 != 0 {
		keysAndValues = append(keysAndValues, "")
	}
	Default.Fatalw(msg, keysAndValues...)
}

// 兼容gormv1
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.Info(args...)
}

// 兼容grpclog
func (l *Logger) Infoln(v ...interface{}) {
	l.Info(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.Warn(v...)
}

func (l *Logger) Warningln(v ...interface{}) {
	l.Warn(v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.Warn(v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.Warn(v...)
}

func (l *Logger) V(level int) bool {
	if level == 3 {
		level = 5
	}
	return l.Logger.Core().Enabled(zapcore.Level(level))
}