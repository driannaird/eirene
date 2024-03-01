package util

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
)

func InstallDepedency(req entity.Module) (*helper.ResponseModule, error) {
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

func DeleteDepedency(req entity.Module) (*helper.ResponseModule, error) {
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

func UpdatePackage(req entity.Module) error {
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

func check_os(os string) string {
	os_release, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		log.Printf("cannot read file because :%s", err.Error())
	}

	contain_os := strings.Contains(string(os_release), os)
	if contain_os {
		return os
	}

	return "Sorry your os not support"
}
