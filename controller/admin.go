package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/model"
	"go-gin/mysql"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func GetAddArticle(c *gin.Context) {
	title := config.Get("app.title")

	c.HTML(http.StatusOK, "addArticle.html", gin.H{
		"title":        title,
		"categoryList": categoryList,
	})
}

func PostAddArticle(c *gin.Context) {
	db := mysql.Client()

	var form articleForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(300, gin.H{"msg": err})
	}

	article := &model.Article{
		Title:        form.Title,
		ShortContent: form.ShortContent,
		Content:      form.Content,
		Tag:          strings.Join(form.Tag, ","),
	}
	db.Create(article)
	// 页面接收
	c.JSON(200, gin.H{"request": form})
}

func GetUpdateArticle(c *gin.Context) {
	title := config.Get("app.title")
	about := config.Get("app.about")

	db := mysql.Client()

	id, _ := strconv.Atoi(c.Param("id"))
	where := model.Article{
		Model: model.Model{Id: id},
	}

	var article = new(model.Article)
	db.First(&article, where)

	c.HTML(http.StatusOK, "updateArticle.html", gin.H{
		"title":        title,
		"article":      article,
		"tag":          strings.Split(article.Tag, ","),
		"about":        template.HTML(about),
		"categoryList": categoryList,
	})
}

func PostUpdateArticle(c *gin.Context) {

	db := mysql.Client()

	var form articleForm
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(300, gin.H{"msg": err})
	}

	article := &model.Article{
		Model:        model.Model{Id: id},
		Title:        form.Title,
		ShortContent: form.ShortContent,
		Content:      form.Content,
		Tag:          strings.Join(form.Tag, ","),
	}
	db.Save(article)
	// 页面接收
	c.JSON(200, gin.H{"request": form})
}
