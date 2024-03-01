package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/util"
)

type ModuleService interface {
	Install(req entity.Module) (*helper.ResponseModule, error)
	Update(req entity.Module) error
	Delete(req entity.Module) (*helper.ResponseModule, error)
	AddSSHKey(req entity.SSHKey) error
}

type moduleservice struct {
	mod util.ModuleInstall
}

func NewModuleService(mod util.ModuleInstall) ModuleService {
	return &moduleservice{mod: mod}
}

func (mod *moduleservice) Install(req entity.Module) (*helper.ResponseModule, error) {
	response, err := mod.mod.InstallDepedency(req)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	return response, nil
}

func (mod *moduleservice) Update(req entity.Module) error {
	err := mod.mod.UpdatePackage(req)
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	return helper.Success("success update package", nil)
}

func (mod *moduleservice) Delete(req entity.Module) (*helper.ResponseModule, error) {
	response, err := mod.mod.DeleteDepedency(req)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	return response, nil
}

func (mod *moduleservice) AddSSHKey(req entity.SSHKey) error {
	err := mod.mod.AddSSHKey(req)
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	return nil
}
