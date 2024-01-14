package main

import (
	"github.com/amrizal94/exam-app-backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func routes(router *gin.Engine, db *gorm.DB) {

	// grouping router version 1
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/register", controllers.Register(db))
			users.POST("/login", controllers.Login(db))
		}
	}
}
