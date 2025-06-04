package entity

type DosenProfile struct {
	UserID int    `gorm:"column:user_id" json:"user_id"`
	Nama   string `gorm:"column:nama" json:"nama"`
}

func (DosenProfile) TableName() string {
	return "profile_dosen"
}
