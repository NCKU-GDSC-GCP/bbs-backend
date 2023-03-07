package entity

import (
	"time"
)

type Comment struct {
	Id        uint      `gorm:"primary_key:auto_increment" json:"id"`
	Content   string    `gorm:"type:varchar(255);" json:"content"`
	AuthorId  uint      `gorm:"primary_key" json:"author_id"`
	User      *User     `gorm:"foreignKey:AuthorId" json:"author,omitempty"`
	PostId    uint      `gorm:"primary_key" json:"post_id"`
	Post      *Post     `json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
