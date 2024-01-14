package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/config"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/models"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
  type User struct {
    Name           string     `validate:"required,min=4,max=18"`
    Email          string     `validate:"required,email,min=8,max=120"`
    Password       string     `validate:"required,min=6,max=36"`
  }

	var input User
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

  var findUser User
  if result := config.GetDB().Where("email = ?", input.Email).First(&findUser); result.Error != nil {
      user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

      config.GetDB().Create(&user)
      token := utils.GenerateJWTData{ID: user.ID,Email: user.Email}
      tokenString, err := utils.GenerateJWT(token)
      if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
      }

      user.Token = tokenString
      config.GetDB().Save(&user)
      selectedData := struct {
          ID    uint          `json:"id"`
          Name  string        `json:"name"`
          Email string        `json:"email"`
          Token string        `json:"token"`
          CreatedAt time.Time `json:"created_at"`
      }{
          ID:    user.ID,
          Name:  user.Name,
          Email: user.Email,
          Token: user.Token,
          CreatedAt: user.CreatedAt,
      }
      c.JSON(http.StatusOK, selectedData)
  } else {
      c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
  }
}

func Login(c *gin.Context) {
  type User struct {
    ID             uint
    Name           string
    Email          string     `validate:"required,email"`
    Password       string     `validate:"required,min=6"`
    Token          string  
    CreatedAt      time.Time
  }

	var input User
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

  var findUser User
  if result := config.GetDB().Where("email = ?", input.Email).First(&findUser); result.Error == nil {
      if findUser.Password == input.Password {
        token := utils.GenerateJWTData{ID: findUser.ID,Email: findUser.Email}
        tokenString, err := utils.GenerateJWT(token)
        if err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
        }
        findUser.Token = tokenString
        config.GetDB().Save(&findUser)
        if result := config.GetDB().Where("email = ?", input.Email).First(&findUser); result.Error == nil{
            selectedData := struct {
                ID          uint            `json:"id"`
                Name        string          `json:"name"`
                Email       string          `json:"email"`
                Token       string          `json:"token"`
                CreatedAt   time.Time       `json:"created_at"`
            }{
                ID:    findUser.ID,
                Name:  findUser.Name,
                Email: findUser.Email,
                Token: findUser.Token,
                CreatedAt: findUser.CreatedAt,
            }
            c.JSON(http.StatusOK, selectedData)
        }
      }else{
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Password wrong"})
      }

  } else {
      c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
  }
}

func Control(c *gin.Context) {
  type User struct {
    ID             uint       `json:"id"`
    Name           string     `json:"name"`
    Email          string     `json:"email"`
    Token          string     `json:"token"`
    CreatedAt      time.Time  `json:"created_at"`
  }


  id, exists := c.Get("id")
	if !exists {
		return
	}

  var user User

  y, ok := id.(float64)
	if ok {
		config.GetDB().First(&user, "id = ?", y)

    c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Post create error"})
	}
}