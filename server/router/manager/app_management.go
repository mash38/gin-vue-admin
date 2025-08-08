package manager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppManagementRouter struct {}

// InitAppManagementRouter 初始化 应用管理 路由信息
func (s *AppManagementRouter) InitAppManagementRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	appManagementRouter := Router.Group("appManagement").Use(middleware.OperationRecord())
	appManagementRouterWithoutRecord := Router.Group("appManagement")
	appManagementRouterWithoutAuth := PublicRouter.Group("appManagement")
	{
		appManagementRouter.POST("createAppManagement", appManagementApi.CreateAppManagement)   // 新建应用管理
		appManagementRouter.DELETE("deleteAppManagement", appManagementApi.DeleteAppManagement) // 删除应用管理
		appManagementRouter.DELETE("deleteAppManagementByIds", appManagementApi.DeleteAppManagementByIds) // 批量删除应用管理
		appManagementRouter.PUT("updateAppManagement", appManagementApi.UpdateAppManagement)    // 更新应用管理
	}
	{
		appManagementRouterWithoutRecord.GET("findAppManagement", appManagementApi.FindAppManagement)        // 根据ID获取应用管理
		appManagementRouterWithoutRecord.GET("getAppManagementList", appManagementApi.GetAppManagementList)  // 获取应用管理列表
	}
	{
	    appManagementRouterWithoutAuth.GET("getAppManagementPublic", appManagementApi.GetAppManagementPublic)  // 应用管理开放接口
	}
}
