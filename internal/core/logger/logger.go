package core_logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct{
	*zap.Logger

	file *os.File
}

func FromContext (ctx context.Context) *Logger {
	log, ok := ctx.Value("log").(*Logger)
	if !ok {
		panic("no logger in context")
	}
	return log
}

func NewLogger(config Config) (*Logger, error) {
	zapLvl := zap.NewAtomicLevel()
	if err:= zapLvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("unmarshal log level %w", err)
	}

	if err:= os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(config.Folder, fmt.Sprintf("%s.log", timestamp))

	logfile, err := os.OpenFile(logFilePath, os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")
	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logfile), zapLvl),
	)

	zapLogger := zap.New(core, zap.AddCaller())

	return &Logger{
		Logger: zapLogger,
		file: logfile,
	}, nil
}

func (logger *Logger) With (field ...zap.Field) *Logger {
	return &Logger{
		Logger: logger.Logger.With(field...),
		file: logger.file,
	}

}

func (logger *Logger) Close() {
	if err := logger.file.Close(); err != nil {
		fmt.Println("failed to close logger:",err )
	}
}
