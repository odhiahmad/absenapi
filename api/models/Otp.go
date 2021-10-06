package models

import uuid "github.com/satori/go.uuid"

type Otp struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	OtpCode string    `gorm:"not null; size:255; column:otp_code"`
	UserID  uuid.UUID `gorm:"null"`
	User    *User     `gorm:"foreignKey:UserID;references:ID"`
}

type OtpRequest struct {
	UserId  string
	OtpCode string
}
