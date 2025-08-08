// 自动生成模板AppManagement
package manager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 应用管理 结构体  AppManagement
type AppManagement struct {
	global.GVA_MODEL
	App         *string        `json:"app" form:"app" gorm:"column:app;" binding:"required"`                     //应用名称
	Customer_id *int           `json:"customer_id" form:"customer_id" gorm:"column:customer_id;"`                //客户ID
	Title       *string        `json:"title" form:"title" gorm:"column:title;"`                                  //代号
	Logo        string         `json:"logo" form:"logo" gorm:"column:logo;"`                                     //logo
	Ico         string         `json:"ico" form:"ico" gorm:"column:ico;"`                                        //ico
	Version     *string        `json:"version" form:"version" gorm:"column:version;"`                            //版本号
	Android     datatypes.JSON `json:"android" form:"android" gorm:"column:android;" swaggertype:"array,object"` //安卓链接
	Ios         datatypes.JSON `json:"ios" form:"ios" gorm:"column:ios;" swaggertype:"array,object"`             //苹果链接
	Nav         datatypes.JSON `json:"nav" form:"nav" gorm:"column:nav;" swaggertype:"array,object"`             //应用链接
	Bg          string         `json:"bg" form:"bg" gorm:"column:bg;"`                                           //背景图片
	Remark      *string        `json:"remark" form:"remark" gorm:"column:remark;"`                               //备注
}

// TableName 应用管理 AppManagement自定义表名 app_management
func (AppManagement) TableName() string {
	return "app_management"
}
