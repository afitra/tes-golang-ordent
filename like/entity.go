package like

import (
	"tes-golang-ordent/article"
	"tes-golang-ordent/user"
)

type Like struct {
	ID         int `gorm:"primaryKey"`
	User_id    int `gorm:"not null"`
	Article_id int `gorm:"not null"`
	User       user.User
	Article    article.Article
}
