package like

type LikeDataInput struct {
	User_id    int `json:"user_id" binding:"required"`
	Article_id int `uri:"id" binding:"required"`
}

type LikeParamInput struct {
	ID int `uri:"id" binding:"required"`
}
