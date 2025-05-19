package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type MahasiswaProfileRepository interface {
	Create(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error)
	GetByUserID(userID int) (*entity.MahasiswaProfile, error)
	Update(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error)
}
