package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-gin/config"
	"go-gin/model"
	mysql2 "go-gin/mysql"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type categoryStruct struct {
	Id   int
	Name string
}

var categoryList = map[string]categoryStruct{
	"1": {Id: 1, Name: "php"},
	"2": {Id: 2, Name: "laravel"},
	"3": {Id: 3, Name: "mysql"},
	"4": {Id: 4, Name: "docker"},
	"5": {Id: 5, Name: "redis"},
	"6": {Id: 6, Name: "rabbitmq"},
	"7": {Id: 7, Name: "go"},
	"8": {Id: 8, Name: "其他"},
}

type articleForm struct {
	Title        string   `form:"title"`
	ShortContent string   `form:"shortContent"`
	Content      string   `form:"content"`
	Tag          []string `form:"tag"`
}

func Index(c *gin.Context) {

	title := config.Get("app.title")
	about := config.Get("app.about")

	db := mysql2.Client()
	size := 100
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "0"))

	var articleList []model.Article
	db.Limit(size).Offset(currentPage * size).Order("id desc").Find(&articleList)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":        title,
		"categoryList": categoryList,
		"articleList":  articleList,
		"about":        template.HTML(about),
	})
}

func GetAddArticle(c *gin.Context) {
	title := config.Get("app.title")

	c.HTML(http.StatusOK, "addArticle.html", gin.H{
		"title":        title,
		"categoryList": categoryList,
	})
}

func PostAddArticle(c *gin.Context) {
	db := mysql2.Client()

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

	db := mysql2.Client()

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

	db := mysql2.Client()

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

func Article(c *gin.Context) {

	title := config.Get("app.title")
	about := config.Get("app.about")

	db := mysql2.Client()
	//defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	where := model.Article{Model: model.Model{Id: id}}

	var article = new(model.Article)
	db.First(&article, where)

	db.Model(article).Update("click_num", gorm.Expr("click_num + ?", 1))

	clickRecord := &model.ClickRecord{
		Ip:        c.ClientIP(),
		ArticleId: id,
	}
	db.Create(clickRecord)

	c.HTML(http.StatusOK, "article.html", gin.H{
		"title":   title,
		"article": article,
		"about":   template.HTML(about),
	})
}
