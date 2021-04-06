package main

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	service2 "user/domain/service"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"user/domain/repository"
	"user/handler"
	user "user/proto/user"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	db,err :=gorm.Open("mysql","root:root@tcp(10.4.7.71:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)

	//只执行一次，数据表初始化
	//rp:=repository.NewUserRepository(db)
	//rp.InitTable()
	err = repository.NewUserRepository(db).InitTable()
	if err !=nil {
		log.Error(err)
	}

	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	//注册Handler
	err = user.RegisterUserHandler(srv.Server(),&handler.User{UserDataService:userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
