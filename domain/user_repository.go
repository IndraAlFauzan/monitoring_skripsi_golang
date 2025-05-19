package domain

import "github.com/indraalfauzan/monitoring_skripsi_golang/entity"

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByEmailandRole(email string, roleID int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
