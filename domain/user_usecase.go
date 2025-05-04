package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type UserUseCase interface {
	Login(username, password string) (*entity.User, error)
	Register(user *entity.User) (*entity.User, error)
}
