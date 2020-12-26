package services

// UserService is an interface which defines
// all the services that are done by `User`
type UserService interface {
	CreateNewUser(email, password, category string) error
	GetUser(ID int) error
}

type userService struct{}

// NewUserService creates an instace of a userService
// struct which will act as the service class for user
func NewUserService() UserService {
	return &userService{}
}

func (u *userService) CreateNewUser(email, password, category string) error {
	return nil
}

func (u *userService) GetUser(ID int) error {
	return nil
}
