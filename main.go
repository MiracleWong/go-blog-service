package main

import (
	"context"
	"fmt"
	"github.com/MiracleWong/go-blog-service/global"
	"github.com/MiracleWong/go-blog-service/internal/model"
	"github.com/MiracleWong/go-blog-service/internal/routers"
	"github.com/MiracleWong/go-blog-service/pkg/logger"
	"github.com/MiracleWong/go-blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)
var (
	ctx context.Context
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatal("init.setupSetting err: %v ", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v ", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v ", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com:MiracleWong/go-blog-service
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:  global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof(ctx,"%s: go-programming-tour-book/%s", "MiracleWong","go-blog-service")
	s.ListenAndServe()
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if  err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if  err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	fmt.Println(global.ServerSetting.ReadTimeout)
	fmt.Println(global.ServerSetting.WriteTimeout)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}