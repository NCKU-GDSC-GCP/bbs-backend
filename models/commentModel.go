package model

import (
	"bbs_backend/entity"
	"bbs_backend/loaders"
	"fmt"
)

func NewCommentModel() CommentModel {
	return &commentModel{}
}

type CommentModel interface {
	GetAllByPost(pid uint) *[]entity.Comment
	GetAllByUser(uid uint) *[]entity.Comment
	Get(id uint) (*entity.Comment, error)
	Create(comment *entity.Comment) (err error)
	Update(comment *entity.Comment) (err error)
	Delete(id uint) error
}

type commentModel struct {
}

func (r *commentModel) Create(comment *entity.Comment) (err error) {
	result := loaders.DB.Debug().Create(comment)
	err = result.Error
	if result.Error != nil {
		fmt.Println("Create failed")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}
	return
}

func (r *commentModel) Get(id uint) (*entity.Comment, error) {
	comment := &entity.Comment{}
	err := loaders.DB.First(comment, "id = ?", id).Error
	if err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}

func (r *commentModel) GetAllByPost(pid uint) *[]entity.Comment {
	comments := make([]entity.Comment, 0)
	loaders.DB.Where("post_id = ?", pid).Find(&comments)
	return &comments
}

func (r *commentModel) GetAllByUser(uid uint) *[]entity.Comment {
	comments := make([]entity.Comment, 0)
	loaders.DB.Where("author_id = ?", uid).Find(&comments)
	return &comments
}

func (r *commentModel) Update(comment *entity.Comment) (err error) {
	result := loaders.DB.Save(&comment)
	err = result.Error
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
		err = fmt.Errorf("RowsAffected Number fault")
	}
	return
}

func (r *commentModel) Delete(id uint) error {
	res := loaders.DB.Delete(&entity.Comment{}, "id = ?", id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("topic=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}
