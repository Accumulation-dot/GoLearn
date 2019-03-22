package main

import (
	"bitbucket.org/forfunn/GoLearn/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	//sessiondb := redis.New(service.Config{
	//	Network:service.DefaultRedisNetwork,
	//	Addr: service.DefaultRedisAddr,
	//	Password: "",
	//	Database: "",
	//	MaxIdle: 0,
	//	MaxActive:0,
	//	IdleTimeout:service.DefaultRedisIdleTimeout,
	//	Prefix:"", })
	//
	//
	//
	//defer sessiondb.Close()


	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "golang_" + defaultTableName
	}

	mysqldb, err := gorm.Open("mysql", "developer:ZXcV1357@$^*@tcp(192.168.44.152:3306)/mysql_go_test?&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println(err.Error())
	}

	//user := models.User{}

	mysqldb.AutoMigrate(&models.User{}, &models.HomeMenu{})

	defer mysqldb.Close()

	iris.RegisterOnInterrupt(func() {
		//sessiondb.Close()
		mysqldb.Close()
	})

	

	

	//session := sessions.New(sessions.Config{
	//	Cookie: "sessionscookieid",
	//	Expires: 45 * time.Minute, })

	//session.UseDatabase(sessiondb)




}

func RedisConfig(app *iris.Application)  {





}



