package entity

import "time"

type User struct {
	Id        uint64 `gorm:"primary_key:auto_increament" json:"id"`
	Username  string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Password  string `gorm:"->;<-;not null" json:"-"`
	Token     string `gorm:"-" json:"token,omitempty"`
	Status    uint8  `gorm:"not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
}
