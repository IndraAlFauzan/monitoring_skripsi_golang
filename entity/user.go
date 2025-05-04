package entity

// call the role
type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:100;not null;unique"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:255;not null"`
	RoleID   int
	Role     Role `gorm:"foreignKey:RoleID"`
}
