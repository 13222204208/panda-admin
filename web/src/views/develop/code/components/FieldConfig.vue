<script lang="ts" setup>
import { computed, ref, onMounted } from "vue";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import Refresh from "~icons/ep/refresh";
import Sort from "~icons/ep/sort";
import { VueDraggable } from "vue-draggable-plus";
import { ElMessage } from "element-plus";
import { getDistinctDictTypes, type DictTypeInfo } from "@/api/dict";
import {
  getColumnConfigOptions,
  getTableColumns,
  type ColumnConfigOptionsData
} from "@/api/generate";

interface FieldData {
  id: number; // 字段ID
  fieldName: string; // 字段名
  fieldComment: string; // 字段描述
  htmlType: string; // 表单组件
  dictType: string; // 绑定的字典
  validationRules: string; // 验证规则
  isEdit: boolean; // 编辑
  isRequired: boolean; // 必填
  isUnique: boolean; // 唯一
  isList: boolean; // 列表
  isQuery: boolean; // 查询
  queryType: string; // 查询方式
}

const props = defineProps<{
  modelValue: FieldData[];
  tableName?: string; // 添加表名属性
}>();

const emit = defineEmits<{
  "update:modelValue": [value: FieldData[]];
}>();

const fieldData = computed({
  get: () => props.modelValue,
  set: value => emit("update:modelValue", value)
});

// 字典类型数据
const dictTypes = ref<DictTypeInfo[]>([]);
const dictTypesLoading = ref(false);

// 配置选项数据
const configOptions = ref<ColumnConfigOptionsData>({
  formModes: [],
  formValidations: [],
  whereModes: []
});
const configOptionsLoading = ref(false);

// 弹窗相关状态
const sortDialogVisible = ref(false);
const sortFieldList = ref<FieldData[]>([]);

// 获取字典类型数据
const loadDictTypes = async () => {
  try {
    dictTypesLoading.value = true;
    const response = await getDistinctDictTypes();
    console.log("响应数据", response);
    if (response.code === 0) {
      dictTypes.value = response.data.types || [];
    } else {
      ElMessage.error("获取字典类型失败");
    }
  } catch (error) {
    console.error("获取字典类型失败:", error);
    ElMessage.error("获取字典类型失败");
  } finally {
    dictTypesLoading.value = false;
  }
};

// 获取配置选项数据
const loadConfigOptions = async () => {
  try {
    configOptionsLoading.value = true;
    const response = await getColumnConfigOptions();
    console.log("配置选项响应数据", response);
    if (response.code === 0) {
      configOptions.value = response.data;
    } else {
      ElMessage.error("获取配置选项失败");
    }
  } catch (error) {
    console.error("获取配置选项失败:", error);
    ElMessage.error("获取配置选项失败");
  } finally {
    configOptionsLoading.value = false;
  }
};

// 组件挂载时加载数据
onMounted(() => {
  loadDictTypes();
  loadConfigOptions();
});

// 同步字段
const syncFields = async () => {
  if (!props.tableName) {
    ElMessage.warning("请先选择表名");
    return;
  }

  try {
    ElMessage.info("正在同步字段信息...");

    // 调用接口获取表字段信息
    const response = await getTableColumns(props.tableName);

    if (response.code === 0 && response.data) {
      const tableColumns = response.data.columns;
      const currentFields = fieldData.value;
      const newFields: FieldData[] = [];

      // 遍历数据库字段，检查是否有新字段需要添加
      tableColumns.forEach((column, index) => {
        // 检查当前字段列表中是否已存在该字段
        const existingField = currentFields.find(
          field => field.fieldName === column.columnName
        );

        if (!existingField) {
          // 如果是新字段，则添加到列表中
          const newField: FieldData = {
            id: Date.now() + index, // 生成唯一ID
            fieldName: column.columnName,
            fieldComment: column.columnComment || column.columnName,
            htmlType: getDefaultHtmlType(column.dataType), // 根据数据类型推断表单组件
            dictType: "",
            validationRules: "",
            isEdit: !isSystemField(column.columnName), // 系统字段默认不可编辑
            isRequired:
              column.isNullable === "NO" && !isSystemField(column.columnName),
            isUnique: column.columnKey === "UNI",
            isList: !isSystemField(column.columnName), // 系统字段默认不在列表显示
            isQuery: false,
            queryType: getDefaultQueryType(column.dataType)
          };
          newFields.push(newField);
        }
      });

      if (newFields.length > 0) {
        // 将新字段添加到现有字段列表中
        fieldData.value = [...currentFields, ...newFields];
        ElMessage.success(`成功同步 ${newFields.length} 个新字段`);
      } else {
        ElMessage.info("没有发现新字段");
      }
    } else {
      ElMessage.error("获取表字段信息失败");
    }
  } catch (error) {
    console.error("同步字段失败:", error);
    ElMessage.error("同步字段失败，请重试");
  }
};

