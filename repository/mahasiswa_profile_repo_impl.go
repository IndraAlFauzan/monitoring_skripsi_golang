package repository

import (
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"gorm.io/gorm"
)

type mahasiswaProfileRepo struct {
	db *gorm.DB
}

func NewMahasiswaProfileRepository(db *gorm.DB) domain.MahasiswaProfileRepository {
	return &mahasiswaProfileRepo{db}
}

func (r *mahasiswaProfileRepo) Create(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error) {
	if err := r.db.Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}
func (r *mahasiswaProfileRepo) GetByUserID(userID int) (*entity.MahasiswaProfile, error) {
	var profile entity.MahasiswaProfile
	if err := r.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}
func (r *mahasiswaProfileRepo) Update(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error) {
	// Update the profile in the database use where clause
	if err := r.db.Where("user_id = ?", profile.UserID).Updates(profile).Error; err != nil {
		return nil, err
	}
	// Return the updated profile
	return profile, nil

}
