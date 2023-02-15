package entity

import (
	"time"
)

type Friend struct {
	UidFrom   uint      `gorm:"primary_key;type:int" json:"uidFrom"`
	UserFrom  User      `gorm:"foreignKey:UidFrom"`
	UidTo     uint      `gorm:"primary_key;type:int" json:"uidTo"`
	UserTo    User      `gorm:"foreignKey:UidTo"`
	Status    uint      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
