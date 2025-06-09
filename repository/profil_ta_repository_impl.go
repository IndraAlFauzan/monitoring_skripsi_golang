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

func (r *profilTARepository) GetProfileTA(userID int) (*entity.ProfileTAWithPembimbing, error) {
	var profile entity.ProfileTA
	if err := r.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	err := r.db.
		Preload("Pembimbings.Dosen").
		Preload("Pembimbings.Status").
		Where("user_id = ?", userID).
		First(&profile).Error

	if err != nil {
		return nil, err
	}

	var pembimbing1, pembimbing2 string
	for _, d := range profile.Pembimbings {
		if d.StatusDosenID == 1 {
			pembimbing1 = d.Dosen.Nama
		} else if d.StatusDosenID == 2 {
			pembimbing2 = d.Dosen.Nama
		}
	}

	result := &entity.ProfileTAWithPembimbing{
		IDProfile:           profile.IDProfile,
		UserID:              profile.UserID,
		JudulSkripsi:        profile.JudulSkripsi,
		StatusBimbingan:     profile.StatusBimbingan,
		CreatedAt:           profile.CreatedAt,
		DosenPembimbingSatu: pembimbing1,
		DosenPembimbingDua:  pembimbing2,
	}

	return result, nil
}

func (r *profilTARepository) CreateProfilTA(profil *entity.ProfileTA) (*entity.ProfileTA, error) {
	if err := r.db.Create(profil).Error; err != nil {
		return nil, err
	}
	return profil, nil
}

func (r *profilTARepository) CreatePembimbing(dosens []*entity.ProfileTADosen) error {
	return r.db.Create(&dosens).Error
}

func (r *profilTARepository) CheckProfileExist(userID int) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.ProfileTA{}).
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
