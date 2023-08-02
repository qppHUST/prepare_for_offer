package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qppHUST/prepareForOffer/web/auth"
	"github.com/qppHUST/prepareForOffer/web/model"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/auth", token)
		api.POST("/tokenCheck", tokenCheck)
	}

	// router.Handle("/hello",http.FileServer(root FileSystem))

}

func token(ctx *gin.Context) {
	var tokenRequest model.TokenRequest
	if err := ctx.ShouldBind(&tokenRequest); err != nil {
		log.Fatal(err)
	}
	token, error := auth.GetJWT(tokenRequest)
	if error != nil {
		log.Fatalln("GetJWT error", error)
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func tokenCheck(ctx *gin.Context) {
	var tokenCheckRequest model.TokenCheckRequest
	if err := ctx.ShouldBind(&tokenCheckRequest); err != nil {
		log.Fatalln(err)
	}
	claims, err := auth.ParseJWT(tokenCheckRequest.Token)
	if err != nil {
		log.Fatalln("parse error", err)
	}
	ctx.JSON(200, claims)
}
