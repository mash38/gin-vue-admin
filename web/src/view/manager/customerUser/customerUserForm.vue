
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入客户名称" />
</el-form-item>
        <el-form-item label="推荐人ID:" prop="parent_id">
    <el-input v-model.number="formData.parent_id" :clearable="true" placeholder="请输入推荐人ID" />
</el-form-item>
        <el-form-item label="客户类型:" prop="c_type">
    <el-input v-model.number="formData.c_type" :clearable="true" placeholder="请输入客户类型" />
</el-form-item>
        <el-form-item label="客户编码:" prop="code">
    <el-input v-model="formData.code" :clearable="true" placeholder="请输入客户编码" />
</el-form-item>
        <el-form-item label="客户来源:" prop="source">
    <el-input v-model.number="formData.source" :clearable="true" placeholder="请输入客户来源" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCustomerUser,
  updateCustomerUser,
  findCustomerUser
} from '@/api/manager/customerUser'

defineOptions({
    name: 'CustomerUserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            name: '',
            parent_id: undefined,
            c_type: undefined,
            code: '',
            source: undefined,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCustomerUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createCustomerUser(formData.value)
               break
             case 'update':
               res = await updateCustomerUser(formData.value)
               break
             default:
               res = await createCustomerUser(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
