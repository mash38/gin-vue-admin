import service from '@/utils/request'
// @Tags CustomerUser
// @Summary 创建客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CustomerUser true "创建客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customerUser/createCustomerUser [post]
export const createCustomerUser = (data) => {
  return service({
    url: '/customerUser/createCustomerUser',
    method: 'post',
    data
  })
}

// @Tags CustomerUser
// @Summary 删除客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CustomerUser true "删除客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customerUser/deleteCustomerUser [delete]
export const deleteCustomerUser = (params) => {
  return service({
    url: '/customerUser/deleteCustomerUser',
    method: 'delete',
    params
  })
}

// @Tags CustomerUser
// @Summary 批量删除客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customerUser/deleteCustomerUser [delete]
export const deleteCustomerUserByIds = (params) => {
  return service({
    url: '/customerUser/deleteCustomerUserByIds',
    method: 'delete',
    params
  })
}

// @Tags CustomerUser
// @Summary 更新客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CustomerUser true "更新客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customerUser/updateCustomerUser [put]
export const updateCustomerUser = (data) => {
  return service({
    url: '/customerUser/updateCustomerUser',
    method: 'put',
    data
  })
}

// @Tags CustomerUser
// @Summary 用id查询客户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CustomerUser true "用id查询客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customerUser/findCustomerUser [get]
export const findCustomerUser = (params) => {
  return service({
    url: '/customerUser/findCustomerUser',
    method: 'get',
    params
  })
}

// @Tags CustomerUser
// @Summary 分页获取客户管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取客户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customerUser/getCustomerUserList [get]
export const getCustomerUserList = (params) => {
  return service({
    url: '/customerUser/getCustomerUserList',
    method: 'get',
    params
  })
}

// @Tags CustomerUser
// @Summary 不需要鉴权的客户管理接口
// @Accept application/json
// @Produce application/json
// @Param data query managerReq.CustomerUserSearch true "分页获取客户管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /customerUser/getCustomerUserPublic [get]
export const getCustomerUserPublic = () => {
  return service({
    url: '/customerUser/getCustomerUserPublic',
    method: 'get',
  })
}
