<script setup lang="ts">
import { ref, computed, reactive } from "vue";
import type { FormRules } from "element-plus";
import type { FormProps } from "../utils/types";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    id: "",
    tableName: "",
    tableComment: "",
    className: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

// 根据表名自动生成类名
const autoGenerateClassName = computed(() => {
  if (!newFormInline.value.tableName) return "";

  // 将下划线命名转换为驼峰命名
  const tableName = newFormInline.value.tableName;
  const words = tableName.split("_");
  const className = words
    .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join("");

  return className;
});

// 监听表名变化，自动生成类名
const handleTableNameChange = () => {
  if (autoGenerateClassName.value && !newFormInline.value.className) {
    newFormInline.value.className = autoGenerateClassName.value;
  }
};

const rules = reactive<FormRules>({
  tableName: [
    { required: true, message: "表名称为必填项", trigger: "blur" },
    {
      pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/,
      message: "表名称必须以字母开头，只能包含字母、数字和下划线",
      trigger: "blur"
    }
  ],
  tableComment: [
    { required: true, message: "表描述为必填项", trigger: "blur" }
  ],
  className: [
    { required: true, message: "实体类名为必填项", trigger: "blur" },
    {
      pattern: /^[A-Z][a-zA-Z0-9]*$/,
      message: "实体类名必须以大写字母开头，只能包含字母和数字",
      trigger: "blur"
    }
  ]
});

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="rules"
    label-width="100px"
  >
    <el-row :gutter="30">
      <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
        <el-form-item label="表名称" prop="tableName">
          <el-input
            v-model="newFormInline.tableName"
            clearable
            placeholder="请输入表名称"
            @input="handleTableNameChange"
          />
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
        <el-form-item label="表描述" prop="tableComment">
          <el-input
            v-model="newFormInline.tableComment"
            clearable
            placeholder="请输入表描述"
          />
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
        <el-form-item label="实体类名" prop="className">
          <el-input
            v-model="newFormInline.className"
            clearable
            placeholder="请输入实体类名"
          >
            <template #append>
              <el-button
                v-if="
                  autoGenerateClassName &&
                  autoGenerateClassName !== newFormInline.className
                "
                @click="newFormInline.className = autoGenerateClassName"
                type="primary"
                size="small"
              >
                自动生成
              </el-button>
            </template>
          </el-input>
        </el-form-item>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24">
        <el-alert title="提示信息" type="info" :closable="false" show-icon>
          <template #default>
            <div class="text-sm">
              <p>
                • 表名称：数据库表的名称，建议使用下划线命名法（如：sys_user）
              </p>
              <p>• 表描述：对该表功能的简要说明</p>
              <p>
                •
                实体类名：生成的Java实体类名称，建议使用驼峰命名法（如：SysUser）
              </p>
              <p>• 系统会根据表名自动生成对应的实体类名</p>
            </div>
          </template>
        </el-alert>
      </el-col>
    </el-row>
  </el-form>
</template>

<style lang="scss" scoped>
:deep(.el-input-group__append) {
  padding: 0 8px;
}
</style>
