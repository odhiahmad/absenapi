package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Wallet struct {
	ID            uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique;index" json:"wallet_id"`
	Balance       uint64          `gorm:"not null;" `
	AccountID     uuid.UUID       `gorm:" null"`
	Account       *Account        `gorm:"foreignKey:AccountID;references:ID" `
	WalletHistory []WalletHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"wallet_history"`
	CreatedAt     time.Time       `gorm:"column:created_at;default:CURRENT_TIMESTAMP" `
	UpdatedAt     time.Time       `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" `
	DeletedAt     *time.Time      `gorm:"column:deleted_at" sql:"index"`
}

func (u *Wallet) Prepare() {
	u.UpdatedAt = time.Now()
}
