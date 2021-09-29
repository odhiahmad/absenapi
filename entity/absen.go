package entity

import "time"

type User struct {
	Id        uint64 `gorm:"primary_key:auto_increament" json:"id"`
	UserId    uint64 `gorm:"not null" json:"-"`
	User      User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Tanggal time.time `gorm:"not null" json:"tanggal"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
}
