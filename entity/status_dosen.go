package entity

type StatusDosen struct {
	ID          int    `gorm:"column:id;primaryKey" json:"id"`
	StatusDosen string `gorm:"column:status_dosen" json:"status_dosen"`
}

func (StatusDosen) TableName() string {
	return "status_dosen"
}
