<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import BasicConfig from "./components/BasicConfig.vue";
import FieldConfig from "./components/FieldConfig.vue";
import {
  getCodeGenRecordDetail,
  updateCodeGenRecord,
  CodeGenRecord
} from "@/api/generate";
import { message } from "@/utils/message";

defineOptions({
  name: "CodeConfig"
});

const route = useRoute();
const router = useRouter();
const activeTab = ref("basic");

// 获取路由传递的tableId
const tableId = ref<number | null>(null);

// 配置数据
const configData = ref({
  // 基本设置
  tableName: "",
  tableComment: "",
  moduleName: "",

  // 功能权限
  hasAdd: true,
  hasDelete: true,
  hasEdit: true,
  hasQuery: true,
  hasExport: false,
  hasBatchDelete: false,

  // 菜单设置
  parentMenuId: null,
  menuName: "",
  menuIcon: "",

  // 关联表设置
  joinTables: [
    {
      tableName: "", // 关联表名称
      alias: "", // 别名
      joinType: "LEFT", // 关联方式: LEFT, RIGHT, INNER
      joinField: "", // 关联字段
      mainField: "" // 主表关联字段
    }
  ]
});

// 字段配置数据
const fieldData = ref([
  {
    fieldName: "title", // 字段名
    fieldComment: "标题", // 字段描述
    htmlType: "input", // 表单组件
    dictType: "", // 绑定的字典
    validationRules: "", // 验证规则
    isEdit: true, // 编辑
    isRequired: true, // 必填
    isUnique: false, // 唯一
    isList: true, // 列表
    isQuery: false, // 查询
    queryType: "LIKE" // 查询方式
  }
]);

// 保存配置
const handleSaveConfig = async () => {
  if (!tableId.value) {
    message("缺少记录ID，无法保存配置", { type: "error" });
    return;
  }

  try {
    // 准备更新数据
    const updateData = {
      tableName: configData.value.tableName,
      tableComment: configData.value.tableComment,
      packageName: configData.value.moduleName, // 使用moduleName作为packageName
      moduleName: configData.value.moduleName,
      options: JSON.stringify({
        ...configData.value,
        fields: fieldData.value
      }),
      columns: JSON.stringify(fieldData.value)
    };

    // 调用更新接口
    const res = await updateCodeGenRecord(tableId.value, updateData);
    if (res.code === 0) {
      console.log("保存配置成功", {
        config: configData.value,
        fields: fieldData.value
      });
      message("配置保存成功", { type: "success" });
    } else {
      message("配置保存失败", { type: "error" });
    }
  } catch (error) {
    console.error("保存配置失败:", error);
    message("配置保存失败", { type: "error" });
  }
};

// 生成代码
const handleGenerateCode = async () => {
  if (!tableId.value) {
    message("缺少记录ID，无法生成代码", { type: "error" });
    return;
  }

  try {
    // 准备生成代码的数据
    const generateData = {
      tableName: configData.value.tableName,
      tableComment: configData.value.tableComment,
      packageName: configData.value.moduleName, // 使用moduleName作为packageName
      moduleName: configData.value.moduleName,
      options: JSON.stringify({
        ...configData.value,
        fields: fieldData.value
      }),
      columns: JSON.stringify(fieldData.value)
    };

    // 调用代码生成接口
    const res = await CodeGenRecord(tableId.value, generateData);
    if (res.code === 0) {
      console.log("代码生成成功", {
        config: configData.value,
        fields: fieldData.value
      });
      message("代码生成成功", { type: "success" });
    } else {
      message("代码生成失败", { type: "error" });
    }
  } catch (error) {
    console.error("代码生成失败:", error);
    message("代码生成失败", { type: "error" });
  }
};

// 获取代码生成记录详情
const loadRecordDetail = async () => {
  if (tableId.value) {
    try {
      const result = await getCodeGenRecordDetail(tableId.value);
      console.log("获取记录详情", result);
      if (result.data) {
        // 更新配置数据 - 解析JSON字符串
        configData.value =
          typeof result.data.options === "string"
            ? JSON.parse(result.data.options)
            : result.data.options;
        // 更新字段数据 - 解析JSON字符串
        fieldData.value =
          typeof result.data.columns === "string"
            ? JSON.parse(result.data.columns)
            : result.data.columns;
      }
    } catch (error) {
      console.error("获取代码生成记录详情失败:", error);
    }
  }
};

// 组件挂载时获取tableId并加载详情
onMounted(() => {
  // 从路由query参数中获取tableId
  const queryTableId = route.query.tableId;
  if (queryTableId) {
    tableId.value = Number(queryTableId);
    loadRecordDetail();
  }
});
// 返回列表
const handleBackToList = () => {
  router.push("/develop/generate");
};
</script>

<template>
  <div class="code-config-container">
    <el-card class="config-card">
      <template #header>
        <div class="card-header">
          <span class="card-title">代码生成配置</span>
          <div class="header-buttons">
            <el-button type="default" @click="handleBackToList"
              >返回列表</el-button
            >
            <el-button type="primary" @click="handleGenerateCode"
              >生成代码</el-button
            >
            <el-button type="info" @click="handleSaveConfig"
              >仅保存配置</el-button
            >
          </div>
        </div>
      </template>

      <!-- 标签页 -->
      <el-tabs v-model="activeTab" class="config-tabs">
        <!-- 基本信息 -->
        <el-tab-pane label="基本信息" name="basic">
          <BasicConfig v-model="configData" />
        </el-tab-pane>

        <!-- 主表字段 -->
        <el-tab-pane label="主表字段" name="fields">
          <FieldConfig v-model="fieldData" :table-name="configData.tableName" />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<style lang="scss" scoped>
.code-config-container {
  margin: 24px;
  height: calc(100vh - 120px);
}

.config-card {
  height: 100%;

  :deep(.el-card__body) {
    padding: 0;
    height: calc(100% - 60px);
    overflow-y: auto;
  }
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
}

.header-buttons {
  display: flex;
  gap: 8px;
}

.config-tabs {
  :deep(.el-tabs__header) {
    margin: 0;
    padding: 0 20px;
    background-color: var(--el-fill-color-extra-light);
    border-bottom: 1px solid var(--el-border-color-light);
  }

  :deep(.el-tabs__content) {
    padding: 20px;
  }
}
</style>
