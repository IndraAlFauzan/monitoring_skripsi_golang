package repository

import (
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"gorm.io/gorm"
)

type profilTARepository struct {
	db *gorm.DB
}

func NewProfilTARepository(db *gorm.DB) domain.ProfileTARepository {
	return &profilTARepository{db}
}

func (r *profilTARepository) CreateProfilTA(profil *entity.ProfielTA) (*entity.ProfielTA, error) {
	if err := r.db.Create(profil).Error; err != nil {
		return nil, err
	}
	return profil, nil
}

func (r *profilTARepository) CreatePembimbing(dosens []*entity.ProfileTADosen) error {
	return r.db.Create(&dosens).Error
}
