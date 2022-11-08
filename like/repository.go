package like

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateLike(like Like) (Like, error)
	FindLike(like LikeDataInput) (Like, error)
	DeleteLike(like Like) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) CreateLike(like Like) (Like, error) {

	err := r.db.Create(&like).Error

	if err != nil {

		return like, err
	}
	return like, nil

}

func (r *repository) FindLike(like LikeDataInput) (Like, error) {
	var data Like
	err := r.db.Where("user_id = ? AND article_id = ?", like.User_id, like.Article_id).Find(&data).Error

	if err != nil {

		return data, err
	}

	return data, nil
}

func (r *repository) DeleteLike(like Like) error {

	var data Like

	err := r.db.Where("user_id = ? AND article_id = ?", like.User_id, like.Article_id).Delete(&data).Error
	if err != nil {

		return err
	}

	return nil
}
