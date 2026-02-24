package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if(result.Error != nil){
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}




func PostIndex(c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)
	
	c.JSON(200, gin.H{
		"posts": posts,
	})
}