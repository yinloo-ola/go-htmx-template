package main

import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views/assets", false)))
	//new template engine
	router.HTMLRender = ginview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/user"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	type User struct {
		Name   string
		Gender string
		Age    int
	}

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
			"users": []User{
				{Name: "Peter", Gender: "Male", Age: 30},
				{Name: "Alice", Gender: "Female", Age: 25},
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.POST("/createUser", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		gender := ctx.PostForm("gender")
		age := ctx.PostForm("age")
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "user.html", User{
			Name: name, Gender: gender, Age: ageInt,
		})
	})

	router.Run(":9090")
}
