package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/CelticAlreadyUse/CMS/internal/config"
	"github.com/CelticAlreadyUse/CMS/internal/helper"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		splitAuth := strings.Split(authHeader, " ")
		if len(splitAuth) != 2 || splitAuth[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		accessToken := splitAuth[1]
		log.Println("Token:", accessToken)

		var claim model.CustomClaims
		token, err := jwt.ParseWithClaims(accessToken, &claim, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(config.JWTSigningKey()), nil
		})

		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
				return
			}
			log.Printf("Token Error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		if helper.IsTokenExpired(claim.ExpiresAt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
			return
		}
		ctx := context.WithValue(c.Request.Context(), model.BearerAuthKey, claim)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
