package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	g := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	g.GET("/", func(context *gin.Context) {
		sugar.Info("index")
		context.String(200, "index")
	})
	g.GET("/api/user/info", func(context *gin.Context) {
		sugar.Warn("/api/user/info")
		stu := Student{Name: "张三", Age: 12}
		context.JSON(200, stu)
	})
	g.GET("/api2/user/info", func(context *gin.Context) {
		stu := Student{Name: "张三api2", Age: 12}
		context.JSON(200, stu)
	})
	g.Run()
}
