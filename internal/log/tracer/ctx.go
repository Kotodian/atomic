package tracer

import (
	"context"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"strings"
)

const (
	DefaultTraceIdName  = "trace_id"
	DefaultSpanIdName   = "span_id"
	DefaultParentIdName = "parent_id"
	DefaultSeparator    = ":"
)

var LogTraceKey struct{}

var NoTraceInfo = errors.New("no trace info")

func decodeTracer(ctx context.Context) ([]string, error) {
	s := make([]string, 0, 4)
	if val, ok := ctx.Value(LogTraceKey).(string); ok {
		s = strings.Split(val, DefaultSeparator)
	} else {
		span := opentracing.SpanFromContext(ctx)
		s = strings.Split(fmt.Sprintf("%v", span), DefaultSeparator)
	}

	if len(s) >= 3 {
		return s, nil
	}
	return []string{}, NoTraceInfo
}

func GetTraceInfo(ctx context.Context) map[string]string {
	s, err := decodeTracer(ctx)
	trace := make(map[string]string)
	if err != nil {
		return trace
	}
	trace[DefaultTraceIdName] = s[0]
	trace[DefaultSpanIdName] = s[1]
	trace[DefaultParentIdName] = s[2]

	return trace
}
