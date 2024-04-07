package main

import (
	"GinStudy/src/main/config"
	"GinStudy/src/main/router"
	member "GinStudy/src/main/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.New()
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义规则
		validate.RegisterValidation("NameValid", member.NameValid)
	}
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}
