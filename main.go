package main

import (
	"fmt"
	_ "github.com/MiracleWong/go-blog-service/docs"
	"github.com/MiracleWong/go-blog-service/global"
	"github.com/MiracleWong/go-blog-service/internal/model"
	"github.com/MiracleWong/go-blog-service/internal/routers"
	"github.com/MiracleWong/go-blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	//err := setupSetting()
	//if err != nil {
	//	log.Fatal("init.setupSetting err: %v ", err)
	//}
	err := setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v ", err)
	}
}

//func init() {
//	viper.SetConfigName("config1") // 读取json配置文件
//	viper.AddConfigPath(".")      // 设置配置文件和可执行二进制文件在用一个目录
//	viper.SetConfigType("json")
//	if err := viper.ReadInConfig(); err != nil {
//		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
//			// Config file not found; ignore error if desired
//			log.Println("no such config file")
//		} else {
//			// Config file was found but another error was produced
//			log.Println("read config error")
//		}
//		log.Fatal(err) // 读取配置文件失败致命错误
//	}
//}

func main() {
	//gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10* time.Second,
		WriteTimeout:  10* time.Second,
		MaxHeaderBytes: 1 << 20,
		//Addr: ":8080",
		//Handler: router,
		//ReadTimeout: global.ServerSetting.ReadTimeout * time.Second,
		//WriteTimeout:  global.ServerSetting.WriteTimeout * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//fmt.Println("获取配置文件的port", viper.GetInt("port"))
	//fmt.Println("获取配置文件的mysql.url", viper.GetString(`mysql.url`))
	//fmt.Println("获取配置文件的mysql.username", viper.GetString(`mysql.username`))
	//fmt.Println("获取配置文件的mysql.password", viper.GetString(`mysql.password`))
	//fmt.Println("获取配置文件的redis", viper.GetStringSlice("redis"))
	//fmt.Println("获取配置文件的smtp", viper.GetStringMap("smtp"))
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if  err != nil {
		return err
	}
	err = s.ReadSection("Sever", &global.ServerSetting)
	fmt.Println(&global.ServerSetting)
	if  err != nil {
		return err
	}
	err = s.ReadSection("Sever", &global.AppSetting)
	if  err != nil {
		return err
	}
	err = s.ReadSection("Sever", &global.DataSetting)
	if  err != nil {
		return err
	}
	//global.ServerSetting.ReadTimeout *= time.Second
	//global.ServerSetting.WriteTimeout *= time.Second
	return err
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine()
	//global.DBEngine, err = model.NewDBEngine(global.DataSetting)
	if err != nil {
		return err
	}
	return nil
}