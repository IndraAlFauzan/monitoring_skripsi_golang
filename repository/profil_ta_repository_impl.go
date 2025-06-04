package repository

import (
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"gorm.io/gorm"
)

type profilTARepository struct {
	db *gorm.DB
}

func (r *profilTARepository) GetProfileTA(userID int) (*entity.ProfileTAWithPembimbing, error) {
	var profile entity.ProfielTA
	if err := r.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	// Ambil pembimbing 1
	var pembimbing1 struct {
		Nama string
	}
	r.db.Table("profile_ta_dosen").
		Select("dosen_profiles.nama").
		Joins("JOIN dosen_profiles ON profile_ta_dosen.dosen_id = dosen_profiles.user_id").
		Where("profile_ta_dosen.profil_ta_id = ? AND profile_ta_dosen.status_dosen_id = ?", profile.IDProfile, 1).
		Scan(&pembimbing1)

	// Ambil pembimbing 2
	var pembimbing2 struct {
		Nama string
	}
	r.db.Table("profile_ta_dosen").
		Select("dosen_profiles.nama").
		Joins("JOIN dosen_profiles ON profile_ta_dosen.dosen_id = dosen_profiles.user_id").
		Where("profile_ta_dosen.profil_ta_id = ? AND profile_ta_dosen.status_dosen_id = ?", profile.IDProfile, 2).
		Scan(&pembimbing2)

	result := &entity.ProfileTAWithPembimbing{
		IDProfile:           profile.IDProfile,
		UserID:              profile.UserID,
		JudulSkripsi:        profile.JudulSkripsi,
		StatusBimbingan:     profile.StatusBimbingan,
		CreatedAt:           profile.CreatedAt,
		DosenPembimbingSatu: pembimbing1.Nama,
		DosenPembimbingDua:  pembimbing2.Nama,
	}

	return result, nil
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
