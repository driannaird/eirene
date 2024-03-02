package util

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
)

type ModuleInstall interface {
	InstallDepedency(req entity.Module) (*helper.ResponseModule, error)
	DeleteDepedency(req entity.Module) (*helper.ResponseModule, error)
	UpdatePackage(req entity.Module) error
	AddSSHKey(req entity.SSHKey) error
}

type mod struct {
}

func NewModuleInstall() ModuleInstall {
	return &mod{}
}

func (m *mod) InstallDepedency(req entity.Module) (*helper.ResponseModule, error) {
	os := check_os(req.OS)
	switch os {
	case "ubuntu":
		response := install_package(req, "/bin/apt install")
		return response, nil
	case "debian":
		response := install_package(req, "/bin/apt install")
		return response, nil
	case "centos":
		response := install_package(req, "/bin/yum install")
		return response, nil
	default:
		return nil, helper.BadRequest("Sorry your os not support")
	}
}

func (m *mod) DeleteDepedency(req entity.Module) (*helper.ResponseModule, error) {
	os := check_os(req.OS)
	switch os {
	case "ubuntu":
		response := install_package(req, "/bin/apt purge")
		return response, nil
	case "debian":
		response := install_package(req, "/bin/apt purge")
		return response, nil
	case "centos":
		response := install_package(req, "/bin/yum purge")
		return response, nil
	default:
		return nil, helper.BadRequest("Sorry your os not support")
	}
}

func (m *mod) UpdatePackage(req entity.Module) error {
	os := check_os(req.OS)
	switch os {
	case "ubuntu":
		err := run_exec("/bin/apt")
		return err
	case "debian":
		err := run_exec("/bin/apt")
		return err
	case "centos":
		err := run_exec("/bin/yum")
		return err
	default:
		return helper.BadRequest("Sorry your os not support")
	}
}

func (m *mod) AddSSHKey(req entity.SSHKey) error {
	f, err := os.OpenFile("~/.ssh/authorized_keys", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return helper.BadRequest("sorry your file is not found")
	}

	defer f.Close()

	_, err = f.WriteString(req.Key)
	if err != nil {
		return helper.InternalServerError("sorry cannot insert ssh-keygen")
	}

	return nil
}

func install_package(req entity.Module, command string) *helper.ResponseModule {
	err := exec.Command(command, req.Package...).Err
	if err != nil {
		log.Printf("Something error when install package :%s", err.Error())
		return &helper.ResponseModule{
			Package: nil,
			Message: "sorry package not installed",
		}
	}

	return &helper.ResponseModule{
		Package: req.Package,
		Message: "Package success installed",
	}
}

func run_exec(command string) error {
	err := exec.Command(command, "update").Err
	if err != nil {
		return helper.BadRequest("Sorry yu cant running this command")
	}

	return helper.Success("success update server", nil)
}

func check_os(_os string) string {
	os_release, err := os.ReadFile("/etc/os-release")
	if err != nil {
		log.Printf("cannot read file because :%s", err.Error())
	}

	contain_os := strings.Contains(string(os_release), _os)
	if contain_os {
		return _os
	}

	return "Sorry your os not support"
}
