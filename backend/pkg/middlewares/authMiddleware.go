package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/config"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/models"
	"github.com/fatihsen-dev/go-fullstack-social-media/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
        }

        authToken := authHeader[len("Bearer "):]
        if authToken == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
        }

        token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
            secret := []byte(utils.GetEnvVariable("JWT_SECRET_KEY"))
            return secret, nil
        })

        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Invalid token"})
            c.Abort()
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            id := claims["id"].(float64)
            var findUser models.User
            config.GetDB().Select("token").First(&findUser, "id = ?", id)
            userToken := findUser.Token
            if userToken == authToken {
                c.Set("id",id)
                c.Next()
            }else {
                c.JSON(http.StatusNotFound, gin.H{"error": "Invalid token"})
                c.Abort()
            }
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
        }
    }
}