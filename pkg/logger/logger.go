package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var log = logrus.New()

type Config struct {
	LogLevel   string
	LogFile    string
	MaxSize    int  // максимальный размер в мегабайтах
	MaxBackups int  // максимальное количество файлов
	MaxAge     int  // максимальный возраст в днях
	Console    bool // вывод в консоль
	JSONFormat bool
}

type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Добавляем информацию о файле и строке
	if _, file, line, ok := runtime.Caller(6); ok {
		entry.Data["source"] = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return f.TextFormatter.Format(entry)
}

func Init(cfg Config) error {
	// Создаем все родительские директории
	logDir := filepath.Dir(cfg.LogFile)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории логов: %v", err)
	}

	// Проверяем права на запись в файл
	file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла логов: %v", err)
	}
	file.Close()

	// Настраиваем ротацию логов
	fileLogger := &lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   true,
	}

	// Настраиваем формат
	if cfg.JSONFormat {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		log.SetFormatter(&CustomFormatter{
			TextFormatter: logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339,
			},
		})
	}

	// Настраиваем вывод в оба места
	writers := []io.Writer{fileLogger}
	if cfg.Console {
		writers = append(writers, os.Stdout)
	}
	log.SetOutput(io.MultiWriter(writers...))

	// Устанавливаем уровень логирования
	if level, err := logrus.ParseLevel(cfg.LogLevel); err != nil {
		return fmt.Errorf("неверный уровень логирования: %v", err)
	} else {
		log.SetLevel(level)
	}

	// Пишем тестовый лог
	Info("Логгер инициализирован", map[string]interface{}{
		"logFile": cfg.LogFile,
		"level":   cfg.LogLevel,
	})

	return nil
}

// Методы логирования
func Info(msg string, fields map[string]interface{}) {
	if fields == nil {
		log.Info(msg)
	} else {
		log.WithFields(fields).Info(msg)
	}
}

func Error(err error, msg string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["error"] = err
	log.WithFields(fields).Error(msg)
}

func Debug(msg string, fields map[string]interface{}) {
	if fields == nil {
		log.Debug(msg)
	} else {
		log.WithFields(fields).Debug(msg)
	}
}

func GetLogger() *logrus.Logger {
	return log
}

func WithContext(ctx map[string]interface{}) *logrus.Entry {
	return log.WithFields(logrus.Fields(ctx))
}
