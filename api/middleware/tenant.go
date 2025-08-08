package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// TenantMiddleware 提取 tenant_id 并注入 Context
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := ""
		// 优先Context
		if val, ok := c.Get(ContextTenantKey); ok {
			tenantID, _ = val.(string)
		}
		// header
		if tenantID == "" {
			tenantID = c.GetHeader("X-Tenant-ID")
		}
		// query
		if tenantID == "" {
			tenantID = c.Query("tenant_id")
		}
		// path param
		if tenantID == "" {
			tenantID = c.Param("tenant_id")
		}
		if tenantID == "" {
			logrus.WithFields(logrus.Fields{
				"header": c.Request.Header,
				"query":  c.Request.URL.RawQuery,
			}).Warn("缺少 tenant_id")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing tenant_id"})
			return
		}
		c.Set(ContextTenantKey, tenantID)
		c.Next()
	}
}
