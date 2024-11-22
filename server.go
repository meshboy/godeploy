package main

import(
	"os"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()
	server.SetTrustedProxies(nil)
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"message": "Hello World!",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}


	server.Run(fmt.Sprintf(":%s", port))
}