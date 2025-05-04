package repository

import (
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	roles := []entity.Role{
		{Name: "admin"},
		{Name: "dosen"},
		{Name: "mahasiswa"},
	}
	for _, role := range roles {
		var count int64
		db.Model(&entity.Role{}).Where("name = ?", role.Name).Count(&count)
		if count == 0 {
			db.Create(&role)
		}
	}
	return nil
}
