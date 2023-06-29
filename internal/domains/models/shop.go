package models

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:VARCHAR(255);unique" json:"name"`
	Description string         `gorm:"type:TEXT" json:"description"`
	Location    string         `gorm:"type:VARCHAR(255)" json:"location"`
	TelShop     string         `gorm:"type:VARCHAR(255)" json:"telShop"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	OwnerID int  `gorm:"not null" json:"owner_id"`
	Owner   User `gorm:"foreignKey:OwnerID" json:"owner"`
}
