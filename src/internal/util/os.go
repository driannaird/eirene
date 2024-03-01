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
		response := install_package(req, "/bin/apt")
		return response, nil
	case "centos":
		response := install_package(req, "/bin/yum")
		return response, nil
	default:
		return nil, helper.BadRequest("Sorry your os not support")
	}
}

func install_package(req entity.Module, command string) *helper.ResponseModule {
	for _, pkg := range req.Package {
		err := exec.Command(command, pkg).Err
		if err != nil {
			log.Printf("Something error when install package :%s", err.Error())
			return &helper.ResponseModule{
				Package: nil,
				Message: "sorry package not installed",
			}
		}
	}

	return &helper.ResponseModule{
		Package: req.Package,
		Message: "Package success installed",
	}
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
