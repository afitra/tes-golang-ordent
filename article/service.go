package article

type Service interface {
	CreateArticle(input ArticleDataInput) (Article, error)
	GetArticleByID(input ArticleParamInput) (Article, error)
	UpdateArticleData(inputID ArticleParamInput, inputData ArticleDataInput) (Article, error)
	DestroyArticle(inputID ArticleParamInput) error
	GetAllArticleData() ([]Article, error)
	CostomFind(userId int, articleId ArticleParamInput) (Article, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateArticle(input ArticleDataInput) (Article, error) {

	article := Article{}

	article.User_id = input.User_id
	article.Title = input.Title
	article.Description = input.Description

	newArticle, err := s.repository.createArticle(article)

	if err != nil {
		return newArticle, err
	}
	return newArticle, nil

}

func (s *service) GetArticleByID(input ArticleParamInput) (Article, error) {

	article, err := s.repository.FindByID(input.ID)

	if err != nil {

		return article, err
	}
	return article, nil
}

func (s *service) UpdateArticleData(input ArticleParamInput, inputData ArticleDataInput) (Article, error) {

	article, err := s.repository.FindByID(input.ID)
	if err != nil {

		return article, err

	}

	article.User_id = inputData.User_id
	article.Title = inputData.Title
	article.Description = inputData.Description

	updateArticle, err := s.repository.Update(article)

	if err != nil {
		return updateArticle, err
	}
	return updateArticle, nil

}

func (s *service) DestroyArticle(input ArticleParamInput) error {
	err := s.repository.Destroy(input.ID)

	if err != nil {

		return err
	}
	return nil
}
func (s *service) GetAllArticleData() ([]Article, error) {
	datas, err := s.repository.FindAll()
	if err != nil {
		return datas, err
	}
	return datas, nil
}

func (s *service) CostomFind(userId int, articleId ArticleParamInput) (Article, error) {
	article, err := s.repository.CostomFind(userId, articleId)

	if err != nil {

		return article, err
	}
	return article, nil
}
