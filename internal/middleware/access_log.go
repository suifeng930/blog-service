package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"time"
)

//构造一个访问日志的writer结构体,

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)

}

// 实现自定义log 重写
func AccessLog() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 将Access Writer 写入流 替换为原来的 http.ResponseWriter
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		// 将Access Writer 写入流 替换为原来的 http.ResponseWriter
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method %s,status_code: %d, begin_time: %d, end_time: %d"
		//将请求日志信息，持久化到自定义log日志组件中
		global.Logger.WithFields(fields).Infof(s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime)


	}

}
