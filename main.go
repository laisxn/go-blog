package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jmoiron/sqlx"
)

var pool *redis.Pool //创建redis连接池
type point struct {
	a, b int
}

func init() {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

var blog_title = "食之有味"
var about = "大鱼吃小鱼，小鱼吃虾米，而我是<em>大小通吃</em>"

type categoryStruct struct {
	Id   int
	Name string
}

type article struct {
	Id           int
	Title        string
	ShortContent string
	Content      string
	Tag          string     `json:"tag"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
}
type articleForm struct {
	Title        string   `form:"title"`
	ShortContent string   `form:"shortContent"`
	Content      string   `form:"content"`
	Tag          []string `form:"tag"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog?parseTime=true")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection success")
	}

	defer db.Close()
	r := gin.Default()

	categoryList := map[string]categoryStruct{
		"1": {Id: 1, Name: "php"},
		"2": {Id: 2, Name: "laravel"},
		"3": {Id: 3, Name: "mysql"},
		"4": {Id: 4, Name: "docker"},
		"5": {Id: 5, Name: "redis"},
		"6": {Id: 6, Name: "rabbitmq"},
	}

	//articleList := map[string]article{}

	r.Static("/assets", "./static/assets")
	r.Static("/editor-md", "./static/editor-md")
	r.LoadHTMLGlob("static/view/*")

	r.GET("/", func(c *gin.Context) {
		var articleList []article
		db.Find(&articleList)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":        blog_title,
			"categoryList": categoryList,
			"articleList":  articleList,
			"about":        template.HTML(about),
		})
	})
	r.GET("/admin/addArticle", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addArticle.html", gin.H{
			"title":        blog_title,
			"categoryList": categoryList,
		})
	})
	r.POST("/admin/addArticle", func(c *gin.Context) {
		var form articleForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(300, gin.H{"msg": err})
		}

		article := &article{
			Title:        form.Title,
			ShortContent: form.ShortContent,
			Content:      form.Content,
			Tag:          strings.Join(form.Tag, ","),
		}
		db.Create(article)
		// 页面接收
		c.JSON(200, gin.H{"request": form})
	})
	r.GET("/article/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		where := article{Id: id}

		var article = new(article)
		db.First(&article, where)

		c.HTML(http.StatusOK, "article.html", gin.H{
			"title":   blog_title,
			"article": article,
			"about":   template.HTML(about),
		})
	})
	r.GET("/long_async", func(c *gin.Context) {
		c1 := pool.Get() //从连接池，取一个链接
		defer c1.Close() //函数运行结束 ，把连接放回连接池

		_, err := c1.Do("Set", "abc", 123)
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err := redis.String(c1.Do("Get", "abc"))
		if err != nil {
			fmt.Println("get abc faild :", err)
			return
		}
		fmt.Println(r)
		pool.Close() //关闭连接池
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "食之无畏"})
	})

	r.Run(":8000")
}
