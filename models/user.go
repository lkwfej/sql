package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(50);unique;not null" json:"username"`
	//Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
