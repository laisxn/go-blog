package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/model"
	"go-gin/mysql"
	"net/http"
)

func AddChatRecord(c *gin.Context) {

	chatRecord := &model.ChatRecord{
		Pid:          0,
		Content:      c.PostForm("content"),
		UserId:       0,
		UserNickname: c.PostForm("user_nickname"),
		Ip:           c.PostForm("ip"),
	}
	db := mysql.Client()
	db.Create(chatRecord)
	// 页面接收
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功"})
	return
}
