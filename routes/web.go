package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/controller"
)

func Load(r *gin.Engine) {
	r.GET("/favicon.ico", func(context *gin.Context) {

	})
	r.GET("/", controller.Index)
	r.GET("/article/:id", controller.Article)
	r.POST("/comment", controller.AddComment)
	r.POST("/chatRecord", controller.AddChatRecord)

	adminGroup := r.Group("/admin")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{config.Get("auth.username"): config.Get("auth.password")}))

	adminGroup.GET("/addArticle", controller.GetAddArticle)
	adminGroup.POST("/addArticle", controller.PostAddArticle)
	adminGroup.GET("/updateArticle/:id", controller.GetUpdateArticle)
	adminGroup.POST("/updateArticle/:id", controller.PostUpdateArticle)
}
