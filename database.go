package main

import (
	"fmt"

	"github.com/amrizal94/exam-app-backend/helper"
	"github.com/amrizal94/exam-app-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func initDB() *gorm.DB {
	dbConfig := &DBConfig{
		User:     helper.Getenv("DBUser", "root"),
		Password: helper.Getenv("DBPass", ""),
		Host:     helper.Getenv("DBHost", "localhost"),
		Port:     helper.Getenv("DBPort", "3306"),
		Name:     helper.Getenv("DBName", "test"),
	}

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + dbConfig.Name + ";")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// migrate schema
	db.AutoMigrate(&models.User{}, &models.Role{})
	roles := []*models.Role{
		{Name: "admin"},
		{Name: "member"},
	}

	db.Create(roles)

	return db
}
