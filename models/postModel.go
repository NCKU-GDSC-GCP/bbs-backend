package model

import (
	"bbs_backend/entity"
	"bbs_backend/loaders"
	"fmt"
)

func NewPostModel() PostModel {
	return &postModel{}
}

type PostModel interface {
	GetAllByTopic(topicId uint) *[]entity.Post
	GetAllByUser(uid uint) *[]entity.Post
	Get(pid uint) (*entity.Post, error)
	Create(post *entity.Post) (err error)
	Update(post *entity.Post) (err error)
	Delete(pid uint) error
}

type postModel struct {
}

func (r *postModel) Create(post *entity.Post) (err error) {
	result := loaders.DB.Debug().Create(post)
	err = result.Error
	if result.Error != nil {
		fmt.Println("Create failed")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}
	return
}

func (r *postModel) Get(pid uint) (*entity.Post, error) {
	post := entity.Post{}
	err := loaders.DB.First(&post, pid).Error
	if err != nil {
		return nil, err
	} else {
		return &post, nil
	}
}

func (r *postModel) GetAllByTopic(topicId uint) *[]entity.Post {
	posts := make([]entity.Post, 0)
	loaders.DB.Where("topic_id = ?", topicId).Find(&posts)
	return &posts
}

func (r *postModel) GetAllByUser(uid uint) *[]entity.Post {
	posts := make([]entity.Post, 0)
	loaders.DB.Where("author_id = ?", uid).Find(&posts)
	return &posts
}

func (r *postModel) Update(post *entity.Post) (err error) {
	result := loaders.DB.Save(&post)
	err = result.Error
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
		err = fmt.Errorf("RowsAffected Number fault")
	}
	return
}

func (r *postModel) Delete(pid uint) error {
	res := loaders.DB.Delete(&entity.Post{}, "id = ?", pid)
	if res.RowsAffected < 1 {
		return fmt.Errorf("post=%d cannot be deleted because it doesn't exist", pid)
	}

	return nil
}
