package service

import (
	"context"
	"log"

	"github.com/jinzhu/copier"
	"github.com/rulanugrh/eirene/src/helper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type IMetric interface {
	GetMetric() (*[]helper.Metric, error)
	GetTracer() trace.Tracer
}

type imetrict struct{}

func NewMetric() IMetric {
	return &imetrict{}
}

func (m *imetrict) GetMetric() (*[]helper.Metric, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, helper.InternalServerError(err.Error())
	}

	var response []helper.Metric

	traceProsessor := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(traceProsessor),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				string(semconv.URLFullKey),
				semconv.ServiceNameKey.String("Eirene"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	var metrc helper.Metric
	copier.Copy(&metrc, tp)

	response = append(response, metrc)
	return &response, nil
}

func (m *imetrict) GetTracer() trace.Tracer {
	tracer := otel.Tracer("eirene-server")
	return tracer
}
