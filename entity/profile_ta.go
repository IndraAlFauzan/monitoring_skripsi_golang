package entity

import "time"

type ProfielTA struct {
	IDProfile       int       `gorm:"column:id_profil;primaryKey" json:"id_profil"`
	UserID          int       `gorm:"column:user_id" json:"user_id"`
	JudulSkripsi    string    `gorm:"column:judul_skripsi" json:"judul_skripsi"`
	StatusBimbingan string    `gorm:"column:status_bimbingan" json:"status_bimbingan"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProfielTA) TableName() string {
	return "profile_ta"
}
