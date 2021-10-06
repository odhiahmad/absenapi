package models

import uuid "github.com/satori/go.uuid"

type Merchant struct {
	ID      uuid.UUID `gorm:"primary_key;unique;column:merchant_id"`
	Name    string    `gorm:"not null; size:255; column:name"`
	Address string    `gorm:"not null; size:255; column:address"`
}
