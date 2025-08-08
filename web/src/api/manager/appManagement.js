import service from '@/utils/request'
// @Tags AppManagement
// @Summary 创建应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppManagement true "创建应用管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /appManagement/createAppManagement [post]
export const createAppManagement = (data) => {
  return service({
    url: '/appManagement/createAppManagement',
    method: 'post',
    data
  })
}

// @Tags AppManagement
// @Summary 删除应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppManagement true "删除应用管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appManagement/deleteAppManagement [delete]
export const deleteAppManagement = (params) => {
  return service({
    url: '/appManagement/deleteAppManagement',
    method: 'delete',
    params
  })
}

// @Tags AppManagement
// @Summary 批量删除应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除应用管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appManagement/deleteAppManagement [delete]
export const deleteAppManagementByIds = (params) => {
  return service({
    url: '/appManagement/deleteAppManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags AppManagement
// @Summary 更新应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppManagement true "更新应用管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appManagement/updateAppManagement [put]
export const updateAppManagement = (data) => {
  return service({
    url: '/appManagement/updateAppManagement',
    method: 'put',
    data
  })
}

// @Tags AppManagement
// @Summary 用id查询应用管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AppManagement true "用id查询应用管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appManagement/findAppManagement [get]
export const findAppManagement = (params) => {
  return service({
    url: '/appManagement/findAppManagement',
    method: 'get',
    params
  })
}

// @Tags AppManagement
// @Summary 分页获取应用管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取应用管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appManagement/getAppManagementList [get]
export const getAppManagementList = (params) => {
  return service({
    url: '/appManagement/getAppManagementList',
    method: 'get',
    params
  })
}

// @Tags AppManagement
// @Summary 不需要鉴权的应用管理接口
// @Accept application/json
// @Produce application/json
// @Param data query managerReq.AppManagementSearch true "分页获取应用管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appManagement/getAppManagementPublic [get]
export const getAppManagementPublic = () => {
  return service({
    url: '/appManagement/getAppManagementPublic',
    method: 'get',
  })
}