// 根据数据类型推断默认的表单组件类型
const getDefaultHtmlType = (dataType: string): string => {
  const type = dataType.toLowerCase();

  if (type.includes("text") || type.includes("longtext")) {
    return "textarea";
  } else if (
    type.includes("int") ||
    type.includes("decimal") ||
    type.includes("float")
  ) {
    return "number";
  } else if (type.includes("date") || type.includes("time")) {
    return "datetime";
  } else if (type.includes("enum")) {
    return "select";
  } else {
    return "input";
  }
};

// 根据数据类型推断默认的查询方式
const getDefaultQueryType = (dataType: string): string => {
  const type = dataType.toLowerCase();

  if (
    type.includes("int") ||
    type.includes("decimal") ||
    type.includes("float")
  ) {
    return "EQ";
  } else if (type.includes("date") || type.includes("time")) {
    return "BETWEEN";
  } else {
    return "LIKE";
  }
};

// 判断是否为系统字段（如创建时间、更新时间等）
const isSystemField = (fieldName: string): boolean => {
  const systemFields = [
    "id",
    "created_at",
    "updated_at",
    "deleted_at",
    "create_time",
    "update_time",
    "delete_time",
    "created_by",
    "updated_by",
    "deleted_by"
  ];
  return systemFields.includes(fieldName.toLowerCase());
};

// 打开字段排序弹窗
const openSortDialog = () => {
  if (fieldData.value.length === 0) {
    ElMessage.warning("暂无字段可排序");
    return;
  }
  // 深拷贝字段数据用于排序
  sortFieldList.value = JSON.parse(JSON.stringify(fieldData.value));
  sortDialogVisible.value = true;
};

// 确认排序
const confirmSort = () => {
  // 更新sort字段
  sortFieldList.value.forEach((field, index) => {
    field.sort = (index + 1) * 10;
  });

  // 应用排序结果
  fieldData.value = [...sortFieldList.value];
  sortDialogVisible.value = false;
  ElMessage.success("字段排序已更新");
};

// 取消排序
const cancelSort = () => {
  sortDialogVisible.value = false;
};

// 拖拽结束事件
const onDragEnd = () => {
  console.log("拖拽结束", sortFieldList.value);
};
</script>

