package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qppHUST/prepareForOffer/web/router"
)

func main() {
	runServer()
}

func runServer() {
	handler := gin.Default()
	router.InitRouter(handler)
	server := http.Server{
		Addr:    ":9999",
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
