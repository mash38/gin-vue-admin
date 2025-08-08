package manager

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AppManagementRouter
	CustomerUserRouter
}

var (
	appManagementApi = api.ApiGroupApp.ManagerApiGroup.AppManagementApi
	customerUserApi  = api.ApiGroupApp.ManagerApiGroup.CustomerUserApi
)
