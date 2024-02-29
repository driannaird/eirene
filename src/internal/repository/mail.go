package repository

import (
	"github.com/rulanugrh/eirene/src/internal/entity"
	"gorm.io/gorm"
)

type MailRepository interface {
	Inbox(user_email string) ([]entity.Mail, error)
	Sent(req entity.Mail) (*entity.Mail, error)
	Starred(user_email string) ([]entity.Mail, error)
	Archived(user_email string) ([]entity.Mail, error)
	Update(id uint, model entity.Mail) (*entity.Mail, error)
	Delete(id uint) error
}

type mailrepo struct {
	db *gorm.DB
}

func NewMailRepository(db *gorm.DB) MailRepository {
	return &mailrepo{
		db: db,
	}
}

func (repo *mailrepo) Inbox(user_email string) ([]entity.Mail, error) {
	var listEmail []entity.Mail
	err := repo.db.Where("user_email = ?", user_email).Find(&listEmail).Error
	if err != nil {
		return nil, err
	}

	return listEmail, nil
}

func (repo *mailrepo) Sent(req entity.Mail) (*entity.Mail, error) {
	err := repo.db.Create(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (repo *mailrepo) Starred(user_email string) ([]entity.Mail, error) {
	var listEmail []entity.Mail
	err := repo.db.Where("type = ?", "starred").Where("user_email = ?", user_email).Find(&listEmail).Error

	if err != nil {
		return nil, err
	}

	return listEmail, nil
}

func (repo *mailrepo) Archived(user_email string) ([]entity.Mail, error) {
	var listEmail []entity.Mail
	err := repo.db.Where("type = ?", "archice").Where("user_email = ?", user_email).Find(&listEmail).Error

	if err != nil {
		return nil, err
	}

	return listEmail, nil
}

func (repo *mailrepo) Update(id uint, model entity.Mail) (*entity.Mail, error) {
	var update entity.Mail
	err := repo.db.Model(&model).Where("id = ?", id).Updates(&update).Error
	if err != nil {
		return nil, err
	}

	return &update, nil
}

func (repo *mailrepo) Delete(id uint) error {
	var mail entity.Mail
	err := repo.db.Where("id = ?", id).Delete(&mail).Error
	if err != nil {
		return err
	}

	return nil
}
