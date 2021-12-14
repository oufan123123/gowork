package main

import (
	"fmt"
	"log"
	"modules/config"
	"modules/handler"
	"modules/pojo"
	"modules/respository"
	"modules/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB          *gorm.DB
	UserHandler handler.UserHandler
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("init database")
	conf := &pojo.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		"Local")
	DB, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatal("connect database fail: %v\n", err)
	}
	DB.SingularTable(true)
	fmt.Println("database init success!!!")

}

func initHandler() {
	UserHandler = handler.UserHandler{
		USC: service.UserService{
			Repo: &respository.UserRespository{
				DB: DB,
			},
		}}
}

func init() {
	initViper()
	initDB()
	initHandler()
}
