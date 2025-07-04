<script setup lang="ts">
import { computed, ref, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { cloneDeep } from "lodash-es";
import Plus from "~icons/ep/plus";
import { IconSelect } from "@/components/ReIcon";
import { useCodeConfig } from "../utils/hook";

const { menuTreeData, fetchMenuTree, tableColumnsData, fetchTableColumns } =
  useCodeConfig();
interface JoinConfig {
  tableName: string; // 关联表名称
  alias: string; // 别名
  joinType: string; // 关联方式: LEFT, RIGHT, INNER
  joinField: string; // 关联字段
  mainField: string; // 主表关联字段
}

interface ConfigData {
  // 基本设置
  tableName: string;
  tableComment: string;
  moduleName: string;

  // 功能权限
  hasAdd: boolean;
  hasDelete: boolean;
  hasEdit: boolean;
  hasQuery: boolean;
  hasExport: boolean;
  hasBatchDelete: boolean;

  // 菜单设置
  parentMenuId: number | null;
  menuName: string;
  menuIcon: string;

  // 关联表设置
  joinTables: JoinConfig[];
}

const props = defineProps<{
  modelValue: ConfigData;
}>();

const emit = defineEmits<{
  "update:modelValue": [value: ConfigData];
}>();

const configData = computed({
  get: () => props.modelValue,
  set: value => emit("update:modelValue", value)
});

// 关联表相关数据
const linkTablesOption = ref<any[]>([]);
const linkColumnsOption = ref<any>({});
const columnsOption = ref<any[]>([]);

// 添加关联表
function addJoin() {
  if (!configData.value.joinTables) {
    configData.value.joinTables = [];
  }

  if (configData.value.joinTables.length >= 20) {
    ElMessage.warning("关联表数量不能超过20个");
    return;
  }

  configData.value.joinTables.push({
    tableName: "",
    alias: "",
    joinType: "LEFT",
    joinField: "",
    mainField: ""
  });
}

// 删除关联表
function delJoin(join: JoinConfig, index: number) {
  if (configData.value.joinTables) {
    // 恢复该表的可选状态
    const tableInfo = linkTablesOption.value.find(
      table => table.value === join.tableName
    );
    if (tableInfo) {
      tableInfo.disabled = false;
    }
    configData.value.joinTables.splice(index, 1);
  }
}

// 处理关联表选择变化
function handleTableChange(join: JoinConfig, tableName: string) {
  if (tableName) {
    const tableInfo = linkTablesOption.value.find(
      table => table.value === tableName
    );
    if (tableInfo) {
      // 使用表名的驼峰写法作为别名
      join.alias = tableToCamelCase(tableName);
      // 标记该表为已选择，避免重复选择
      tableInfo.disabled = true;
    }
    // 清空之前选择的字段
    join.joinField = "";
  } else {
    // 如果清空选择，恢复该表的可选状态
    const tableInfo = linkTablesOption.value.find(
      table => table.value === join.tableName
    );
    if (tableInfo) {
      tableInfo.disabled = false;
    }
    join.alias = "";
    join.joinField = "";
  }
}

// 将表名转换为驼峰命名
function tableToCamelCase(tableName: string): string {
  return tableName
    .split("_")
    .map((word, index) => {
      if (index === 0) {
        return word.toLowerCase();
      }
      return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
    })
    .join("");
}

// 获取指定表的字段选项
function getTableColumns(tableName: string) {
  if (!tableName || !tableColumnsData.value) {
    return [];
  }

  const table = tableColumnsData.value.find(t => t.tableName === tableName);
  if (table && table.columns) {
    return table.columns.map(column => ({
      label: `${column.columnName}(${column.columnComment})`,
      value: column.columnName,
      dataType: column.dataType,
      columnType: column.columnType,
      isNullable: column.isNullable,
      columnDefault: column.columnDefault,
      columnKey: column.columnKey,
      extra: column.extra
    }));
  }

  return [];
}
// 模拟API调用 - 替换为使用真实的表字段数据
async function loadTableSelect() {
  // 从 tableColumnsData 中获取表选项
  if (tableColumnsData.value && tableColumnsData.value.length > 0) {
    return tableColumnsData.value.map(table => ({
      label: `${table.tableName}(${table.tableComment})`,
      value: table.tableName,
      defAlias: tableToCamelCase(table.tableName),
      daoName: table.tableName
        .split("_")
        .map(word => word.charAt(0).toUpperCase() + word.slice(1))
        .join("")
    }));
  }
  return [];
}

async function loadColumnSelect(tableName: string) {
  // 从 tableColumnsData 中根据表名获取字段选项
  if (tableColumnsData.value && tableColumnsData.value.length > 0) {
    const table = tableColumnsData.value.find(t => t.tableName === tableName);
    if (table && table.columns) {
      return table.columns.map(column => ({
        label: `${column.columnName}(${column.columnComment})`,
        value: column.columnName,
        dataType: column.dataType,
        columnType: column.columnType,
        isNullable: column.isNullable,
        columnDefault: column.columnDefault,
        columnKey: column.columnKey,
        extra: column.extra
      }));
    }
  }
  return [];
}

// 初始化数据
onMounted(async () => {
  await fetchMenuTree();
  await fetchTableColumns();
  try {
    // 加载表选项
    // if (configData.value.dbName) {
    const tables = await loadTableSelect();
    linkTablesOption.value = cloneDeep(tables);
    // }

    // 加载主表字段
    if (configData.value.tableName) {
      columnsOption.value = await loadColumnSelect(configData.value.tableName);
    }

    // 加载已有关联表的字段选项
    if (configData.value.options?.join) {
      for (const join of configData.value.options.join) {
        if (join.linkTable) {
          linkColumnsOption.value[join.uuid] = await loadColumnSelect(
            join.linkTable
          );
        }
      }
    }
  } catch (error) {
    console.error("初始化数据失败:", error);
    ElMessage.error("初始化数据失败");
  }
});
</script>

<template>
  <div class="basic-config">
    <el-form :model="configData" label-width="120px" class="config-form">
      <div class="form-section">
        <div class="section-header">
          <span class="section-title">基本设置</span>
        </div>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="表名称" prop="tableName">
              <el-input v-model="configData.tableName" disabled />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="表描述" prop="tableComment">
              <el-input
                v-model="configData.tableComment"
                placeholder="请输入表描述"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="模块名" prop="moduleName">
              <el-input
                v-model="configData.moduleName"
                placeholder="请输入模块名"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 功能权限设置 -->
        <el-row :gutter="20" style="margin-top: 20px">
          <el-col :span="24">
            <el-form-item label="功能权限">
              <el-row :gutter="20">
                <el-col :span="4">
                  <el-checkbox v-model="configData.hasAdd">增加</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="configData.hasDelete">删除</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="configData.hasEdit">修改</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="configData.hasQuery" disabled
                    >查询</el-checkbox
                  >
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="configData.hasExport">导出</el-checkbox>
                </el-col>
                <el-col :span="4">
                  <el-checkbox v-model="configData.hasBatchDelete"
                    >批量删除</el-checkbox
                  >
                </el-col>
              </el-row>
            </el-form-item>
          </el-col>
        </el-row>
      </div>

      <!-- 菜单设置 -->
      <div class="form-section">
        <div class="section-header">
          <span class="section-title">菜单设置</span>
        </div>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="上级菜单">
              <el-cascader
                v-model="configData.parentMenuId"
                :options="menuTreeData"
                :props="{
                  value: 'id',
                  label: 'title',
                  children: 'children',
                  checkStrictly: true
                }"
                :emit-path="false"
                clearable
                filterable
                placeholder="请选择上级菜单"
                style="width: 100%"
                :show-all-levels="false"
                empty-text="暂无数据"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="菜单名称">
              <el-input
                v-model="configData.menuName"
                placeholder="请输入菜单名称"
                clearable
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="菜单图标">
              <IconSelect v-model="configData.menuIcon" class="w-full" />
            </el-form-item>
          </el-col>
        </el-row>
      </div>

      <!-- 关联表设置 -->
      <div class="form-section">
        <div class="section-header">
          <div class="section-title-wrapper">
            <span class="section-title">关联表设置</span>
            <el-button
              type="primary"
              size="small"
              @click="addJoin"
              :disabled="(configData.options?.join?.length || 0) >= 20"
            >
              <el-icon><Plus /></el-icon>
              新增关联表
            </el-button>
          </div>
        </div>

        <el-alert
          v-if="(configData.options?.join?.length || 0) > 3"
          type="warning"
          :show-icon="true"
          :closable="false"
          style="margin-bottom: 16px"
        >
          关联表数量建议在三个以下，过多的关联表可能影响查询性能
        </el-alert>

        <!-- 关联表设置部分修改 -->
        <div
          v-for="(join, index) in configData.joinTables"
          :key="index"
          class="join-item"
        >
          <div class="join-item-header">
            <span class="join-item-title">关联表 {{ index + 1 }}</span>
            <el-button
              @click="delJoin(join, index)"
              size="small"
              type="danger"
              text
            >
              移除
            </el-button>
          </div>

          <el-row :gutter="16">
            <el-col :span="6">
              <el-form-item label="关联表名称">
                <el-select
                  v-model="join.tableName"
                  placeholder="请选择关联表"
                  filterable
                  clearable
                  style="width: 100%"
                  @change="handleTableChange(join, $event)"
                >
                  <el-option
                    v-for="table in linkTablesOption"
                    :key="table.value"
                    :label="table.label"
                    :value="table.value"
                    :disabled="table.disabled"
                  />
                </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="4">
              <el-form-item label="别名">
                <el-input v-model="join.alias" placeholder="请输入别名" />
              </el-form-item>
            </el-col>

            <el-col :span="4">
              <el-form-item label="关联方式">
                <el-select
                  v-model="join.joinType"
                  placeholder="请选择"
                  style="width: 100%"
                >
                  <el-option label="LEFT JOIN" value="LEFT" />
                  <el-option label="RIGHT JOIN" value="RIGHT" />
                  <el-option label="INNER JOIN" value="INNER" />
                </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="5">
              <el-form-item label="关联字段">
                <el-select
                  v-model="join.joinField"
                  placeholder="请选择关联字段"
                  filterable
                  clearable
                  style="width: 100%"
                  :disabled="!join.tableName"
                >
                  <el-option
                    v-for="column in getTableColumns(join.tableName)"
                    :key="column.value"
                    :label="column.label"
                    :value="column.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="5">
              <el-form-item label="主表关联字段">
                <el-select
                  v-model="join.mainField"
                  placeholder="请选择主表字段"
                  filterable
                  clearable
                  style="width: 100%"
                >
                  <el-option
                    v-for="column in columnsOption"
                    :key="column.value"
                    :label="column.label"
                    :value="column.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
        </div>
      </div>
    </el-form>
  </div>
