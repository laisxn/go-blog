package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/model"
	"go-gin/mysql"
	"net/http"
)

func AddChatRecord(c *gin.Context) {

	post_api_auth_key := c.PostForm("api_auth_key")
	api_auth_key := config.Get("app.api_auth_key")

	if api_auth_key != post_api_auth_key {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}

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
