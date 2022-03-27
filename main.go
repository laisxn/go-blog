package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "go-gin/config"
	"go-gin/routes"
	"os"
	"path/filepath"
)

func main() {

	r := gin.Default()
	//不使用代理
	r.SetTrustedProxies(nil)

	path, _ := os.Getwd()
	r.Static("/assets", "./static/assets")
	r.Static("/editor-md", "./static/editor-md")
	r.LoadHTMLGlob(filepath.Join(path, "static/view/*"))

	routes.Load(r)

	r.Run(":8000")
}
