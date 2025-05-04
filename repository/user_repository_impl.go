package repository

import (
	"github.com/indraalfauzan/monitoring_skripsi_golang/domain"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func (u *userRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return &entity.User{}, err
	}
	// Ambil kembali user lengkap dengan Role-nya
	if err := u.db.Preload("Role").First(user, user.ID).Error; err != nil {
		// Gagal preload role, tapi user tetap bisa dikembalikan
		return user, nil
	}
	return user, nil

}

// GetUserByID implements domain.UserRepository.
func (u *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Preload("Role").Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// Create implements domain.UserRepository.
