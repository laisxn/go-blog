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

var category model.Category
var categoryList = category.GetList()

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

	size, _ := strconv.Atoi(c.DefaultQuery("size", "100"))
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	category, _ := strconv.Atoi(c.DefaultQuery("category", "0"))

	var articleList []model.Article
	articleModel := db.Limit(size).Offset((currentPage - 1) * size).Order("id desc")
	if category > 0 {
		articleModel = articleModel.Where("FIND_IN_SET(?,tag)", category)
	}
	articleModel.Find(&articleList)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":        title,
		"categoryList": categoryList,
		"articleList":  articleList,
		"about":        template.HTML(about),
		"currentPage":  currentPage,
		"size":         currentPage,
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
		"title":        title,
		"article":      article,
		"about":        template.HTML(about),
		"categoryList": categoryList,
	})
}
