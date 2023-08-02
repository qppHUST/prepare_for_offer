package model

type TokenRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type TokenCheckRequest struct {
	Token string `form:"token" binding:"required"`
}
