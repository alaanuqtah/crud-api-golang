package controllers

import (
	"crud-api-golang/initializers"
	"crud-api-golang/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//GET DATA OFF REQ BODY
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//CREATE POST

	//RETURN IT
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//get the id of the url
	id := c.Param("id")

	//get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find post where updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model((&post)).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {

	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	//respond with them
	c.JSON(200, gin.H{
		"message": "successfuly deleted",
	})

}
