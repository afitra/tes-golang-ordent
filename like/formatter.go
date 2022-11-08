package like

import (
	"tes-golang-ordent/article"
)

type LikeFormatter struct {
	ID         int `json:"id"`
	User_id    int `json:"user_id"`
	Article_id int `json:"article_id"`
}

func FormatLike(like Like) LikeFormatter {

	likeFormatter := LikeFormatter{}
	likeFormatter.ID = like.ID

	likeFormatter.User_id = like.User_id
	likeFormatter.Article_id = like.Article_id

	return likeFormatter
}

type LikeIncludeFormatter struct {
	ID int `json:"id"`
	// User_id    int                             `json:"user_id"`
	// Artilce_id int                             `json:"like_id"`
	User    UserFormatter                   `json:"user"`
	Article article.ArticleIncludeFormatter `json:"article"`
}
type UserFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatIncludeLike(like Like) LikeIncludeFormatter {

	likeIncludeFormatter := LikeIncludeFormatter{}
	likeIncludeFormatter.ID = like.ID

	user := UserFormatter{}
	user.Name = user.Name
	user.Email = user.Email

	article := article.FormatIncludeArticle(like.Article)

	likeIncludeFormatter.User = user
	likeIncludeFormatter.Article = article

	return likeIncludeFormatter
}
func FormatAllLike(inputData []Like) []LikeIncludeFormatter {
	if len(inputData) == 0 {
		return []LikeIncludeFormatter{}
	}

	allFormatter := []LikeIncludeFormatter{}
	for _, item := range inputData {

		temp := FormatIncludeLike(item)

		allFormatter = append(allFormatter, temp)
	}

	return allFormatter
}
