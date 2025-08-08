package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/repo"
)

// GET /api/v1/user/list
func UserListHandler(c *gin.Context) {
	logrus.Info("查询用户列表")
	c.JSON(http.StatusOK, gin.H{
		"data": []string{"userA", "userB"},
		"msg":  "success",
	})
}

// POST /api/v1/auth/reset_password
func ResetPasswordHandler(c *gin.Context) {
	type Req struct {
		Email       string `json:"email" binding:"required"`
		Code        string `json:"code" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 验证code（你需要实现验证码校验逻辑，比如Redis缓存里查找）
	// ok := service.VerifyCode(req.Email, req.Code)
	ok := true // 伪代码
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码无效或已过期"})
		return
	}
	// 更新数据库密码（你需要实现加密存储）
	err := repo.UpdateUserPasswordByEmail(req.Email, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "重置失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "密码已重置"})
}

// POST /api/v1/auth/forgot_password
func ForgotPasswordHandler(c *gin.Context) {
	type Req struct {
		Email string `json:"email" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 查询用户是否存在
	user, err := repo.FindUserByEmail(req.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "如果邮箱存在，将发送邮件"}) // 避免暴露注册状态
		return
	}
	// 生成验证码或重置链接并发送（这里你可以实现自己的邮件发送逻辑）
	// code := service.GenerateCode()
	// service.SendEmail(req.Email, code)
	// service.CacheCode(req.Email, code)
	c.JSON(http.StatusOK, gin.H{"msg": "如果邮箱存在，将发送邮件"})
}
