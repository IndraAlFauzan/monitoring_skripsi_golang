package entity

type MahasiswaProfile struct {
	UserID    int    `gorm:"primaryKey"`
	NIM       string `json:"nim"`
	Nama      string `json:"nama"`
	NoHP      string `json:"no_hp"`
	PhotoPath string `json:"photo_path"`
}
