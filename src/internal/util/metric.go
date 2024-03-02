package util

import (
	"context"

	"github.com/rulanugrh/eirene/src/helper"
	"go.opentelemetry.io/otel/attribute"
	tr "go.opentelemetry.io/otel/trace"
)

func Tracer(trace tr.Tracer, spanName string) (tr.Span, error) {
	ctx, span := trace.Start(context.Background(), spanName)
	if ctx.Err() != nil {
		return nil, helper.InternalServerError(ctx.Err().Error())
	}
	return span, nil
}

func TracerWithAttribute(trace tr.Tracer, spanName string, attr string) (tr.Span, error) {
	ctx, span := trace.Start(context.Background(), spanName, tr.WithAttributes(attribute.String("id", attr)))
	if ctx.Err() != nil {
		return nil, helper.InternalServerError(ctx.Err().Error())
	}
	return span, nil
}
