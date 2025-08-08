package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/internal/repo"
)

var roomRepo = &repo.RoomRepo{}

// 查询列表
func RoomListHandler(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	rooms, err := roomRepo.List(context.TODO(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

// 新增
func RoomCreateHandler(c *gin.Context) {
	var req model.Room
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TenantID = c.GetString("tenant_id") // 多租户隔离
	err := roomRepo.Create(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 更新
func RoomUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delete(req, "_id")
	err := roomRepo.Update(context.TODO(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 删除
func RoomDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := roomRepo.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
