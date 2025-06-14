package userService

import (
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type UserService interface {
	CreateUser(email string, password string) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id string, email string, password string) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(email string, password string) (User, error) {
	id := types.UUID(uuid.New())

	newUser := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id string) (User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return User{}, err
	}
	return s.repo.GetUserByID(uid)
}

func (s *userService) UpdateUser(id string, email string, password string) (User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return User{}, err
	}

	existingUser, err := s.repo.GetUserByID(uid)
	if err != nil {
		return User{}, err
	}

	existingUser.Email = email
	existingUser.Password = password

	if err := s.repo.UpdateUser(existingUser); err != nil {
		return User{}, err
	}

	return existingUser, nil
}

func (s *userService) DeleteUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUser(uid)
}
