package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/model"
	"go-gin/mysql"
	"math/rand"
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
		c.JSON(300, gin.H{"msg": err})
	}

	comment := &model.Comment{
		ArticleId:    form.ArticleId,
		Content:      form.Content,
		UserId:       0,
		UserNickName: randStr(10),
		Ip:           c.ClientIP(),
	}
	db.Create(comment)
	// 页面接收
	c.JSON(200, gin.H{"request": form})
}
