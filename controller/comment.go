package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/model"
	"go-gin/mysql"
	"html/template"
	"math/rand"
	"net/http"
)

type commentForm struct {
	ArticleId int    `form:"article_id"`
	Content   string `form:"content"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//随机字符串
func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AddComment(c *gin.Context) {
	db := mysql.Client()

	var form commentForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}

	comment := &model.Comment{
		Pid:          0,
		ArticleId:    form.ArticleId,
		Content:      template.HTMLEscapeString(form.Content),
		UserId:       0,
		UserNickname: randStr(10),
		Ip:           c.ClientIP(),
	}
	db.Create(comment)
	// 页面接收
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功"})
}
