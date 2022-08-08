package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-gin/config"
	"go-gin/model"
	"go-gin/mysql"
	"html/template"
	"net/http"
	"strconv"
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

	db := mysql.Client()

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

func Article(c *gin.Context) {

	title := config.Get("app.title")
	about := config.Get("app.about")
	incClickCum := config.GetToBool("app.inc_click_num")

	db := mysql.Client()
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
	if incClickCum {
		db.Create(clickRecord)
	}

	c.HTML(http.StatusOK, "article.html", gin.H{
		"title":        title,
		"article":      article,
		"about":        template.HTML(about),
		"categoryList": categoryList,
	})
}
