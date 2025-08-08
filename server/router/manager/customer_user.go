package manager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerUserRouter struct {}

// InitCustomerUserRouter 初始化 客户管理 路由信息
func (s *CustomerUserRouter) InitCustomerUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	customerUserRouter := Router.Group("customerUser").Use(middleware.OperationRecord())
	customerUserRouterWithoutRecord := Router.Group("customerUser")
	customerUserRouterWithoutAuth := PublicRouter.Group("customerUser")
	{
		customerUserRouter.POST("createCustomerUser", customerUserApi.CreateCustomerUser)   // 新建客户管理
		customerUserRouter.DELETE("deleteCustomerUser", customerUserApi.DeleteCustomerUser) // 删除客户管理
		customerUserRouter.DELETE("deleteCustomerUserByIds", customerUserApi.DeleteCustomerUserByIds) // 批量删除客户管理
		customerUserRouter.PUT("updateCustomerUser", customerUserApi.UpdateCustomerUser)    // 更新客户管理
	}
	{
		customerUserRouterWithoutRecord.GET("findCustomerUser", customerUserApi.FindCustomerUser)        // 根据ID获取客户管理
		customerUserRouterWithoutRecord.GET("getCustomerUserList", customerUserApi.GetCustomerUserList)  // 获取客户管理列表
	}
	{
	    customerUserRouterWithoutAuth.GET("getCustomerUserPublic", customerUserApi.GetCustomerUserPublic)  // 客户管理开放接口
	}
}
