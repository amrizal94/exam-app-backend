package controllers

import (
	"net/http"
	"strings"

	"github.com/amrizal94/exam-app-backend/app"
	"github.com/amrizal94/exam-app-backend/helper"
	"github.com/amrizal94/exam-app-backend/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	type userValidation struct {
		Name     string `valid:"required"`
		Username string `valid:"required"`
		Email    string `valid:"email,required"`
		Password string `valid:"required,minstringlength(6)~password: must be at least 6 characters"`
	}
	return func(ctx *gin.Context) {

		// get request body
		request := app.User{}
		if err := ctx.ShouldBind(&request); err != nil {
			helper.ErrJSON(ctx, http.StatusBadRequest, err.Error())
		}

		// put data request body into struct userValidation for validation
		userValidate := &userValidation{
			Name:     request.Name,
			Username: request.Username,
			Email:    request.Email,
			Password: request.Password,
		}
		_, err := govalidator.ValidateStruct(userValidate)
		if err != nil {
			message := strings.Replace(err.Error(), ";", ", ", -1)
			helper.ErrJSON(ctx, http.StatusBadRequest, message)
			return
		}

		// prepare to create a new user
		hash, err := helper.HashPassword(request.Password)
		if err != nil {
			helper.ErrJSON(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		code := uuid.New()
		createUser := models.User{
			Name:     request.Name,
			Username: request.Username,
			Password: hash,
			Email:    request.Email,
			Code:     code,
		}

		// save the user in the database
		result := db.Create(&createUser)
		if result.Error != nil {
			err := result.Error.Error()
			if strings.Contains(err, "Error 1062 (23000)") {
				err = "username or email has been registered"
				helper.ErrJSON(ctx, http.StatusBadRequest, err)
				return
			}
			helper.ErrJSON(ctx, http.StatusInternalServerError, err)
			return
		}

		if result.RowsAffected == 0 {
			err := "database has not been changed. Perhaps the query is incorrect"
			helper.ErrJSON(ctx, http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "user has been registered",
		})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// get request body
		request := app.User{}
		if err := ctx.ShouldBind(&request); err != nil {
			helper.ErrJSON(ctx, http.StatusBadRequest, err.Error())
		}

		// check username and get password hashed from database
		password := request.Password
		user := models.User{}
		result := db.Select("id", "password").Where("username = ?", request.Username).Find(&user)
		if result.RowsAffected == 0 {
			err := "user not registered"
			helper.ErrJSON(ctx, http.StatusUnauthorized, err)
			return
		}

		// check password
		if !helper.CheckPasswordHash(password, user.Password) {
			helper.ErrJSON(ctx, http.StatusUnauthorized, "wrong password")
			return
		}

		// generate token for jwt authorization
		token, err := helper.GenerateToken(user.ID)
		if err != nil {
			helper.ErrJSON(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  token,
		})

	}
}
