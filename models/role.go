package models

type Role struct {
	RoleID   uint   `gorm:"primaryKey"`
	RoleName string `gorm:"size:50;not null"`
}
