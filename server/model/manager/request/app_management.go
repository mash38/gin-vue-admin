
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppManagementSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      App  *string `json:"app" form:"app"` 
      Title  *string `json:"title" form:"title"` 
    request.PageInfo
}
