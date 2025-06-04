package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type ProfileTARepository interface {
	CreateProfilTA(profil *entity.ProfielTA) (*entity.ProfielTA, error)
	CreatePembimbing(dosens []*entity.ProfileTADosen) error
	GetProfileTA(userID int) (*entity.ProfileTAWithPembimbing, error)
}
