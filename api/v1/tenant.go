package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/internal/repo"
)

var tenantRepo = &repo.TenantRepo{}

// 查询列表
func TenantListHandler(c *gin.Context) {
	// tenantID := c.GetString("tenant_id") // 如果需要租户隔离
	list, err := tenantRepo.List(context.TODO(), "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// 新增
func TenantCreateHandler(c *gin.Context) {
	var req model.Tenant
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	err := tenantRepo.Create(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 更新
func TenantUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delete(req, "_id") // 防止_id被更改
	err := tenantRepo.Update(context.TODO(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 删除
func TenantDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := tenantRepo.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
