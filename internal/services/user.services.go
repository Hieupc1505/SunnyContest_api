package services

type IUserService interface {
	Register() error
	Login() error
}

type userService struct{}

func NewUserService() IUserService {
	return &userService{}
}

func (s *userService) Register() error {
	return nil
}

func (s *userService) Login() error {
	return nil
}
