package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type UserUseCase interface {
	Login(username, password string) (*entity.User, error)
	RegisterMhs(user *entity.User) (*entity.User, error)
	RegisterUser(user *entity.User) (*entity.User, error)
}
