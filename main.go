package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// init database
	db := initDB()

	// server router using gin engine default
	router := gin.Default()

	//call router function
	routes(router, db)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "get request")
	})
	router.Run()
}
