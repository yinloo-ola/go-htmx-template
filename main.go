package main

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views/assets", false)))
	//new template engine
	router.HTMLRender = ginview.Default()

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

	router.GET("/users", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "users.html", gin.H{"users": []User{
			{Name: "Peter", Gender: "Male", Age: 30},
			{Name: "Alice", Gender: "Female", Age: 25},
		}})
	})

	router.Run(":9090")
}
