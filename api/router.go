package api

import (
	"github.com/gin-gonic/gin"
	"server/api/middleware"
	"server/api/v1"
)

func RegisterRoutes(r *gin.Engine) {
	// 公共接口：只需要租户ID
	public := r.Group("/api/v1/public")
	public.Use(middleware.TenantMiddleware())
	{
		public.GET("/info", v1.PublicInfoHandler)
		// 如有更多租户通用信息接口，可继续加
	}

	// 管理后台接口：需登录 + 租户身份
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.JWTAuthMiddleware(), middleware.TenantMiddleware())
	{
		// 租户管理
		admin.GET("/tenant/list", v1.TenantListHandler)
		admin.POST("/tenant", v1.TenantCreateHandler)
		admin.PUT("/tenant/:id", v1.TenantUpdateHandler)
		admin.DELETE("/tenant/:id", v1.TenantDeleteHandler)

		// 房源管理
		admin.GET("/room/list", v1.RoomListHandler)
		admin.POST("/room", v1.RoomCreateHandler)
		admin.PUT("/room/:id", v1.RoomUpdateHandler)
		admin.DELETE("/room/:id", v1.RoomDeleteHandler)

		// 账单规则管理
		admin.GET("/rule/list", v1.RuleListHandler)
		admin.POST("/rule", v1.RuleCreateHandler)
		admin.PUT("/rule/:id", v1.RuleUpdateHandler)
		admin.DELETE("/rule/:id", v1.RuleDeleteHandler)

		// 账单管理
		admin.GET("/bill/list", v1.BillListHandler)
		admin.POST("/bill", v1.BillCreateHandler)
		admin.PUT("/bill/:id", v1.BillUpdateHandler)
		admin.DELETE("/bill/:id", v1.BillDeleteHandler)

		// 用户管理（如有子账号/操作员）
		admin.GET("/user/list", v1.UserListHandler)
		// ...其它管理接口
	}

	// 登录相关，不需要任何中间件
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", v1.LoginHandler)
		// 注册、找回密码等
		auth.POST("/forgot_password", v1.ForgotPasswordHandler) // 第一步，发送验证码
		auth.POST("/reset_password", v1.ResetPasswordHandler)   // 第二步，重置密码
	}

	// 也可以加租户自己的自助接口，如账单自助查询、微信推送等
	tenantSelf := r.Group("/api/v1/tenant")
	tenantSelf.Use(middleware.TenantMiddleware())
	{
		tenantSelf.GET("/bill/list", v1.TenantSelfBillListHandler)
		tenantSelf.POST("/bill/pay", v1.TenantSelfPayHandler)
	}
}
