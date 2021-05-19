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
	"github.com/go-playground/validator/v10"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)
var (
	ctx context.Context
    validate *validator.Validate
)


type User struct {
	Name  string `validate:"required"` //非空
	Age   uint8  `validate:"gte=0,lte=130"` //  0<=Age<=130
	Email string `validate:"required,email"` //非空，email格式
	//dive关键字代表 进入到嵌套结构体进行判断
	Address []*Address `validate:"dive"` //  可以拥有多个地址
}
type Address struct {
	Province string `validate:"required"` //非空
	City     string `validate:"required"` //非空
	Phone    string `validate:"numeric,len=11"` //数字类型，长度为11
}


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
// @termsOfService https://github.com/miraclewong/go-blog-service
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


	//validate = validator.New() //初始化（赋值）
	//validateStruct()           //结构体校验
	//validateVariable()         //变量校验
}


func validateStruct() {
	address := Address{
		Province: "重庆",
		City:     "重庆",
		Phone:    "13366663333x",
	}
	user := User{
		Name:  "江洲",
		Age:   23,
		Email: "jz@163.com",
		Address: []*Address{&address},
	}
	err := validate.Struct(user)
	if err != nil {
		//断言为：validator.ValidationErrors，类型为：[]FieldError
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("Namespace:", e.Namespace())
			fmt.Println("Field:", e.Field())
			fmt.Println("StructNamespace:", e.StructNamespace())
			fmt.Println("StructField:", e.StructField())
			fmt.Println("Tag:", e.Tag())
			fmt.Println("ActualTag:", e.ActualTag())
			fmt.Println("Kind:", e.Kind())
			fmt.Println("Type:", e.Type())
			fmt.Println("Value:", e.Value())
			fmt.Println("Param:", e.Param())
			fmt.Println()
		}

		fmt.Println("结构体输入数据类型错误！")
		return
	} else {
		fmt.Println("结构体校验通过")
	}
}
//变量校验
func validateVariable() {
	myEmail := "123@qq.com" //邮箱地址：xx@xx.com
	err := validate.Var(myEmail, "required,email")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("变量校验通过！")
	}
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