package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs" //必须添加这个，不然会报错 not yet registered swag
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

func NewRouter() *gin.Engine {

	engine := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	} else {
		engine.Use(middleware.AccessLog())
		engine.Use(middleware.Recovery())
	}
	engine.Use(middleware.RateLimiter(methodLimiters))
	engine.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	// 采用国际化处理
	engine.Use(middleware.Translations())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 文件上传中间件
	upload := api.NewUpload()
	engine.POST("/upload/file", upload.UploadFile)
	// 静态资源访问配置
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 鉴权信息 jwt
	engine.POST("/auth", api.GetAuth)
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := engine.Group("/api/v1")
	// 将jwt 纳入到 api/v1 中
	apiv1.Use(middleware.JWT())
	{
		//tags
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// article

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return engine

}

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
