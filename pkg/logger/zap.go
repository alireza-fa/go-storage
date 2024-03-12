package logger

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var zapLogger *zap.SugaredLogger

type ZapLog struct {
	zap *zap.SugaredLogger
	cfg *Config
}

var zapLogLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func NewZap(cfg *Config) *ZapLog {
	log := &ZapLog{cfg: cfg}
	log.Init()
	return log
}

func (log *ZapLog) getLogLevel() zapcore.Level {
	level, exists := zapLogLevelMapping[log.cfg.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (log *ZapLog) Init() {
	once.Do(func() {
		fileName := fmt.Sprintf("%s%s-%s.%s", log.cfg.FilePath, time.Now().Format("2006-01-02"), uuid.New(), "log")
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    1,
			MaxAge:     20,
			LocalTime:  true,
			MaxBackups: 5,
			Compress:   true,
		})

		conf := zap.NewProductionEncoderConfig()
		conf.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(conf),
			w,
			log.getLogLevel(),
		)

		logger := zap.New(core, zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel),
		).Sugar()

		logger = logger.With("AppName", "Ghofle", "LoggerName", "Zaplog")

		zapLogger = logger
	})

	log.zap = zapLogger
}

func logParamsToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(extra))

	for key, value := range extra {
		params = append(params, string(key))
		params = append(params, value)
	}

	return params
}

func prepareLogInfo(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub

	return logParamsToZapParams(extra)
}

func (log *ZapLog) Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	log.zap.Debugw(message, params...)
}

func (log *ZapLog) Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	log.zap.Infow(message, params...)
}

func (log *ZapLog) Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	log.zap.Warnw(message, params...)
}

func (log *ZapLog) Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	log.zap.Errorw(message, params...)
}

func (log *ZapLog) Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	log.zap.Fatalw(message, params...)
}
