package services

import (
	"errors"
	"github.com/Trend20/go-shoppers-api/internal/modules/users/models"
	"github.com/Trend20/go-shoppers-api/internal/modules/users/repositories"
	"github.com/Trend20/go-shoppers-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(UserRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

// Register creates a new user with a hashed password
func (s *UserService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepo.Create(user)
}

// Login authenticates a user and returns a JWT token if valid
func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAll()
}

// GetUserByID retrieves a single user by their ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.GetByID(id)
}

// update user
func (s *UserService) UpdateUser(id uint, updatedUser *models.User) error {
	return s.UserRepo.Update(id, updatedUser)
}

// delete user
func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepo.Delete(id)
}
