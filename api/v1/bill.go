package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/internal/repo"
)

var billRepo = &repo.BillRepo{}

// POST /api/v1/tenant/bill/pay
func TenantSelfPayHandler(c *gin.Context) {
	type PayRequest struct {
		BillID        string `json:"bill_id" binding:"required"`
		PaymentMethod string `json:"payment_method" binding:"required"` // 如 "wechat", "alipay", "cash"
		Voucher       string `json:"voucher"`                           // 支付凭证图片链接
		PayTime       string `json:"pay_time"`                          // 支付时间（可前端传，建议服务器自动生成）
	}
	var req PayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//tenantID := c.GetString("tenant_id")
	// 可校验此bill_id是否属于该tenantID，防止越权
	update := map[string]interface{}{
		"status":         "paid",
		"payment_method": req.PaymentMethod,
		"voucher":        req.Voucher,
	}
	if req.PayTime != "" {
		update["pay_time"] = req.PayTime
	}
	err := billRepo.Update(context.TODO(), req.BillID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "pay success"})
}

// GET /api/v1/tenant/bill/list
func TenantSelfBillListHandler(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenant_id is required"})
		return
	}
	// 可按需加筛选条件，比如 status/period
	bills, err := billRepo.List(context.TODO(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bills)
}

// 查询列表
func BillListHandler(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	bills, err := billRepo.List(context.TODO(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bills)
}

// 新增
func BillCreateHandler(c *gin.Context) {
	var req model.Bill
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TenantID = c.GetString("tenant_id")
	err := billRepo.Create(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 更新
func BillUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delete(req, "_id")
	err := billRepo.Update(context.TODO(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 删除
func BillDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := billRepo.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
