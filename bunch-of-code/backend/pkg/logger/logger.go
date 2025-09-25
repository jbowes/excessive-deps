package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, kv ...any)
	Error(msg string, kv ...any)
}

type zapLogger struct{ l *zap.SugaredLogger }

func New(env string) Logger {
	var cfg zap.Config
	if env == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	lg, _ := cfg.Build()
	return &zapLogger{l: lg.Sugar()}
}

func (z *zapLogger) Info(msg string, kv ...any) {
	z.l.Infow(msg, pairs(kv...)...)
}

func (z *zapLogger) Error(msg string, kv ...any) {
	z.l.Errorw(msg, pairs(kv...)...)
}

func pairs(kv ...any) []any {
	// ensure even length
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}
	// stringify keys
	for i := 0; i < len(kv); i += 2 {
		if _, ok := kv[i].(string); !ok {
			kv[i] = fmt.Sprintf("%v", kv[i])
		}
	}
	return kv
}
