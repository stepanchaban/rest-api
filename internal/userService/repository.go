package userService

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error
	GetAllUsers() ([]User, error)
	GetUserByID(id uuid.UUID) (User, error)
	UpdateUser(user User) error
	DeleteUser(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (r *userRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uuid.UUID) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}