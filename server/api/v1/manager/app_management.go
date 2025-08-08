package manager

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/manager"
    managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/manager/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AppManagementApi struct {}



// CreateAppManagement 创建应用管理
// @Tags AppManagement
// @Summary 创建应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manager.AppManagement true "创建应用管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /appManagement/createAppManagement [post]
func (appManagementApi *AppManagementApi) CreateAppManagement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var appManagement manager.AppManagement
	err := c.ShouldBindJSON(&appManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appManagementService.CreateAppManagement(ctx,&appManagement)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAppManagement 删除应用管理
// @Tags AppManagement
// @Summary 删除应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manager.AppManagement true "删除应用管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /appManagement/deleteAppManagement [delete]
func (appManagementApi *AppManagementApi) DeleteAppManagement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := appManagementService.DeleteAppManagement(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAppManagementByIds 批量删除应用管理
// @Tags AppManagement
// @Summary 批量删除应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /appManagement/deleteAppManagementByIds [delete]
func (appManagementApi *AppManagementApi) DeleteAppManagementByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := appManagementService.DeleteAppManagementByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAppManagement 更新应用管理
// @Tags AppManagement
// @Summary 更新应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body manager.AppManagement true "更新应用管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /appManagement/updateAppManagement [put]
func (appManagementApi *AppManagementApi) UpdateAppManagement(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var appManagement manager.AppManagement
	err := c.ShouldBindJSON(&appManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appManagementService.UpdateAppManagement(ctx,appManagement)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAppManagement 用id查询应用管理
// @Tags AppManagement
// @Summary 用id查询应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询应用管理"
// @Success 200 {object} response.Response{data=manager.AppManagement,msg=string} "查询成功"
// @Router /appManagement/findAppManagement [get]
func (appManagementApi *AppManagementApi) FindAppManagement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reappManagement, err := appManagementService.GetAppManagement(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reappManagement, c)
}
// GetAppManagementList 分页获取应用管理列表
// @Tags AppManagement
// @Summary 分页获取应用管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query managerReq.AppManagementSearch true "分页获取应用管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appManagement/getAppManagementList [get]
func (appManagementApi *AppManagementApi) GetAppManagementList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo managerReq.AppManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := appManagementService.GetAppManagementInfoList(ctx,pageInfo)
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

// GetAppManagementPublic 不需要鉴权的应用管理接口
// @Tags AppManagement
// @Summary 不需要鉴权的应用管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appManagement/getAppManagementPublic [get]
func (appManagementApi *AppManagementApi) GetAppManagementPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    appManagementService.GetAppManagementPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的应用管理接口信息",
    }, "获取成功", c)
}
