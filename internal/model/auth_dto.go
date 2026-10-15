package model

type LoginReq struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginRes struct {
	Id string `json:"id"`
}
