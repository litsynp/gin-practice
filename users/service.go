package users

type UserService interface {
	CreateUser(user User) (User, error)
	FindUserById(id int64) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUserById(id int64) error
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(user User) (User, error) {
	err := s.userRepository.Create(&user)
	return user, err
}

func (s *userService) FindUserById(id int64) (User, error) {
	user, err := s.userRepository.FindById(id)
	return user, err
}

func (s *userService) UpdateUser(user User) (User, error) {
	err := s.userRepository.Update(&user)
	return user, err
}

func (s *userService) DeleteUserById(id int64) error {
	err := s.userRepository.DeleteById(id)
	return err
}
