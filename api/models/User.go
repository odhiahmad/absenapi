package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string     `gorm:"size:255;null"`
	PhoneNumber uint64     `gorm:"null;unique" `
	Email       string     `gorm:"size:100;null;unique" `
	BirthPlace  string     `gorm:"size:100;null" `
	BirthDate   string     `gorm:"type:date;null" `
	MotherName  string     `gorm:"size: 255;null" `
	BankAccount string     `gorm:"not null" `
	Account     Account    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}


type UserRequest struct {
	Username    string
	Password    string
	Name        string
	PhoneNumber uint64
	Email       string
	BirthPlace  string
	BirthDate   string
	MotherName  string
	BankAccount string
}

type UserResponse struct {
	ID          uuid.UUID
	Name        string
	Username    string
	PhoneNumber uint64
	Email       string
	BirthPlace  string
	BirthDate   string
	MotherName  string
	BankAccount string
	Account     *Account
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (u *User) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UserRequest) Validate(action string) error {
	//PhoneNumberStr := fmt.Sprint(u.PhoneNumber)
	//fmt.Println("phone string",PhoneNumberStr)
	switch strings.ToLower(action) {
	case "update":
		if u.PhoneNumber == 0 {
			return errors.New("Required Phone Number")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil

	default:
		if u.PhoneNumber == 0 {
			return errors.New("Required Phone Number")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil
	}

}
