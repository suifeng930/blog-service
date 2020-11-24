package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)


// 主要作用是控制应用程序的初始化流程，起到将配置文件内容映射到应用配置结构体中
func init() {
	err := setupSetting()
	if err!=nil {
		log.Fatalf("init setupSetting err:%v",err)

	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v",err)
	}
}
func main() {

	//engine := gin.Default()
	//
	//engine.GET("/ping", func(context *gin.Context) {
	//	context.JSON(200,gin.H{"message":"pong"})
	//
	//})
	//
	//engine.Run()


	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":"+global.ServerSetting.HttpPort,
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
	setting, err := setting.NewSetting()
	if err!=nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err!=nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DataBaseSetting)
	if err!=nil {
		return err
	}
	//log.Println("global.ServerSetting:",global.ServerSetting)
	//log.Println("global.AppSetting:",global.AppSetting)
	//log.Println("global.DataBaseSetting:",global.DataBaseSetting)
	global.ServerSetting.ReadTimeout *=time.Second
	global.ServerSetting.WriteTimeout *=time.Second
	return nil

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err!=nil {
		return err

	}
	return nil

}