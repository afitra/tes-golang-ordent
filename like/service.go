package like

type Service interface {
	CreateLike(input LikeDataInput) (Like, error)
	FindLike(input LikeDataInput) (Like, error)
	DeleteLike(input LikeDataInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateLike(input LikeDataInput) (Like, error) {

	like := Like{}

	like.User_id = input.User_id
	like.Article_id = input.Article_id

	newLike, err := s.repository.CreateLike(like)

	if err != nil {
		return newLike, err
	}
	return newLike, nil

}

func (s *service) FindLike(input LikeDataInput) (Like, error) {
	like, err := s.repository.FindLike(input)

	if err != nil {

		return like, err
	}
	return like, nil
}

func (s *service) DeleteLike(input LikeDataInput) error {

	like := Like{}

	like.User_id = input.User_id
	like.Article_id = input.Article_id

	err := s.repository.DeleteLike(like)

	if err != nil {
		return err
	}
	return nil

}
