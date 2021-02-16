package common

import (
	"html/template"

	"github.com/micro/go-micro/v2/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	fileName := "micro.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			fileName:   fileName,
			MaxSize:    521,
			MaxBackups: 0,
			LocalTime:  true,
			Commpress:  true,
		})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		zap.NewAtomicLevelAt(zap.DebugLevel)
	)
	log:=zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1)
	)
	logger = log.Sugar()
}

func Debug(args ...interface{}){
	logger.Debug(args)
}

func Debugf(template string, args ...interface{}){
	logger.Debugf(template,args...)
}
func Info(args ...interface{}){
	logger.Info(args...)
}
func Info(template string, args ...interface{}){
	logger.Infof(template,args...)
}

func Warn(args ...interface{}){
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}){
	logger.Warnf(template,args...)
}
func Error(args ...interface){
	logger.Error(args...)
}
func Errorf(template string, args ...interface{}){
	logger.Errorf(template,args...)
}

func DPainc(args ...interface{}){
	logger.DPainc(args...)
}
func DPaincf(template string, args ...interface{}){
	logger.DPaincf(template,args...)
}
