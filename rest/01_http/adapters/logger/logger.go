package logger

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

type Logger struct{}

func (l Logger) Debug(ctx context.Context, msg string) {
	logger.WithContext(ctx).Debug(msg)
}

func (l Logger) Info(ctx context.Context, msg string) {
	logger.WithContext(ctx).Info(msg)
}

func (l Logger) Error(ctx context.Context, msg string) {
	logger.WithContext(ctx).Error(msg)
}
