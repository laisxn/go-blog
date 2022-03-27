package mysql

import (
	"github.com/jinzhu/gorm"
	"go-gin/config"
	"sync"
)

type connect struct {
	client *gorm.DB
}

var once = sync.Once{}

var _connect *connect

func connectDb() {

	mysql := config.Instance().Section("mysql")

	db, err := gorm.Open("mysql", mysql.Key("username").String()+":"+mysql.Key("password").String()+"@tcp("+mysql.Key("host").String()+":"+mysql.Key("port").String()+")/"+mysql.Key("database").String()+"?parseTime=true&loc=Local")

	if err != nil {
		panic(err)
	}

	_connect = &connect{
		client: db,
	}
}

func Client() *gorm.DB {
	if _connect == nil {
		once.Do(func() {
			connectDb()
		})
	}

	return _connect.client
}
