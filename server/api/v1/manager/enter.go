package manager

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AppManagementApi
	CustomerUserApi
}

var (
	appManagementService = service.ServiceGroupApp.ManagerServiceGroup.AppManagementService
	customerUserService  = service.ServiceGroupApp.ManagerServiceGroup.CustomerUserService
)
