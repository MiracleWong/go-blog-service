package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MiracleWong/go-blog-service/internal/model"
	"github.com/MiracleWong/go-blog-service/internal/routers"
	"github.com/gin-gonic/gin"

	"github.com/MiracleWong/go-blog-service/global"
	"github.com/MiracleWong/go-blog-service/pkg/setting"

	_ "github.com/MiracleWong/go-blog-service/docs"
)

// 初始化配置读取，自动执行

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v ", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupSetting err: %v ", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book

func main() {

	//dsn := "root:jxkj@123@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("连接成功，db名称: ", db)

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DataSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	fmt.Println("settings : ")
	fmt.Println(global.ServerSetting)
	fmt.Println(global.DataSetting)
	fmt.Println(global.AppSetting)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
