package article

import "gorm.io/gorm"

type Repository interface {
	createArticle(article Article) (Article, error)
	FindByID(ID int) (Article, error)
	Update(article Article) (Article, error)
	Destroy(ID int) error
	FindAll() ([]Article, error)
	CostomFind(userId int, articleId ArticleParamInput) (Article, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) createArticle(article Article) (Article, error) {

	err := r.db.Create(&article).Error

	if err != nil {

		return article, err
	}
	return article, nil

}
func (r *repository) FindByID(ID int) (Article, error) {
	var article Article

	err := r.db.Where("id = ?", ID).Find(&article).Error

	if err != nil {

		return article, err
	}

	return article, nil

}
func (r *repository) Update(article Article) (Article, error) {

	err := r.db.Save(&article).Error

	if err != nil {

		return article, err
	}
	return article, nil

}
func (r *repository) Destroy(ID int) error {
	var article Article

	err := r.db.Where("id = ?", ID).Delete(&article).Error
	if err != nil {

		return err
	}

	return nil

}
func (r *repository) FindAll() ([]Article, error) {
	var data []Article
	err := r.db.Preload("User").Find(&data).Error

	if err != nil {

		return data, err
	}
	return data, nil
}

func (r *repository) CostomFind(userId int, articleId ArticleParamInput) (Article, error) {
	var article Article

	err := r.db.Where("id = ? AND user_id = ? ", articleId.ID, userId).Find(&article).Error

	if err != nil {

		return article, err
	}

	return article, nil
}
