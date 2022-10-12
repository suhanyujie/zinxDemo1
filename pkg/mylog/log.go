package mylog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func NewLogger() *zap.Logger {
	w := GetWriter()
	enc := GetEncoder()
	// core 的核心在于：writer encoder levelEnable 三组配置
	core := zapcore.NewCore(enc, w, zapcore.DebugLevel)
	// l := zap.New(core)
	l := zap.New(core, zap.AddCaller())
	l.Sugar()

	return l
}

func GetLogger() *zap.Logger {
	if logger == nil {
		logger = NewLogger()
	}

	return logger
}

func GetWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	//return zapcore.AddSync(file)
	writer := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxAge:     5,
		MaxBackups: 7,
		LocalTime:  false,
		Compress:   false,
	}

	return zapcore.AddSync(writer)
}

func GetEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	// 设置时间编码配置
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// level 的展示使用大写字母
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(config)
}
