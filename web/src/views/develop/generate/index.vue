<script setup lang="ts">
import { ref } from "vue";
import { useGenerate } from "./utils/hook";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import SqlPreview from "./form/sql.vue"; // 添加这行
import ImportTable from "./form/import.vue";
import Delete from "~icons/ep/delete";
import EditPen from "~icons/ep/edit-pen";
import Refresh from "~icons/ep/refresh";
import AddFill from "~icons/ri/add-circle-line";
import Download from "~icons/ep/download";
import View from "~icons/ep/view";

defineOptions({
  name: "CodeGenerate"
});

const formRef = ref();
const tableRef = ref();

const {
  form,
  loading,
  columns,
  dataList,
  selectedNum,
  pagination,
  deviceDetection,
  onSearch,
  resetForm,
  onbatchDel,
  handleCode,
  handleUpdate,
  handleDelete,
  handleGenerate,
  handlePreview,
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange
} = useGenerate(tableRef);

const sqlVisible = ref(false);
// 创建表
const onCreate = () => {
  sqlVisible.value = true;
};
// 导入表
const importVisible = ref(false);
const onImport = () => {
  console.log("导入表");
  importVisible.value = true;
};
</script>

<template>
  <div class="main">
    <el-form
      ref="formRef"
      :inline="true"
      :model="form"
      class="search-form bg-bg_color w-full pl-8 pt-[12px] overflow-auto"
    >
      <el-form-item label="表名：" prop="tableName">
        <el-input
          v-model="form.tableName"
          placeholder="请输入表名"
          clearable
          class="w-[180px]!"
        />
      </el-form-item>
      <el-form-item label="表描述：" prop="tableComment">
        <el-input
          v-model="form.tableComment"
          placeholder="请输入表描述"
          clearable
          class="w-[180px]!"
        />
      </el-form-item>
      <el-form-item label="创建时间：" prop="createTime">
        <el-date-picker
          v-model="form.createTime"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          class="w-[240px]!"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :icon="useRenderIcon('ri/search-line')"
          :loading="loading"
          @click="onSearch"
        >
          搜索
        </el-button>
        <el-button :icon="useRenderIcon(Refresh)" @click="resetForm(formRef)">
          重置
        </el-button>
      </el-form-item>
    </el-form>

    <PureTableBar title="代码生成" :columns="columns" @refresh="onSearch">
      <template #buttons>
        <el-button
          type="primary"
          :icon="useRenderIcon('ep:plus')"
          @click="onCreate"
        >
          创建表
        </el-button>
        <el-button
          type="primary"
          :icon="useRenderIcon(AddFill)"
          @click="onImport"
        >
          导入表
        </el-button>
      </template>
      <template v-slot="{ size, dynamicColumns }">
        <pure-table
          ref="tableRef"
          row-key="id"
          adaptive
          :adaptiveConfig="{ offsetBottom: 108 }"
          align-whole="center"
          table-layout="auto"
          :loading="loading"
          :size="size"
          :data="dataList"
          :columns="dynamicColumns"
          :pagination="{ ...pagination, size }"
          :header-cell-style="{
            background: 'var(--el-fill-color-light)',
            color: 'var(--el-text-color-primary)'
          }"
          @selection-change="handleSelectionChange"
          @page-size-change="handleSizeChange"
          @page-current-change="handleCurrentChange"
        >
          <template #operation="{ row }">
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(EditPen)"
              @click="handleCode(row)"
            >
              生成编辑
            </el-button>
            <el-popconfirm
              :title="`是否确认删除表名为${row.tableName}的这条数据`"
              @confirm="handleDelete(row)"
            >
              <template #reference>
                <el-button
                  class="reset-margin"
                  link
                  type="primary"
                  :size="size"
                  :icon="useRenderIcon(Delete)"
                >
                  删除
                </el-button>
              </template>
            </el-popconfirm>
          </template>
        </pure-table>
      </template>
    </PureTableBar>
    <SqlPreview v-model:visible="sqlVisible" />
    <ImportTable v-model:visible="importVisible" @success="onSearch" />
  </div>
</template>

<style lang="scss" scoped>
:deep(.el-button:focus-visible) {
  outline: none;
}

.main {
  margin: 24px 24px 0 !important;
}

.search-form {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
