package api

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
	Name string `json:"name"`
	jwt.StandardClaims
}

// tokenを基にユーザー情報を返す
func Home(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		var user model.User
		if err := db.DB.Where("name = ?", claims.Name).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}
