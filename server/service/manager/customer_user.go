
package manager

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manager"
    managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/manager/request"
)

type CustomerUserService struct {}
// CreateCustomerUser 创建客户管理记录
// Author [yourname](https://github.com/yourname)
func (customerUserService *CustomerUserService) CreateCustomerUser(ctx context.Context, customerUser *manager.CustomerUser) (err error) {
	err = global.GVA_DB.Create(customerUser).Error
	return err
}

// DeleteCustomerUser 删除客户管理记录
// Author [yourname](https://github.com/yourname)
func (customerUserService *CustomerUserService)DeleteCustomerUser(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&manager.CustomerUser{},"id = ?",ID).Error
	return err
}

// DeleteCustomerUserByIds 批量删除客户管理记录
// Author [yourname](https://github.com/yourname)
func (customerUserService *CustomerUserService)DeleteCustomerUserByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]manager.CustomerUser{},"id in ?",IDs).Error
	return err
}

// UpdateCustomerUser 更新客户管理记录
// Author [yourname](https://github.com/yourname)
func (customerUserService *CustomerUserService)UpdateCustomerUser(ctx context.Context, customerUser manager.CustomerUser) (err error) {
	err = global.GVA_DB.Model(&manager.CustomerUser{}).Where("id = ?",customerUser.ID).Updates(&customerUser).Error
	return err
}

// GetCustomerUser 根据ID获取客户管理记录
// Author [yourname](https://github.com/yourname)
func (customerUserService *CustomerUserService)GetCustomerUser(ctx context.Context, ID string) (customerUser manager.CustomerUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&customerUser).Error
	return
}
// GetCustomerUserInfoList 分页获取客户管理记录
// Author [yourname](https://github.com/yourname)
func (customerUserService *CustomerUserService)GetCustomerUserInfoList(ctx context.Context, info managerReq.CustomerUserSearch) (list []manager.CustomerUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&manager.CustomerUser{})
    var customerUsers []manager.CustomerUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Code != nil && *info.Code != "" {
        db = db.Where("code = ?", *info.Code)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&customerUsers).Error
	return  customerUsers, total, err
}
func (customerUserService *CustomerUserService)GetCustomerUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
