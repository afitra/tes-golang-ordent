package article

import (
	"tes-golang-ordent/user"
)

type Article struct {
	ID          int    `gorm:"primaryKey"`
	User_id     int    `gorm:"not null"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	User        user.User
}
