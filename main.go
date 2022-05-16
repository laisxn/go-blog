package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"go-gin/config"
	_ "go-gin/config"
	"go-gin/routes"
	"io"
	"os"
	"path/filepath"
)

func init() {
	gin.SetMode(config.Get("app.debug_model"))

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	fErr, _ := os.Create("gin_err.log")
	gin.DefaultErrorWriter = io.MultiWriter(fErr, os.Stdout)

}

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