</template>

<style lang="scss" scoped>
.basic-config {
  .config-form {
    .form-section {
      margin-bottom: 24px;
      padding: 20px;
      background-color: var(--el-fill-color-extra-light);
      border-radius: 8px;
      border: 1px solid var(--el-border-color-light);
      transition: all 0.3s ease;

      &:hover {
        border-color: var(--el-border-color);
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
      }

      &:last-child {
        margin-bottom: 0;
      }
    }

    .section-header {
      margin-bottom: 20px;
      padding-bottom: 12px;
      border-bottom: 2px solid var(--el-color-primary-light-8);
      position: relative;
    }

    .section-title-wrapper {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .section-title {
      font-size: 16px;
      font-weight: 600;
      color: var(--el-text-color-primary);
      position: relative;

      &::before {
        content: "";
        position: absolute;
        left: -8px;
        top: 50%;
        transform: translateY(-50%);
        width: 4px;
        height: 16px;
        background-color: var(--el-color-primary);
        border-radius: 2px;
      }
    }
  }
}

.join-item {
  margin-bottom: 20px;
  padding: 20px;
  background-color: var(--el-fill-color-blank);
  border: 1px solid var(--el-border-color-lighter);
  border-radius: 8px;
  transition: all 0.3s ease;

  &:hover {
    border-color: var(--el-color-primary-light-7);
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  }

  &:last-child {
    margin-bottom: 0;
  }

  .join-item-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 16px;
    padding-bottom: 8px;
    border-bottom: 1px solid var(--el-border-color-extra-light);
  }

  .join-item-title {
    font-size: 14px;
    font-weight: 500;
    color: var(--el-text-color-regular);
  }
}

.empty-state {
  text-align: center;
  padding: 60px 0;
  background-color: var(--el-fill-color-blank);
  border-radius: 8px;
  border: 2px dashed var(--el-border-color-light);
}

.feedback-text {
  font-size: 12px;
  color: var(--el-color-danger);
  margin-top: 4px;
  line-height: 1.4;
}

:deep(.el-form-item) {
  margin-bottom: 18px;

  .el-form-item__label {
    font-weight: 500;
    color: var(--el-text-color-regular);
  }
}

:deep(.el-checkbox) {
  margin-right: 0;

  .el-checkbox__label {
    font-size: 14px;
    font-weight: 500;
  }
}

:deep(.el-alert) {
  .el-alert__content {
    font-size: 13px;
  }
}

:deep(.el-empty) {
  .el-empty__description {
    color: var(--el-text-color-placeholder);
  }
}
</style>
