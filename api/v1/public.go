package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/api/middleware"
)

func PublicInfoHandler(c *gin.Context) {
	tenantID := c.GetString(middleware.ContextTenantKey)
	c.JSON(http.StatusOK, gin.H{
		"tenant_id": tenantID,
		"msg":       "this is a public info.",
	})
}
