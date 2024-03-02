package helper

import (
	"time"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tr "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

type MailList struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	From       string     `json:"from"`
	To         string     `json:"to"`
	Message    string     `json:"msg"`
	Attachment Attachment `json:"attachment"`
}

type Attachment struct {
	File  string `json:"file"`
	Link  string `json:"link"`
	Video string `json:"video"`
}

type MailCreate struct {
	Title   string `json:"title"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"msg"`
}

type MailUpdate struct {
	Title   string `json:"title"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"msg"`
	Type    string `json:"type"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type UserLogin struct {
	Token string `json:"token"`
}

type ResponseModule struct {
	Package interface{} `json:"package"`
	Message string      `json:"message"`
}

type Image struct {
	File string `json:"file"`
	Link string `json:"link"`
}

type File struct {
	File string `json:"file"`
	Link string `json:"link"`
}

type Metric struct {
	embedded.TracerProvider `json:"embeded"`

	NamedTracer map[instrumentation.Scope]*tr.Tracer

	Sampler     trace.Sampler
	IDGenerator trace.IDGenerator
	SpanLimits  trace.SpanLimits
	Resources   *resource.Resource
}

type DockerImage struct {
	ID          string   `json:"id"`
	Tag         []string `json:"tag"`
	Created     int64    `json:"created"`
	Size        int64    `json:"size"`
	VirtualSize int64    `json:"virtual_size"`
	Labels      map[string]string
}

type InspectDockerImage struct {
	ID            string    `json:"id"`
	Tag           []string  `json:"tag"`
	Created       time.Time `json:"created"`
	Container     string    `json:"container"`
	OS            string    `json:"os"`
	Architecture  string    `json:"architecture"`
	Size          int64     `json:"size"`
	VirtualSize   int64     `json:"virtual_size"`
	Author        string    `json:"author"`
	DockerVersion string    `json:"docker_version"`
}
