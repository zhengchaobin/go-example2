package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		//context.String(http.StatusOK, "hello world")

		context.JSON(http.StatusOK, gin.H{
			"name" : "zcb",
			"age" : 12,
			"gender" : "男",
		})
	})

	router.Run(":8088")

}
