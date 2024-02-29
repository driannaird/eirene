package entity

import "gorm.io/gorm"

type Mail struct {
	gorm.Model
	Title      string     `json:"title" validate:"required"`
	Type       string     `json:"type"`
	UserEmail  string     `json:"user_email"`
	To         string     `json:"to" validate:"required"`
	Message    string     `json:"msg" validate:"required"`
	Attachment Attachment `json:"attachment" gorm:"polymorphic:Attach;"`
}

type Attachment struct {
	gorm.Model
	AttachID   uint8  `json:"attach_id"`
	AttachType string `json:"attach_type"`
	Link       string `json:"link"`
	File       string `json:"file"`
	Video      string `json:"video"`
	Code       string `json:"code"`
}
