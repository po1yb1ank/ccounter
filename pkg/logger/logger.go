package logger

import "go.uber.org/zap"

type ILogger interface {
	Error(string)
	Info(string)
	Debug(string)
}

type ZapSugaredLogger struct {
	logger *zap.SugaredLogger
}

func NewZapSugaredLogger() ILogger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("failed to init zap sugared logger")
	}

	return &ZapSugaredLogger{
		logger: zapLogger.Sugar(),
	}
}

func (l *ZapSugaredLogger) Error(msg string) {
	l.logger.Error(msg)
}
func (l *ZapSugaredLogger) Info(msg string) {
	l.logger.Info(msg)
}
func (l *ZapSugaredLogger) Debug(msg string) {
	l.logger.Debug(msg)
}
