package manager

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/manager"
    managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/manager/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CustomerUserApi struct {}



// CreateCustomerUser 创建客户管理
// @Tags CustomerUser
// @Summary 创建客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manager.CustomerUser true "创建客户管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /customerUser/createCustomerUser [post]
func (customerUserApi *CustomerUserApi) CreateCustomerUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var customerUser manager.CustomerUser
	err := c.ShouldBindJSON(&customerUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerUserService.CreateCustomerUser(ctx,&customerUser)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCustomerUser 删除客户管理
// @Tags CustomerUser
// @Summary 删除客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manager.CustomerUser true "删除客户管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /customerUser/deleteCustomerUser [delete]
func (customerUserApi *CustomerUserApi) DeleteCustomerUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := customerUserService.DeleteCustomerUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCustomerUserByIds 批量删除客户管理
// @Tags CustomerUser
// @Summary 批量删除客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /customerUser/deleteCustomerUserByIds [delete]
func (customerUserApi *CustomerUserApi) DeleteCustomerUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := customerUserService.DeleteCustomerUserByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCustomerUser 更新客户管理
// @Tags CustomerUser
// @Summary 更新客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manager.CustomerUser true "更新客户管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /customerUser/updateCustomerUser [put]
func (customerUserApi *CustomerUserApi) UpdateCustomerUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var customerUser manager.CustomerUser
	err := c.ShouldBindJSON(&customerUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerUserService.UpdateCustomerUser(ctx,customerUser)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCustomerUser 用id查询客户管理
// @Tags CustomerUser
// @Summary 用id查询客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询客户管理"
// @Success 200 {object} response.Response{data=manager.CustomerUser,msg=string} "查询成功"
// @Router /customerUser/findCustomerUser [get]
func (customerUserApi *CustomerUserApi) FindCustomerUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	recustomerUser, err := customerUserService.GetCustomerUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(recustomerUser, c)
}
// GetCustomerUserList 分页获取客户管理列表
// @Tags CustomerUser
// @Summary 分页获取客户管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query managerReq.CustomerUserSearch true "分页获取客户管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /customerUser/getCustomerUserList [get]
func (customerUserApi *CustomerUserApi) GetCustomerUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo managerReq.CustomerUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := customerUserService.GetCustomerUserInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetCustomerUserPublic 不需要鉴权的客户管理接口
// @Tags CustomerUser
// @Summary 不需要鉴权的客户管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /customerUser/getCustomerUserPublic [get]
func (customerUserApi *CustomerUserApi) GetCustomerUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    customerUserService.GetCustomerUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的客户管理接口信息",
    }, "获取成功", c)
}
