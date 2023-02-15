package entity

import (
	"time"
)

type Post struct {
	Id        uint   `gorm:"primary_key:auto_increment" json:"id"`
	Title     string `gorm:"type:varchar(255);not null;" json:"title"`
	Content   string `gorm:"type:varchar(255);" json:"content"`
	AuthorId  uint   `gorm:"primary_key" json:"author_id"`
	User      User   `gorm:"foreignKey:AuthorId"`
	TopicId   uint   `gorm:"primary_key" json:"topic_id"`
	Topic     Topic
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
