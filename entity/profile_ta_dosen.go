package entity

type ProfileTADosen struct {
	ID            int `gorm:"primaryKey" json:"id"`
	ProfilTAID    int `gorm:"column:profil_ta_id" json:"profil_ta_id"`
	DosenID       int `gorm:"column:dosen_id" json:"dosen_id"`
	StatusDosenID int `gorm:"column:status_dosen_id" json:"status_dosen_id"` // 1/2
}

func (ProfileTADosen) TableName() string {
	return "profile_ta_dosen"
}
