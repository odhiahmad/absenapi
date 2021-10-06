package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type statusTransaction string

const (
	WITHDRAW statusTransaction = "WITHDRAW"
	PAYMENT	 statusTransaction = "PAYMENT"
	TOPUP 	 statusTransaction = "TOPUP"
)

type WalletHistory struct {
	ID              uuid.UUID         `gorm:"primary_key;unique;index;type:uuid;default:uuid_generate_v4()" `
	TransactionDate time.Time         `gorm:"not null;type:date"`
	Amount          uint64            `gorm:"not null" `
	TypeTransaction statusTransaction `gorm:"not null;" `
	WalletID        uuid.UUID         `gorm:"not null"`
	Wallet          *Wallet           `gorm:"foreignKey:WalletID;references:ID" `
	CreatedAt       time.Time         `gorm:"column:created_at;default:CURRENT_TIMESTAMP" `
	UpdatedAt       time.Time         `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" `
	DeletedAt       *time.Time        `gorm:"column:deleted_at" sql:"index"`
}

func (wh *WalletHistory) Prepare() {
	wh.CreatedAt = time.Now()
	wh.UpdatedAt = time.Now()
}
