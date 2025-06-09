package entity

type DosenProfile struct {
	UserID    int    `gorm:"column:user_id;primaryKey" json:"user_id"`
	Nama      string `gorm:"column:nama" json:"nama"`
	Keahlian  string `gorm:"column:keahlian" json:"keahlian"`
	NIDN      string `gorm:"column:nidn" json:"nidn"`
	NoHP      string `gorm:"column:no_hp" json:"no_hp"`
	PhotoPath string `gorm:"column:photo_path" json:"photo_path"`
}

func (DosenProfile) TableName() string {
	return "dosen_profiles"
}
