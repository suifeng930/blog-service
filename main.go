package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/go-programming-tour-book/blog-service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port string
	runMode string
	config string
)


// 主要作用是控制应用程序的初始化流程，起到将配置文件内容映射到应用配置结构体中
func init() {

	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err:%v", err)

	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)

	}
	// 新增路由监控
	err = setupTracer()
	if err != nil {

		log.Fatalf("init.setupTracer err: %v", err)
	}
	err = setupFlag()
	if err != nil {
		log.Fatalf("init setupFlag() err:%v", err)

	}
}

// @title   blog -service
// @version 1.0.0
// @description go blog demo
// @termsOfService https://github.com/go-programming-tour-book
func main() {

	//global.Logger.Infof(context.Background(),"%s:  go-programing-tour-book/%s", "xiaoma", "blog-service")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println("s.ListenAndServe()", err)

	}

}

func setupSetting() error {


	// 新增 命令行启动参数配置
	setting, err := setting.NewSetting(strings.Split(config,",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err

	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err

	}
	//log.Println("global.ServerSetting:",global.ServerSetting)
	//log.Println("global.AppSetting:",global.AppSetting)
	//log.Println("global.DataBaseSetting:",global.DataBaseSetting)
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if port !=""{
		global.ServerSetting.HttpPort=port
	}
	if runMode !=""{
		global.ServerSetting.RunMode=runMode

	}
	return nil

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err

	}
	return nil

}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {

	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831")
	if err != nil {
		return err

	}
	//注入到全局变量tracer中，以便于后续在中间件中使用，或定义不同到自定义span 中大点使用
	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {

	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式 ")
	flag.StringVar(&config,"config","configs/","指定要使用的配置文件路径")
	flag.Parse()
	return nil

}