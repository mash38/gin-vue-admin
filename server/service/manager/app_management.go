
package manager

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manager"
    managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/manager/request"
)

type AppManagementService struct {}
// CreateAppManagement 创建应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) CreateAppManagement(ctx context.Context, appManagement *manager.AppManagement) (err error) {
	err = global.GVA_DB.Create(appManagement).Error
	return err
}

// DeleteAppManagement 删除应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService)DeleteAppManagement(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&manager.AppManagement{},"id = ?",ID).Error
	return err
}

// DeleteAppManagementByIds 批量删除应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService)DeleteAppManagementByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]manager.AppManagement{},"id in ?",IDs).Error
	return err
}

// UpdateAppManagement 更新应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService)UpdateAppManagement(ctx context.Context, appManagement manager.AppManagement) (err error) {
	err = global.GVA_DB.Model(&manager.AppManagement{}).Where("id = ?",appManagement.ID).Updates(&appManagement).Error
	return err
}

// GetAppManagement 根据ID获取应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService)GetAppManagement(ctx context.Context, ID string) (appManagement manager.AppManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&appManagement).Error
	return
}
// GetAppManagementInfoList 分页获取应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService)GetAppManagementInfoList(ctx context.Context, info managerReq.AppManagementSearch) (list []manager.AppManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&manager.AppManagement{})
    var appManagements []manager.AppManagement
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.App != nil && *info.App != "" {
        db = db.Where("app LIKE ?", "%"+ *info.App+"%")
    }
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title = ?", *info.Title)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&appManagements).Error
	return  appManagements, total, err
}
func (appManagementService *AppManagementService)GetAppManagementPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
