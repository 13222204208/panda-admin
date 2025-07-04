<script setup lang="ts">
import { ref, reactive } from "vue";
import type { FormRules } from "element-plus";
import type { BatchFormProps } from "../utils/types";
import { Plus, Delete } from "@element-plus/icons-vue";

const props = withDefaults(defineProps<BatchFormProps>(), {
  formInline: () => ({
    dictType: "",
    title: "",
    dictItems: [
      {
        dictLabel: "",
        dictValue: "",
        sort: 1,
        remark: ""
      }
    ]
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

const rules = reactive<FormRules>({
  dictType: [
    { required: true, message: "字典类型为必填项", trigger: "blur" },
    {
      pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/,
      message: "字典类型必须以字母开头，只能包含字母、数字和下划线",
      trigger: "blur"
    }
  ],
  title: [
    { required: true, message: "字典标题为必填项", trigger: "blur" },
    {
      min: 2,
      max: 50,
      message: "字典标题长度应在2-50个字符之间",
      trigger: "blur"
    }
  ]
});

// 动态验证规则（删除状态验证）
const getDynamicRules = (index: number) => {
  return {
    [`dictItems.${index}.dictLabel`]: [
      { required: true, message: "字典标签为必填项", trigger: "blur" }
    ],
    [`dictItems.${index}.dictValue`]: [
      { required: true, message: "字典值为必填项", trigger: "blur" }
    ],
    [`dictItems.${index}.sort`]: [
      {
        type: "number",
        message: "排序值必须为数字",
        trigger: "blur",
        transform: (value: string) => Number(value)
      }
    ]
  };
};

// 添加字典项
const addDictItem = () => {
  newFormInline.value.dictItems.push({
    dictLabel: "",
    dictValue: "",
    sort: newFormInline.value.dictItems.length + 1,
    remark: ""
  });
};

// 删除字典项
const removeDictItem = (index: number) => {
  if (newFormInline.value.dictItems.length > 1) {
    newFormInline.value.dictItems.splice(index, 1);
    // 重新排序
    newFormInline.value.dictItems.forEach((item, idx) => {
      item.sort = idx + 1;
    });
  }
};

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
    label-width="80px"
    class="batch-form"
  >
    <!-- 字典类型 -->
    <div class="form-section">
      <el-row :gutter="24">
        <el-col :span="12">
          <el-form-item label="字典类型" prop="dictType" class="dict-type-item">
            <el-input
              v-model="newFormInline.dictType"
              clearable
              placeholder="请输入字典类型（如 sex、status）"
              style="width: 320px"
            />
            <div class="form-tip">
              字典类型用于分组管理，建议使用英文字母开头
            </div>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="字典标题" prop="title" class="dict-type-item">
            <el-input
              v-model="newFormInline.title"
              clearable
              placeholder="请输入字典标题（如 性别、状态）"
              style="width: 320px"
            />
            <div class="form-tip">字典标题用于显示，建议使用中文描述</div>
          </el-form-item>
        </el-col>
      </el-row>
    </div>

    <!-- 字典项列表 -->
    <el-divider content-position="left" class="section-divider">
      <span class="section-title">字典项配置</span>
    </el-divider>

    <div class="dict-items-container">
      <div
        v-for="(item, index) in newFormInline.dictItems"
        :key="index"
        class="dict-item-card"
      >
        <div class="dict-item-header">
          <span class="dict-item-title">字典项 {{ index + 1 }}</span>
          <el-button
            v-if="newFormInline.dictItems.length > 1"
            type="danger"
            :icon="Delete"
            size="small"
            circle
            @click="removeDictItem(index)"
            title="删除此项"
            class="delete-btn"
          />
        </div>

        <el-row :gutter="16" class="dict-item-content">
          <el-col :xs="24" :sm="12" :md="6">
            <el-form-item
              label="标签"
              :prop="`dictItems.${index}.dictLabel`"
              :rules="getDynamicRules(index)[`dictItems.${index}.dictLabel`]"
              class="form-item-compact"
            >
              <el-input
                v-model="item.dictLabel"
                clearable
                placeholder="如：男、启用"
              />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="6">
            <el-form-item
              label="值"
              :prop="`dictItems.${index}.dictValue`"
              :rules="getDynamicRules(index)[`dictItems.${index}.dictValue`]"
              class="form-item-compact"
            >
              <el-input
                v-model="item.dictValue"
                clearable
                placeholder="如：1、enabled"
              />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="4">
            <el-form-item
              label="排序"
              :prop="`dictItems.${index}.sort`"
              :rules="getDynamicRules(index)[`dictItems.${index}.sort`]"
              class="form-item-compact"
            >
              <el-input-number
                v-model="item.sort"
                :min="0"
                :max="9999"
                controls-position="right"
                class="w-full"
              />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="24" :md="8">
            <el-form-item label="备注" class="form-item-compact">
              <el-input
                v-model="item.remark"
                placeholder="备注信息（可选）"
                maxlength="100"
                show-word-limit
              />
            </el-form-item>
          </el-col>
        </el-row>
      </div>
    </div>

    <!-- 添加按钮 -->
    <div class="add-button-container">
      <el-button
        type="primary"
        :icon="Plus"
        @click="addDictItem"
        class="add-item-btn"
      >
        添加字典项
      </el-button>
    </div>
  </el-form>
</template>

<style scoped>
.batch-form {
  padding: 8px 0;
}

.form-section {
  margin-bottom: 24px;
}

.dict-type-item {
  margin-bottom: 8px;
}

.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
  line-height: 1.4;
}

.section-divider {
  margin: 24px 0 20px 0;
}

.section-title {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
}

.dict-items-container {
  margin-bottom: 20px;
}

.dict-item-card {
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 16px;
  background-color: #fafbfc;
  transition: all 0.3s ease;
  position: relative;
}

.dict-item-card:hover {
  border-color: #409eff;
  box-shadow: 0 2px 12px rgba(64, 158, 255, 0.12);
}

.dict-item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #ebeef5;
}

.dict-item-title {
  font-weight: 500;
  color: #303133;
  font-size: 14px;
}

.delete-btn {
  opacity: 0.7;
  transition: opacity 0.3s ease;
}

.delete-btn:hover {
  opacity: 1;
}

.dict-item-content {
  margin: 0;
}

.form-item-compact {
  margin-bottom: 12px;
}

.form-item-compact:last-child {
  margin-bottom: 0;
}

.add-button-container {
  margin-top: 16px;
  text-align: center;
}

.add-item-btn {
  width: 200px;
  height: 40px;
  border-style: dashed;
  background-color: #f8f9fa;
  border-color: #d0d7de;
  color: #656d76;
  font-size: 14px;
  transition: all 0.3s ease;
}

.add-item-btn:hover {
  background-color: #409eff;
  border-color: #409eff;
  color: white;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(64, 158, 255, 0.3);
}

/* 响应式优化 */
@media (max-width: 768px) {
  .dict-item-card {
    padding: 12px;
  }

  .dict-item-content .el-col {
    margin-bottom: 8px;
  }

  .add-item-btn {
    width: 100%;
  }
}

/* 表单项标签优化 */
:deep(.el-form-item__label) {
  font-weight: 500;
  color: #606266;
}

/* 输入框聚焦效果优化 */
:deep(.el-input__wrapper) {
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #c0c4cc inset;
}

:deep(.el-input.is-focus .el-input__wrapper) {
  box-shadow: 0 0 0 1px #409eff inset;
}
</style>
