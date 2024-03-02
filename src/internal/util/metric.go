package util

import (
	"context"

	"github.com/rulanugrh/eirene/src/helper"
	"go.opentelemetry.io/otel/trace"
)

func Tracer(trace trace.Tracer, spanName string) (trace.Span, error) {
	ctx, span := trace.Start(context.Background(), spanName)
	if ctx.Err() != nil {
		return nil, helper.InternalServerError(ctx.Err().Error())
	}
	return span, nil
}
