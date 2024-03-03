package helper

import (
	"time"

	docker "github.com/fsouza/go-dockerclient"
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

type Config struct {
	Hostname   string                   `json:"hostname,omitempty"`
	Domainname string                   `json:"domain_name,omitempty"`
	Image      string                   `json:"image,omitempty"`
	Tty        bool                     `json:"tty"`
	OpenStdin  bool                     `json:"bool"`
	Env        []string                 `json:"env"`
	Port       map[docker.Port]struct{} `json:"port"`
}

type Container struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Created      time.Time `json:"created"`
	Path         string    `json:"path"`
	HostnamePath string    `json:"hostname_path"`
	HostsPath    string    `json:"host_path"`
	Config       *Config   `json:"config"`
}

type Port struct {
	PrivatePort int64  `json:"private_port"`
	PublicPort  int64  `json:"public_port"`
	Type        string `json:"type"`
	IP          string `json:"ip"`
}

type ListContainer struct {
	ID      string `json:"id"`
	Image   string `json:"image"`
	Command string `json:"command"`
	Status  string `json:"status"`
	State   string `json:"state"`
	Created int64  `json:"created"`
	Ports   []Port `json:"ports"`
}

type InspectContainer struct {
	ID           string                   `json:"id"`
	Image        string                   `json:"image"`
	HostnamePath string                   `json:"hostname_path"`
	HostsPath    string                   `json:"hosts_path"`
	Name         string                   `json:"name"`
	OpenStdin    bool                     `json:"open_stdin"`
	TTY          bool                     `json:"tty"`
	Env          []string                 `json:"env"`
	Port         map[docker.Port]struct{} `json:"port"`
}

type HistoryImage struct {
	ID       string   `json:"id"`
	Tags     []string `json:"tags"`
	Created  int64    `json:"created"`
	CreateBy string   `json:"create_by"`
	Size     int64    `json:"size"`
	Comment  string   `json:"comment"`
}

type Volume struct {
	Name       string            `json:"name"`
	Driver     string            `json:"driver"`
	Labels     map[string]string `json:"label"`
	DriverOpts map[string]string `json:"driver_ops"`
}
