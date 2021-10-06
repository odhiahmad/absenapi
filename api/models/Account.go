package models

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Username  string     `gorm:"unique;not null; size: 255" json:"username"`
	Password  string     `gorm:"not null; size: 50" json:"password"`
	IsActive  bool       `gorm:"not null; column:is_active"`
	UserID    uuid.UUID  `gorm:"null"`
	User      *User       `gorm:"foreignKey:UserID;references:ID;not null" json:"user"`
	Wallet    Wallet     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

func (a *Account) Prepare() error {
	hashedPassword, err := Hash(a.Password)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	a.IsActive = true
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Validate(action string) error {
	switch strings.ToLower(action) {
	case "forgot":
		if a.Password == "" {
			return errors.New("Minimum eight characters, at least one letter and one number")
		}
		return nil
	default:
		if a.Username == "" {
			return errors.New("Required username")
		}
		if a.Password == "" {
			return errors.New("Minimum eight characters, at least one letter and one number")
		}
		return nil
	}
}
