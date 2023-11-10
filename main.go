package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MiracleWong/go-blog-service/internal/routers"
	"github.com/gin-gonic/gin"

	"github.com/MiracleWong/go-blog-service/global"
	"github.com/MiracleWong/go-blog-service/pkg/setting"
)

// 初始化配置读取，自动执行

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v ", err)
	}
}

func main() {
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
