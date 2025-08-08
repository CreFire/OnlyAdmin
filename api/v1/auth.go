package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"server/internal/repo"
)

// 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	TenantID string `json:"tenant_id" binding:"required"`
}

// 登录接口
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	user, err := repo.FindUserByNameAndTenant(req.Username, req.TenantID)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}
	if user.Password != req.Password { // 生产一定要加密存储和校验
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	// 生成 JWT Token
	claims := jwt.MapClaims{
		"user_id":   user.ID.Hex(),
		"username":  user.Username,
		"tenant_id": user.TenantID,
		"exp":       time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtStr, err := token.SignedString([]byte(viper.GetString("jwt_secret")))
	if err != nil {
		logrus.Error("JWT 生成失败: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token生成失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": jwtStr})
}
