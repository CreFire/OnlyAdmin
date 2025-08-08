package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/internal/repo"
)

var ruleRepo = &repo.RuleRepo{}

// 查询列表
func RuleListHandler(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	rules, err := ruleRepo.List(context.TODO(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rules)
}

// 新增
func RuleCreateHandler(c *gin.Context) {
	var req model.Rule
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TenantID = c.GetString("tenant_id")
	err := ruleRepo.Create(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 更新
func RuleUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delete(req, "_id")
	err := ruleRepo.Update(context.TODO(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 删除
func RuleDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := ruleRepo.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
