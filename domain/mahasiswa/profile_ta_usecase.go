package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type ProfileTAUseCase interface {
	AjukanTA(userID int, judul string, dosen1ID int, dosen2ID int) (*entity.ProfielTA, error)
	GetProfileTA(userID int) (*entity.ProfileTAWithPembimbing, error)
}
