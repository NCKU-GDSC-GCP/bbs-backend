package model

import (
	"bbs_backend/entity"
	"bbs_backend/loaders"
	"fmt"
)

func NewTopicModel() TopicModel {
	return &topicModel{}
}

type TopicModel interface {
	GetAll() (*[]entity.Topic, error)
	Get(name string) (*entity.Topic, error)
	Create(topic *entity.Topic) (err error)
	Update(topic *entity.Topic) (err error)
	Delete(name string) error
}

type topicModel struct {
}

func (r *topicModel) Create(topic *entity.Topic) (err error) {
	result := loaders.DB.Debug().Create(topic)
	err = result.Error
	if result.Error != nil {
		fmt.Println("Create failed")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}
	return
}

func (r *topicModel) Get(name string) (*entity.Topic, error) {
	topic := &entity.Topic{}
	err := loaders.DB.First(topic, "name = ?", name).Error
	if err != nil {
		return nil, err
	} else {
		return topic, nil
	}
}

func (r *topicModel) GetAll() (*[]entity.Topic, error) {
	topics := &[]entity.Topic{}
	err := loaders.DB.Find(topics).Error
	if err != nil {
		return nil, err
	} else {
		return topics, nil
	}
}

func (r *topicModel) Update(topic *entity.Topic) (err error) {
	result := loaders.DB.Save(&topic)
	err = result.Error
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
		err = fmt.Errorf("RowsAffected Number fault")
	}
	return
}

func (r *topicModel) Delete(name string) error {
	res := loaders.DB.Delete(&entity.Topic{}, "name = ?", name)
	if res.RowsAffected < 1 {
		return fmt.Errorf("topic=%s cannot be deleted because it doesn't exist", name)
	}

	return nil
}
