package middleware

import (
	"log"
	"net/http"
	"os"
	"sports-backend/db"
	"sports-backend/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// tokenからユーザー情報を返す
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Split(authHeader, " ")[1]

		secretKey := os.Getenv("JWT_SECRET_KEY")

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			var user model.User
			if err := db.DB.Where("name = ?", claims.Name).First(&user).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				c.Abort()
				return
			}

			c.Set("user", user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
