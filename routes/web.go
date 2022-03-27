package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin/controller"
)

func Load(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.GET("/admin/addArticle", controller.GetAddArticle)
	r.POST("/admin/addArticle", controller.PostAddArticle)
	r.GET("/admin/updateArticle/:id", controller.GetUpdateArticle)
	r.POST("/admin/updateArticle/:id", controller.PostUpdateArticle)
	r.GET("/article/:id", controller.Article)
}
