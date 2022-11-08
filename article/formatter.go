package article

type ArticleFormatter struct {
	ID          int    `json:"id"`
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func FormatArticle(article Article) ArticleFormatter {

	articleFormatter := ArticleFormatter{}
	articleFormatter.ID = article.ID

	articleFormatter.User_id = article.User_id
	articleFormatter.Title = article.Title
	articleFormatter.Description = article.Description

	return articleFormatter
}

type ArticleIncludeFormatter struct {
	ID          int           `json:"id"`
	User_id     int           `json:"user_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	User        UserFormatter `json:"user"`
}

type UserFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatIncludeArticle(article Article) ArticleIncludeFormatter {

	articleIncludeFormatter := ArticleIncludeFormatter{}
	articleIncludeFormatter.ID = article.ID
	articleIncludeFormatter.User_id = article.User_id
	articleIncludeFormatter.Title = article.Title
	articleIncludeFormatter.Description = article.Description

	user := article.User

	articleUser := UserFormatter{}
	articleUser.Name = user.Name
	articleUser.Email = user.Email

	// articleLike := like.FormatIncludeLike()

	articleIncludeFormatter.User = articleUser

	return articleIncludeFormatter
}
func FormatAllArticle(inputData []Article) []ArticleIncludeFormatter {
	if len(inputData) == 0 {
		return []ArticleIncludeFormatter{}
	}

	allFormatter := []ArticleIncludeFormatter{}
	for _, item := range inputData {

		temp := FormatIncludeArticle(item)

		allFormatter = append(allFormatter, temp)
	}

	return allFormatter
}
