package modles

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `gorm:"unique"`
	Email    string    `gorm:"unique"`
	Password string
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`
}
