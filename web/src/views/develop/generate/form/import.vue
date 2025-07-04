<script setup lang="ts">
import { ref, watch } from "vue";
import { ElMessage } from "element-plus";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import {
  getTableList,
  importTables,
  type TableItem,
  type ImportTableInfo
} from "@/api/generate";

// 图标导入
import Database from "~icons/ri/database-2-line";
import CheckLine from "~icons/ri/check-line";
import CloseLine from "~icons/ri/close-line";

interface Props {
  visible: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:visible", "success"]);

const tableRef = ref();
const loading = ref(false);
const tableData = ref<TableItem[]>([]);
const selectedTables = ref([]);

const columns = [
  {
    type: "selection",
    width: 55,
    align: "center",
    fixed: "left"
  },
  {
    label: "表名称",
    prop: "tableName",
    minWidth: 140,
    showOverflowTooltip: true
  },
  {
    label: "表描述",
    prop: "tableComment",
    minWidth: 160,
    showOverflowTooltip: true
  },
  {
    label: "创建时间",
    prop: "createTime",
    minWidth: 180,
    showOverflowTooltip: true
  },
  {
    label: "更新时间",
    prop: "updateTime",
    minWidth: 180,
    showOverflowTooltip: true
  }
];

const pagination = ref({
  total: 0,
  pageSize: 10,
  currentPage: 1,
  background: true
});

// 获取数据
const fetchData = async (
  currentPage = pagination.value.currentPage,
  pageSize = pagination.value.pageSize
) => {
  loading.value = true;
  try {
    const { data } = await getTableList({
      currentPage,
      pageSize
    });
    tableData.value = data.list;
    pagination.value.total = data.total;
    pagination.value.currentPage = data.currentPage;
    pagination.value.pageSize = pageSize;
  } catch (error) {
    ElMessage.error("获取数据库表列表失败");
  } finally {
    loading.value = false;
  }
};

// 分页变化处理
const handleCurrentChange = (currentPage: number) => {
  fetchData(currentPage, pagination.value.pageSize);
};

const handleSizeChange = (pageSize: number) => {
  fetchData(1, pageSize);
};

// 选择变化
const handleSelectionChange = selection => {
  selectedTables.value = selection;
};

// 取消选择
const onSelectionCancel = () => {
  selectedTables.value = [];
  // 清除表格的选中状态
  tableRef.value?.getTableRef()?.clearSelection();
};

// 确定导入
const onSubmit = async () => {
  if (selectedTables.value.length === 0) {
    ElMessage.warning("请选择要导入的表");
    return;
  }
  loading.value = true;
  try {
    // 只提取表名和表注释
    const tables: ImportTableInfo[] = selectedTables.value.map(item => ({
      tableName: item.tableName,
      tableComment: item.tableComment
    }));

    const { data } = await importTables({ tables });

    if (data.success) {
      ElMessage.success(
        `成功导入 ${data.count} 个表${data.errors && data.errors.length > 0 ? `，部分失败` : ""}`
      );
    }

    if (data.errors && data.errors.length > 0) {
      ElMessage.warning(`导入过程中出现错误：${data.errors.join(", ")}`);
    }

    emit("success");
    handleClose();
  } catch (error) {
    ElMessage.error("导入失败");
  } finally {
    loading.value = false;
  }
};

const handleClose = () => {
  selectedTables.value = [];
  emit("update:visible", false);
};

// 监听对话框可见性变化
watch(
  () => props.visible,
  newVal => {
    if (newVal) {
      fetchData();
    }
  }
);
</script>

<template>
  <el-dialog
    :model-value="props.visible"
    title="导入数据表"
    width="1000px"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @update:model-value="emit('update:visible', $event)"
    @close="handleClose"
  >
    <template #header>
      <div class="flex items-center">
        <el-icon class="mr-2 text-[var(--el-color-primary)]">
          <component :is="useRenderIcon(Database)" />
        </el-icon>
        <span class="text-lg font-medium">导入数据表</span>
      </div>
    </template>

    <!-- 选择提示栏 -->
    <div
      v-if="selectedTables.length > 0"
      v-motion-fade
      class="mb-4 pl-4 flex items-center bg-[var(--el-fill-color-light)] h-[46px] rounded-md"
    >
      <div class="flex-auto">
        <span
          class="text-[rgba(42,46,54,0.5)] dark:text-[rgba(220,220,242,0.5)]"
          style="font-size: var(--el-font-size-base)"
        >
          已选 {{ selectedTables.length }} 项
        </span>
        <el-button type="primary" text @click="onSelectionCancel">
          取消选择
        </el-button>
      </div>
      <div class="mr-4">
        <el-tag type="success" size="small">
          <el-icon class="mr-1">
            <component :is="useRenderIcon(CheckLine)" />
          </el-icon>
          {{ selectedTables.length }} 个表待导入
        </el-tag>
      </div>
    </div>

    <!-- 数据表格 -->
    <pure-table
      ref="tableRef"
      row-key="tableName"
      align-whole="center"
      table-layout="auto"
      :loading="loading"
      :data="tableData"
      :columns="columns"
      :pagination="pagination"
      :header-cell-style="{
        background: 'var(--el-fill-color-light)',
        color: 'var(--el-text-color-primary)',
        fontWeight: '500'
      }"
      :row-style="{ cursor: 'pointer' }"
      @selection-change="handleSelectionChange"
      @page-size-change="handleSizeChange"
      @page-current-change="handleCurrentChange"
    >
      <template #empty>
        <el-empty description="暂无数据表" :image-size="100">
          <el-button type="primary" @click="fetchData"> 刷新数据 </el-button>
        </el-empty>
      </template>
    </pure-table>

    <template #footer>
      <div class="flex items-center justify-between">
        <div class="flex items-center text-[var(--el-text-color-secondary)]">
          <el-icon class="mr-1">
            <component :is="useRenderIcon(Database)" />
          </el-icon>
          <span>共 {{ pagination.total }} 个数据表</span>
          <span v-if="selectedTables.length > 0" class="ml-4">
            ，已选择 {{ selectedTables.length }} 个
          </span>
        </div>
        <div class="flex items-center space-x-2">
          <el-button :icon="useRenderIcon(CloseLine)" @click="handleClose">
            取 消
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(CheckLine)"
            :loading="loading"
            :disabled="selectedTables.length === 0"
            @click="onSubmit"
          >
            {{
              loading
                ? "导入中..."
                : `确定导入${selectedTables.length > 0 ? `(${selectedTables.length})` : ""}`
            }}
          </el-button>
        </div>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
:deep(.el-dialog__header) {
  padding: 20px 20px 10px;
  border-bottom: 1px solid var(--el-border-color-lighter);
}

:deep(.el-dialog__body) {
  padding: 20px;
}

:deep(.el-dialog__footer) {
  padding: 10px 20px 20px;
  border-top: 1px solid var(--el-border-color-lighter);
}

// 表格行悬停效果
:deep(.el-table__row:hover) {
  background-color: var(--el-fill-color-light);
}

// 选中行样式
:deep(.el-table__row.current-row) {
  background-color: var(--el-color-primary-light-9);
}

// 加载状态优化
:deep(.el-loading-mask) {
  background-color: rgba(255, 255, 255, 0.8);
}

// 响应式设计
@media (max-width: 768px) {
  :deep(.el-dialog) {
    width: 95% !important;
    margin: 5vh auto;
  }
}
</style>
