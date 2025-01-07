package repositories

import (
	"errors"
	"github.com/Trend20/go-shoppers-api/internal/modules/users/models"
	"github.com/Trend20/go-shoppers-api/pkg/db"
)

type UserRepository struct{}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, db.DB.Error) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(id uint, updatedUser *models.User) error {
	var existingUser = models.User{}
	if err := db.DB.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, db.DB.Error) {
			return errors.New("user not found")
		}
		return err
	}
	if err := db.DB.Model(&existingUser).Updates(updatedUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, db.DB.Error) {
			return errors.New("user not found")
		}
		return err
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Authenticate(email, hashedPassword string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ? AND password = ?", email, hashedPassword).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
