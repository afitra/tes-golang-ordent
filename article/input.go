package article

type ArticleDataInput struct {
	User_id     int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type ArticleParamInput struct {
	ID int `uri:"id" binding:"required"`
}
