package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type ModuleService interface {
	Install(req entity.Module) (*helper.ResponseModule, error)
	Update(req entity.Module) error
	Delete(req entity.Module) (*helper.ResponseModule, error)
	AddSSHKey(req entity.SSHKey) error
}

type moduleservice struct {
	mod      util.ModuleInstall
	validate middleware.IValidate
	trace    trace.Tracer
}

func NewModuleService(mod util.ModuleInstall, validate middleware.IValidate) ModuleService {
	return &moduleservice{mod: mod, validate: validate, trace: otel.Tracer("module-service")}
}

func (mod *moduleservice) Install(req entity.Module) (*helper.ResponseModule, error) {
	span, err := util.Tracer(mod.trace, "installPackage")
	if err != nil {
		return nil, err
	}

	defer span.End()

	err = mod.validate.Validate(req)
	if err != nil {
		return nil, mod.validate.ValidationMessage(err)
	}

	response, err := mod.mod.InstallDepedency(req)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	return response, nil
}

func (mod *moduleservice) Update(req entity.Module) error {
	span, err := util.Tracer(mod.trace, "updatePackage")
	if err != nil {
		return err
	}

	defer span.End()

	err = mod.validate.Validate(req)
	if err != nil {
		return mod.validate.ValidationMessage(err)
	}

	err = mod.mod.UpdatePackage(req)
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	return helper.Success("success update package", nil)
}

func (mod *moduleservice) Delete(req entity.Module) (*helper.ResponseModule, error) {
	span, err := util.Tracer(mod.trace, "deletePackage")
	if err != nil {
		return nil, err
	}

	defer span.End()

	err = mod.validate.Validate(req)
	if err != nil {
		return nil, mod.validate.ValidationMessage(err)
	}

	response, err := mod.mod.DeleteDepedency(req)
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	return response, nil
}

func (mod *moduleservice) AddSSHKey(req entity.SSHKey) error {
	span, err := util.Tracer(mod.trace, "addSSHKey")
	if err != nil {
		return err
	}

	defer span.End()

	err = mod.validate.Validate(req)
	if err != nil {
		return mod.validate.ValidationMessage(err)
	}

	err = mod.mod.AddSSHKey(req)
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	return nil
}
