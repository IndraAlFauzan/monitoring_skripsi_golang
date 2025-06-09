package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type ProfileTARepository interface {
	CreateProfilTA(profil *entity.ProfileTA) (*entity.ProfileTA, error)
	CreatePembimbing(dosens []*entity.ProfileTADosen) error
	GetProfileTA(userID int) (*entity.ProfileTAWithPembimbing, error)
	CheckProfileExist(userID int) (bool, error)
}
