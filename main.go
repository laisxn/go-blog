package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"go-gin/config"
	_ "go-gin/config"
	"go-gin/mysql"
	"go-gin/routes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	gin.SetMode(config.Get("app.debug_model"))

	f, _ := os.Create("./runtime/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	fErr, _ := os.Create("./runtime/gin_err.log")
	gin.DefaultErrorWriter = io.MultiWriter(fErr, os.Stdout)

	var initSql int
	flag.IntVar(&initSql, "init_sql", 0, "是否初始导入sql文件，默认0 否 1是")
	flag.Parse()

	if initSql == 1 { //初始导入sql
		sqls, _ := ioutil.ReadFile("。/blog.sql")
		sqlArr := strings.Split(string(sqls), ";")
		for key, sql := range sqlArr {
			if sql == "" || key == len(sqlArr)-1 {
				continue
			}
			mysql.Client().Exec(sql)
		}
	}
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

	r.Run(":" + config.Get("app.web_port"))
}