<template>
  <div class="field-config">
    <div class="field-header">
      <div class="field-actions">
        <el-button
          type="default"
          :icon="useRenderIcon(Refresh)"
          @click="syncFields"
        >
          同步字段
        </el-button>
        <el-button
          type="warning"
          :icon="useRenderIcon(Sort)"
          @click="openSortDialog"
        >
          移动字段
        </el-button>
      </div>
    </div>

    <div class="field-table">
      <el-table
        :data="fieldData"
        border
        style="width: 100%"
        table-layout="auto"
        :flexible="true"
      >
        <el-table-column
          type="index"
          label="#"
          min-width="50"
          align="center"
          fixed="left"
        />

        <el-table-column label="字段列名" max-width="100" show-overflow-tooltip>
          <template #default="{ row, $index }">
            <span class="field-name-text">
              {{ row.fieldName || "未设置" }}
            </span>
          </template>
        </el-table-column>

        <el-table-column label="字段描述" min-width="140" show-overflow-tooltip>
          <template #default="{ row, $index }">
            <el-input
              v-model="row.fieldComment"
              placeholder="请输入描述"
              size="small"
            />
          </template>
        </el-table-column>

        <el-table-column label="表单组件" min-width="120" show-overflow-tooltip>
          <template #default="{ row, $index }">
            <el-select
              v-model="row.htmlType"
              size="small"
              style="width: 100%"
              placeholder="选择组件"
              :loading="configOptionsLoading"
            >
              <el-option
                v-for="option in configOptions.formModes"
                :key="option.value"
                :label="option.label"
                :value="option.value"
              />
            </el-select>
          </template>
        </el-table-column>

        <el-table-column label="绑定字典" min-width="120" show-overflow-tooltip>
          <template #default="{ row, $index }">
            <el-select
              v-model="row.dictType"
              size="small"
              style="width: 100%"
              placeholder="选择字典类型"
              clearable
              filterable
              :loading="dictTypesLoading"
            >
              <el-option
                v-for="dictType in dictTypes"
                :key="dictType.dictType"
                :label="`${dictType.title} (${dictType.dictType})`"
                :value="dictType.dictType"
              />
            </el-select>
          </template>
        </el-table-column>

        <el-table-column label="验证规则" min-width="120" show-overflow-tooltip>
          <template #default="{ row, $index }">
            <el-select
              v-model="row.validationRules"
              size="small"
              style="width: 100%"
              placeholder="选择规则"
              :loading="configOptionsLoading"
            >
              <el-option
                v-for="option in configOptions.formValidations"
                :key="option.value"
                :label="option.label"
                :value="option.value"
              />
            </el-select>
          </template>
        </el-table-column>

        <el-table-column label="编辑" min-width="60" align="center">
          <template #default="{ row, $index }">
            <el-checkbox v-model="row.isEdit" />
          </template>
        </el-table-column>

        <el-table-column label="必填" min-width="60" align="center">
          <template #default="{ row, $index }">
            <el-checkbox v-model="row.isRequired" />
          </template>
        </el-table-column>

        <el-table-column label="唯一" min-width="60" align="center">
          <template #default="{ row, $index }">
            <el-checkbox v-model="row.isUnique" />
          </template>
        </el-table-column>

        <el-table-column label="列表" min-width="60" align="center">
          <template #default="{ row, $index }">
            <el-checkbox v-model="row.isList" />
          </template>
        </el-table-column>

        <el-table-column label="查询" min-width="60" align="center">
          <template #default="{ row, $index }">
            <el-checkbox v-model="row.isQuery" />
          </template>
        </el-table-column>

        <el-table-column label="查询方式" min-width="110" show-overflow-tooltip>
          <template #default="{ row, $index }">
            <el-select
              v-model="row.queryType"
              size="small"
              style="width: 100%"
              placeholder="查询方式"
              :loading="configOptionsLoading"
            >
              <el-option
                v-for="option in configOptions.whereModes"
                :key="option.value"
                :label="option.label"
                :value="option.value"
              />
            </el-select>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 字段排序弹窗 -->
    <el-dialog
      v-model="sortDialogVisible"
      title="字段排序"
      width="500px"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <div class="sort-dialog-content">
        <div class="sort-tips">
          <span class="tips-text">拖拽调整字段显示顺序</span>
        </div>

        <div class="sort-list">
          <VueDraggable
            v-model="sortFieldList"
            :animation="150"
            ghost-class="ghost-item"
            chosen-class="chosen-item"
            drag-class="drag-item"
            @end="onDragEnd"
          >
            <div
              v-for="(field, index) in sortFieldList"
              :key="field.id"
              class="sort-item"
            >
              <div class="sort-item-content">
                <div class="drag-handle">
                  <el-icon><Sort /></el-icon>
                </div>
                <div class="field-info">
                  <span class="field-name">
                    {{ field.fieldName || `字段${index + 1}` }}
                  </span>
                  <span class="field-comment" v-if="field.fieldComment">
                    {{ field.fieldComment }}
                  </span>
                </div>
                <div class="field-type">
                  {{ field.fieldType }}
                </div>
                <div class="sort-index">
                  {{ index + 1 }}
                </div>
              </div>
            </div>
          </VueDraggable>
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="cancelSort">取消</el-button>
          <el-button type="primary" @click="confirmSort">确认排序</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style lang="scss" scoped>
.field-config {
  .field-header {
    margin-bottom: 16px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .field-actions {
    display: flex;
    gap: 8px;
  }

  .field-table {
    // 表格容器设置
    overflow-x: auto;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);

    :deep(.el-table) {
      // 表格自适应布局
      table-layout: auto;
      min-width: 1200px; // 设置最小宽度确保内容不会过度压缩

      // 表头样式
      .el-table__header-wrapper {
        .el-table__header {
          th {
            background-color: var(--el-fill-color-light);
            color: var(--el-text-color-primary);
            font-weight: 600;
            padding: 12px 8px;
            border-bottom: 2px solid var(--el-border-color);

            // 表头文字不换行
            .cell {
              white-space: nowrap;
              overflow: hidden;
              text-overflow: ellipsis;
            }
          }
        }
      }

      // 表格行样式
      .el-table__body-wrapper {
        .el-table__row {
          &:hover {
            background-color: var(--el-fill-color-extra-light);
          }

          td {
            padding: 8px;
            border-bottom: 1px solid var(--el-border-color-lighter);

            .cell {
              padding: 0 4px;
            }
          }
        }
      }

      // 输入框样式
      .el-input {
        .el-input__wrapper {
          box-shadow: none;
          border: 1px solid var(--el-border-color-light);
          border-radius: 4px;
          transition: border-color 0.3s ease;

          &:hover {
            border-color: var(--el-border-color);
          }

          &.is-focus {
            border-color: var(--el-color-primary);
          }
        }
      }

      // 选择框样式
      .el-select {
        width: 100%;

        .el-input {
          .el-input__wrapper {
            border: 1px solid var(--el-border-color-light);

            &:hover {
              border-color: var(--el-border-color);
            }
          }
        }
      }

      // 复选框样式
      .el-checkbox {
        .el-checkbox__input {
          .el-checkbox__inner {
            border-radius: 3px;
          }
        }
      }

      // 数字输入框样式
      .el-input-number {
        width: 100%;
      }

      // 固定列阴影
      .el-table__fixed-left {
        box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
      }

      .el-table__fixed-right {
        box-shadow: -2px 0 8px rgba(0, 0, 0, 0.1);
      }
    }
  }
}

