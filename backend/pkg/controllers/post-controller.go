package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/config"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	err := config.GetDB().Preload("User", func(db *gorm.DB) *gorm.DB {
    	return db.Select("id, name, email, created_at")
	}).Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	type Post struct {
		Title             string     `validate:"required,min=10,max=40"`
		Subtitle          string     `validate:"required,min=15,max=80"`
		Description       string     `validate:"required,min=100,max=700"`
	}

	id, exists := c.Get("id")
	if !exists {
		return
	}

	var input Post
	if err := c.ShouldBindJSON(&input); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
  	}

	validate := validator.New()
	err := validate.Struct(&input)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s validation failed on field %s", err.Tag(), err.Field()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
  	}

	y, ok := id.(float64)
	if ok {
		post := models.Post{Title: input.Title, Subtitle: input.Subtitle, Description: input.Description,Owner: uint(y)}
		config.GetDB().Create(&post)

		err2 := config.GetDB().Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", id)
		}).First(&post).Error
		if err2 != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
			return
		}
		c.JSON(http.StatusOK, post)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post create error"})
	}
}

func GetOnePost(c *gin.Context){
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Post ID"})
		return
	}

	var post models.Post
	err2  := config.GetDB().Preload("User").Where("id = ?", id).First(&post).Error
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context){
	idParam := c.Param("id")
	userId, exists := c.Get("id")
	if !exists {
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Post ID"})
		return
	}

	y, ok := userId.(float64)
	if ok {
		var input models.Post
		if err := c.ShouldBindJSON(&input); err != nil {
			if err.Error() == "EOF" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		var dbPost models.Post
		if result := config.GetDB().Where("id = ?", id).First(&dbPost); result.Error == nil {
			if uint(y) == dbPost.Owner {
				dbPost.ID = uint(id)
				if len(input.Title) > 9 && len(input.Title) < 41  {
					dbPost.Title = input.Title
					config.GetDB().Save(&dbPost)
				}
				if len(input.Subtitle) > 14 && len(input.Subtitle) < 81  {
					dbPost.Subtitle = input.Subtitle
					config.GetDB().Save(&dbPost)
				}
				if len(input.Description) > 99 && len(input.Description) < 601  {
					dbPost.Description = input.Description
					config.GetDB().Save(&dbPost)
				}
				if result := config.GetDB().Where("id = ?", id).First(&dbPost); result.Error == nil {
					c.JSON(http.StatusOK, result.Value)
					return
				}else{
					c.JSON(http.StatusBadRequest, gin.H{"error":"Post update error"})
					return
				}
			}else{
				c.JSON(http.StatusUnauthorized, gin.H{"error":"You are not post owner"})
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post not found"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post update error"})
	}
}

func DeletePost(c *gin.Context){
	idParam := c.Param("id")
	userId, exists := c.Get("id")
	if !exists {
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Post ID"})
		return
	}


	y, ok := userId.(float64)
	if ok {
		var post models.Post
	if result := config.GetDB().Where("id = ?", id).First(&post); result.Error == nil {
		if post.Owner == uint(y) {
			result := config.GetDB().Where("id = ?", id).Delete(&post);
			c.JSON(http.StatusOK, result.Value)
			return
			}else{
				c.JSON(http.StatusUnauthorized, gin.H{"error":"You are not post owner"})
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post not found"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post delete error"})
	}
}