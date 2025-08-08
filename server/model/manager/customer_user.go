
// 自动生成模板CustomerUser
package manager
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 客户管理 结构体  CustomerUser
type CustomerUser struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //客户名称
  Parent_id  *int `json:"parent_id" form:"parent_id" gorm:"column:parent_id;"`  //推荐人ID
  C_type  *int `json:"c_type" form:"c_type" gorm:"comment:客户类型ID;column:c_type;"`  //客户类型
  Code  *string `json:"code" form:"code" gorm:"column:code;"`  //客户编码
  Source  *int `json:"source" form:"source" gorm:"column:source;"`  //客户来源
}


// TableName 客户管理 CustomerUser自定义表名 customer_user
func (CustomerUser) TableName() string {
    return "customer_user"
}





