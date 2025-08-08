package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manager"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(manager.AppManagement{}, manager.CustomerUser{})
	if err != nil {
		return err
	}
	return nil
}
