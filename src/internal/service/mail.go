package service

import (
	"github.com/rulanugrh/eirene/src/helper"
	"github.com/rulanugrh/eirene/src/internal/entity"
	"github.com/rulanugrh/eirene/src/internal/middleware"
	"github.com/rulanugrh/eirene/src/internal/repository"
	"github.com/rulanugrh/eirene/src/internal/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type MailService interface {
	Inbox(user_email string) (*[]helper.MailList, error)
	Sent(req entity.Mail) (*helper.MailCreate, error)
	Starred(user_email string) (*[]helper.MailList, error)
	Archived(user_email string) (*[]helper.MailList, error)
	Update(id uint, model entity.Mail) (*helper.MailUpdate, error)
	Delete(id uint) error
}

type mailservice struct {
	repo     repository.MailRepository
	validate middleware.IValidate
	trace    trace.Tracer
}

func NewMailService(repo repository.MailRepository, validate middleware.IValidate) MailService {
	return &mailservice{
		repo:     repo,
		validate: validate,
		trace:    otel.Tracer("mail-service"),
	}
}

func (m *mailservice) Inbox(user_email string) (*[]helper.MailList, error) {
	span, err := util.Tracer(m.trace, "inboxMail")
	if err != nil {
		return nil, err
	}

	defer span.End()

	data, err := m.repo.Inbox(user_email)
	if err != nil {
		return nil, helper.NotFound("email not found")
	}

	var response []helper.MailList
	for _, dt := range data {
		result := helper.MailList{
			ID:      dt.ID,
			Title:   dt.Title,
			From:    dt.UserEmail,
			To:      dt.To,
			Message: dt.Message,
			Attachment: helper.Attachment{
				File:  dt.Attachment.File,
				Video: dt.Attachment.Video,
				Link:  dt.Attachment.Link,
			},
		}

		response = append(response, result)
	}

	return &response, nil
}

func (m *mailservice) Sent(req entity.Mail) (*helper.MailCreate, error) {
	span, err := util.Tracer(m.trace, "sentMail")
	if err != nil {
		return nil, err
	}

	defer span.End()

	err = m.validate.Validate(req)
	if err != nil {
		return nil, m.validate.ValidationMessage(err)
	}

	data, err := m.repo.Sent(req)
	if err != nil {
		return nil, helper.InternalServerError("cannot create email")
	}

	response := helper.MailCreate{
		Title:   data.Title,
		Message: data.Message,
		To:      data.To,
		From:    data.UserEmail,
	}

	return &response, nil
}

func (m *mailservice) Starred(user_email string) (*[]helper.MailList, error) {
	span, err := util.Tracer(m.trace, "starredMail")
	if err != nil {
		return nil, err
	}

	defer span.End()

	data, err := m.repo.Starred(user_email)
	if err != nil {
		return nil, helper.NotFound("email not found")
	}

	var response []helper.MailList
	for _, dt := range data {
		result := helper.MailList{
			ID:      dt.ID,
			Title:   dt.Title,
			From:    dt.UserEmail,
			To:      dt.To,
			Message: dt.Message,
			Attachment: helper.Attachment{
				File:  dt.Attachment.File,
				Video: dt.Attachment.Video,
				Link:  dt.Attachment.Link,
			},
		}

		response = append(response, result)
	}

	return &response, nil
}

func (m *mailservice) Archived(user_email string) (*[]helper.MailList, error) {
	span, err := util.Tracer(m.trace, "archiveMail")
	if err != nil {
		return nil, err
	}

	defer span.End()

	data, err := m.repo.Archived(user_email)
	if err != nil {
		return nil, helper.NotFound("email not found")
	}

	var response []helper.MailList
	for _, dt := range data {
		result := helper.MailList{
			ID:      dt.ID,
			Title:   dt.Title,
			From:    dt.UserEmail,
			To:      dt.To,
			Message: dt.Message,
			Attachment: helper.Attachment{
				File:  dt.Attachment.File,
				Video: dt.Attachment.Video,
				Link:  dt.Attachment.Link,
			},
		}

		response = append(response, result)
	}

	return &response, nil
}

func (m *mailservice) Update(id uint, model entity.Mail) (*helper.MailUpdate, error) {
	span, err := util.Tracer(m.trace, "updateMail")
	if err != nil {
		return nil, err
	}

	defer span.End()

	data, err := m.repo.Update(id, model)
	if err != nil {
		return nil, helper.InternalServerError("cannot update type email")
	}

	response := helper.MailUpdate{
		Title:   data.Title,
		Message: data.Message,
		To:      data.To,
		From:    data.UserEmail,
		Type:    model.Type,
	}

	return &response, nil
}

func (m *mailservice) Delete(id uint) error {
	span, err := util.Tracer(m.trace, "deleteMail")
	if err != nil {
		return err
	}

	defer span.End()

	if err = m.repo.Delete(id); err != nil {
		return helper.InternalServerError("cannot delete email")
	}

	return nil
}