// 排序弹窗样式
.sort-dialog-content {
  .sort-tips {
    margin-bottom: 12px;

    .tips-text {
      font-size: 13px;
      color: var(--el-text-color-regular);
    }
  }

  .sort-list {
    max-height: 350px;
    overflow-y: auto;
    border: 1px solid var(--el-border-color-lighter);
    border-radius: 4px;
    padding: 4px;
    background-color: var(--el-fill-color-extra-light);
  }
}

.sort-item {
  margin-bottom: 4px;

  &:last-child {
    margin-bottom: 0;
  }

  .sort-item-content {
    display: flex;
    align-items: center;
    padding: 8px 10px;
    background-color: var(--el-fill-color-blank);
    border: 1px solid var(--el-border-color-lighter);
    border-radius: 4px;
    cursor: move;
    transition: all 0.2s ease;
    min-height: 36px;

    &:hover {
      border-color: var(--el-color-primary-light-5);
      background-color: var(--el-color-primary-light-9);
    }

    .drag-handle {
      margin-right: 8px;
      color: var(--el-text-color-placeholder);
      font-size: 14px;
      cursor: grab;
      display: flex;
      align-items: center;

      &:active {
        cursor: grabbing;
      }
    }

    .field-info {
      flex: 1;
      display: flex;
      align-items: center;
      gap: 8px;
      min-width: 0;

      .field-name {
        font-size: 13px;
        font-weight: 500;
        color: var(--el-text-color-primary);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }

      .field-comment {
        font-size: 11px;
        color: var(--el-text-color-regular);
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        opacity: 0.8;

        &::before {
          content: "(";
        }

        &::after {
          content: ")";
        }
      }
    }

    .field-type {
      margin-right: 10px;
      padding: 1px 6px;
      background-color: var(--el-color-primary-light-8);
      color: var(--el-color-primary-dark-2);
      border-radius: 3px;
      font-size: 11px;
      white-space: nowrap;
    }

    .sort-index {
      width: 24px;
      height: 24px;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: var(--el-color-primary);
      color: white;
      border-radius: 50%;
      font-size: 12px;
      font-weight: 500;
      flex-shrink: 0;
    }
  }
}

// 拖拽状态样式
.ghost-item {
  opacity: 0.4;
  background-color: var(--el-color-primary-light-9) !important;
}

.chosen-item {
  transform: scale(1.01);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

.drag-item {
  transform: rotate(2deg);
  opacity: 0.9;
}

// 响应式设计
@media (max-width: 1400px) {
  .field-table {
    :deep(.el-table) {
      min-width: 1000px;
    }
  }
}

@media (max-width: 1200px) {
  .field-table {
    :deep(.el-table) {
      min-width: 900px;

      .el-table__header-wrapper,
      .el-table__body-wrapper {
        .cell {
          font-size: 13px;
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .field-config {
    .field-header {
      flex-direction: column;
      gap: 12px;
      align-items: stretch;
    }

    .field-actions {
      justify-content: center;
      flex-wrap: wrap;
    }
  }

  .field-table {
    :deep(.el-table) {
      min-width: 800px;

      .el-input,
      .el-select {
        .el-input__wrapper {
          .el-input__inner {
            font-size: 12px;
          }
        }
      }
    }
  }

  // 弹窗在移动端的适配
  :deep(.el-dialog) {
    width: 95% !important;
    margin: 5vh auto;
  }

  .sort-item {
    .sort-item-content {
      padding: 8px;

      .field-info {
        .field-name {
          font-size: 13px;
        }

        .field-comment {
          font-size: 11px;
        }
      }

      .field-type {
        font-size: 11px;
        margin-right: 8px;
      }

      .sort-index {
        width: 28px;
        height: 28px;
        font-size: 12px;
      }
    }
  }
}
.field-name-text {
  color: var(--el-text-color-regular);
  font-size: 13px;
  padding: 4px 0;
  display: inline-block;
  width: 100%;

  &:empty::before {
    content: "未设置";
    color: var(--el-text-color-placeholder);
  }
}
</style>
