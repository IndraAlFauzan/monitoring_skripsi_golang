package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type MahasiswaProfileUseCase interface {
	CreateProfile(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error)
	GetProfile(userID int) (*entity.MahasiswaProfile, error)
	UpdateProfile(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error)
}
