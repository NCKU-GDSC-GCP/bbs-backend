package entity

import "time"

type Topic struct {
	Id          uint      `gorm:"primary_key:auto_increment" json:"id"`
	Name        string    `gorm:"primary_key;type:varchar(255);not null;unique" json:"name"`
	Description string    `gorm:"type:varchar(255);not null;" json:"description"`
	Cover       string    `gorm:"type:varchar(255);not null" json:"cover"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
