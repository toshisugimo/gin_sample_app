package main

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"app/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/tmpl/*.html")

	r.GET("/", func(c *gin.Context) {
		todos := controller.List()
		c.HTML(http.StatusOK, "index.html", gin.H {
			"todos": todos,
		})
	})

	r.GET("/todo/:id/edit", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := controller.Get(id)
		c.HTML(http.StatusOK, "edit.html", gin.H {
			"todo": todo,
		})
    })

	r.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		status := c.PostForm("status")
		controller.Create(text, status)
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/todo/:id/update", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := c.PostForm("text")
		status := c.PostForm("status")
		controller.Update(id, text, status)
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/todo/:id/delete", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		controller.Delete(id)
		c.Redirect(http.StatusFound, "/")
	})

	r.Run()
}
