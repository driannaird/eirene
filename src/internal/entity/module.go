package entity

type Module struct {
	OS      string   `json:"os" form:"os" validate:"required"`
	Package []string `json:"package" form:"package"`
}

type SSHKey struct {
	Key string `json:"ssh_key" form:"ssh_key" validate:"required"`
}
