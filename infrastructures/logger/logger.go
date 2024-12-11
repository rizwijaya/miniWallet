package logger

import (
	"log"
	"os"

	"github.com/rizwijaya/miniWallet/infrastructures/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.SugaredLogger
	debugMode bool
)

func ErrorLogger(mainDir string) (zapcore.Core, error) {
	// Konfigurasi untuk output error log ke file
	errorEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleSink := zapcore.AddSync(os.Stdout)

	errorFile, err := os.Create(mainDir + "miniWallet_error.log")
	if err != nil {
		return nil, err
	}

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	fileSink := zapcore.AddSync(errorFile)
	multiSink := zapcore.NewMultiWriteSyncer(fileSink, consoleSink)

	return zapcore.NewCore(errorEncoder, multiSink, errorLevel), nil
}

func InfoLogger(mainDir string) (zapcore.Core, error) {
	// Konfigurasi untuk output info log ke file
	infoEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleSink := zapcore.AddSync(os.Stdout)

	infoFile, err := os.Create(mainDir + "miniWallet_info.log")
	if err != nil {
		return nil, err
	}

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	fileSink := zapcore.AddSync(infoFile)
	multiSink := zapcore.NewMultiWriteSyncer(fileSink, consoleSink)

	return zapcore.NewCore(infoEncoder, multiSink, infoLevel), nil
}

func NewLogger(config config.LoadConfig) {
	mainDir := "./log/"
	errorCore, err := ErrorLogger(mainDir)
	if err != nil {
		log.Fatal(err)
	}

	infoCore, err := InfoLogger(mainDir)
	if err != nil {
		log.Fatal(err)
	}

	// Gabungkan core untuk logger
	core := zapcore.NewTee(errorCore, infoCore)

	// Buat logger berdasarkan konfigurasi
	logger = zap.New(core).Sugar()
	debugMode = config.App.Debug
}

func Sync() {
	logger.Sync()
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func Debug(args ...interface{}) {
	if debugMode {
		logger.Debug(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if debugMode {
		logger.Debugf(template, args...)
	}
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}
