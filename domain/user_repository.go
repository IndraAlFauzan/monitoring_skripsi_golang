package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
