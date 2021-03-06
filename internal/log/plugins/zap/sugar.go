package zap

import (
	"atomic/internal/log/tracer"
	"context"
	"fmt"
	"go.uber.org/zap"
)

func getCtxFields(args ...interface{}) []zap.Field {
	//判断是否有context
	if len(args) > 0 {
		if ctx, ok := args[len(args)-1].(context.Context); ok {
			return getTraceField(ctx)
		}
	}
	return []zap.Field{}
}

//
func (l *Log) Debug(s string, args ...interface{}) {
	l.logger.Debug(s, getCtxFields(args...)...)
}
func (l *Log) Info(s string, args ...interface{}) {
	l.logger.Info(s, getCtxFields(args...)...)
}
func (l *Log) Warn(s string, args ...interface{}) {
	l.logger.Warn(s, getCtxFields(args...)...)
}
func (l *Log) Error(s error, args ...interface{}) {
	l.logger.Error(s.Error(), getCtxFields(args...)...)
}
func (l *Log) Panic(s string, args ...interface{}) {
	l.logger.Panic(s, getCtxFields(args...)...)
}
func (l *Log) Fatal(s string, args ...interface{}) {
	l.logger.Fatal(s, getCtxFields(args...)...)
}

//

//判断其他类型--start
func getOtherFields(format string, args ...interface{}) (string, []zap.Field) {
	//判断是否有context
	l := len(args)
	if l > 0 {
		if ctx, ok := args[l-1].(context.Context); ok {
			return fmt.Sprintf(format, args[:l-1]...), getTraceField(ctx)
		} else {
			return fmt.Sprintf(format, args[:l]...), []zap.Field{}
		}
	}
	return format, []zap.Field{}
}

func (l *Log) Debugf(format string, args ...interface{}) {
	s, f := getOtherFields(format, args...)
	l.logger.Debug(s, f...)
}

func (l *Log) Infof(format string, args ...interface{}) {
	s, f := getOtherFields(format, args...)
	l.logger.Info(s, f...)
}

func (l *Log) Warnf(format string, args ...interface{}) {
	s, f := getOtherFields(format, args...)
	l.logger.Warn(s, f...)
}

func (l *Log) Errorf(format error, args ...interface{}) {
	s, f := getOtherFields(format.Error(), args...)
	l.logger.Error(s, f...)
}

func (l *Log) Panicf(format string, args ...interface{}) {
	s, f := getOtherFields(format, args...)
	l.logger.Panic(s, f...)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
	s, f := getOtherFields(format, args...)
	l.logger.Fatal(s, f...)
}

// 获取链路跟踪添加列
func getTraceField(ctx context.Context) []zap.Field {
	fm := tracer.GetTraceInfo(ctx)
	zf := make([]zap.Field, 0)
	for k, v := range fm {
		zf = append(zf, zap.String(k, v))
	}
	return zf
}
