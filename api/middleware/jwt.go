package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const (
	ContextTenantKey = "tenant_id"
	ContextUserIDKey = "user_id"
	JWTClaimTenantID = "tenant_id"
	JWTClaimUserID   = "user_id"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
			return
		}
		tokenString = parts[1]
		secret := viper.GetString("jwt_secret")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}
		tenantID := getClaimString(claims, JWTClaimTenantID)
		userID := getClaimString(claims, JWTClaimUserID)
		c.Set(ContextTenantKey, tenantID)
		c.Set(ContextUserIDKey, userID)
		c.Next()
	}
}

// 支持int/string类型的Claim读取
func getClaimString(claims jwt.MapClaims, key string) string {
	if val, ok := claims[key]; ok {
		switch v := val.(type) {
		case string:
			return v
		case float64: // jwt-go数字会变成float64
			return fmt.Sprintf("%.0f", v)
		}
	}
	return ""
}
