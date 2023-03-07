package model

import (
	"bbs_backend/entity"
	"bbs_backend/loaders"
	"fmt"
)

func NewFriendModel() FriendModel {
	return &friendModel{}
}

type FriendModel interface {
	Create(friend *entity.Friend) (err error)
	GetAll(uid uint) (*[]entity.Friend, error)
	Update(friend *entity.Friend) (err error)
	Delete(uid uint, tarUid uint) error
}

type friendModel struct {
}

func (m *friendModel) Create(friend *entity.Friend) (err error) {
	result := loaders.DB.Debug().Create(friend)
	err = result.Error
	if result.Error != nil {
		fmt.Println("Create failed")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}
	return
}

func (m *friendModel) GetAll(uid uint) (*[]entity.Friend, error) {
	friends := &[]entity.Friend{}
	err := loaders.DB.Where("uid_to = ? OR uid_from = ?", uid, uid).Find(friends).Error
	if err != nil {
		return nil, err
	} else {
		return friends, nil
	}
}

func (m *friendModel) Update(friend *entity.Friend) (err error) {
	result := loaders.DB.Save(&friend)
	err = result.Error
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
		err = fmt.Errorf("RowsAffected Number fault")
	}
	return
}

func (m *friendModel) Delete(uid uint, tarUid uint) error {
	res := loaders.DB.Where("(uid_to = ? AND uid_from = ?) OR (uid_to = ? AND uid_from = ?)", uid, tarUid, tarUid, uid).Delete(&entity.Friend{})
	if res.RowsAffected < 1 {
		return fmt.Errorf("friendId=%d cannot be deleted because it doesn't exist", tarUid)
	}

	return nil
}
