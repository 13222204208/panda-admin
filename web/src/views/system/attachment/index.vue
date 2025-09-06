<script setup lang="ts">
import { ref, onMounted } from "vue";
import { message } from "@/utils/message";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { PureTableBar } from "@/components/RePureTableBar";
import { useAttachment } from "./utils/hook";
import { uploadFile } from "@/api/attachment";
import { getKeyList, extractFields } from "@pureadmin/utils";
import type { UploadFile, UploadRequestOptions } from "element-plus";

import Upload from "~icons/ep/upload";
import Delete from "~icons/ep/delete";
import Refresh from "~icons/ep/refresh";
import EditPen from "~icons/ep/edit-pen";
import Download from "~icons/ep/download";

defineOptions({
  name: "AttachmentManagement"
});

const {
  form,
  loading,
  columns,
  dataList,
  pagination,
  selectedNum,
  onSearch,
  resetForm,
  openDialog,
  handleDelete,
  handleDownload,
  handleSizeChange,
  handleCurrentChange,
  handleSelectionChange,
  onbatchDel
} = useAttachment();

// 组件挂载时执行搜索
onMounted(() => {
  onSearch();
});

// 文件列表
const fileList = ref<UploadFile[]>([]);
const uploadRef = ref();
const formRef = ref();
const tableRef = ref();

// 上传配置 - 选择文件后直接上传
const uploadConfig = {
  action: "", // 不使用默认action，使用自定义上传
  multiple: true,
  showFileList: false, // 隐藏文件列表
  listType: "text" as const,
  accept: "image/*,.pdf,.doc,.docx,.xls,.xlsx,.ppt,.pptx,.txt,.zip,.rar",
  autoUpload: true, // 启用自动上传
  limit: 10, // 最大文件数量
  beforeUpload: onBeforeUpload,
  onExceed: onExceed
};

/** 上传文件前校验 */
function onBeforeUpload(file: File): boolean {
  // 文件大小限制 50MB
  const isValidSize = file.size / 1024 / 1024 < 50;
  if (!isValidSize) {
    message("文件大小不能超过 50MB!", { type: "error" });
    return false;
  }

  // 文件类型校验
  const allowedTypes = [
    "image/",
    "application/pdf",
    "application/msword",
    "application/vnd.openxmlformats-officedocument",
    "text/",
    "application/zip",
    "application/x-rar"
  ];

  const isValidType = allowedTypes.some(type => file.type.startsWith(type));
  if (!isValidType) {
    message("不支持的文件类型!", { type: "error" });
    return false;
  }

  return true;
}

/** 超出最大上传数时触发 */
function onExceed(): void {
  message(`最多上传 ${uploadConfig.limit} 个文件，请先删除部分文件再上传`, {
    type: "warning"
  });
}

/** 移除上传的文件 */
function handleRemove(file: UploadFile): void {
  const index = fileList.value.findIndex(item => item.uid === file.uid);
  if (index > -1) {
    fileList.value.splice(index, 1);
  }
}

/** 自定义上传函数 - 直接上传版本 */
async function handleUpload(options: UploadRequestOptions): Promise<void> {
  const { file, onProgress, onSuccess, onError } = options;

  try {
    // 显示上传进度
    onProgress?.({ percent: 0 } as any);

    const response = await uploadFile(file as File);

    if (response.code === 0) {
      onProgress?.({ percent: 100 } as any);
      onSuccess?.(response.data as any);
      message(`文件"${file.name}"上传成功`, { type: "success" });

      // 刷新列表
      await onSearch();
    } else {
      onError?.(new Error(response.message) as any);
      message(`文件"${file.name}"上传失败：${response.message}`, {
        type: "error"
      });
    }
  } catch (error) {
    console.error("上传文件异常:", error);
    onError?.(error as any);
    message(`文件"${file.name}"上传失败，请稍后重试`, { type: "error" });
  }
}

// 移除不需要的函数
// submitUpload 和 clearFiles 函数可以删除

/** 批量上传所有文件 */
function submitUpload(): void {
  if (fileList.value.length === 0) {
    message("请先选择要上传的文件", { type: "warning" });
    return;
  }

  uploadRef.value?.submit();
}

/** 清空文件列表 */
function clearFiles(): void {
  uploadRef.value?.clearFiles();
  fileList.value = [];
}
</script>

<template>
  <div class="main">
    <!-- 搜索表单 -->
    <el-form
      ref="formRef"
      :inline="true"
      :model="form"
      class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px]"
    >
      <el-form-item label="文件名：" prop="fileName">
        <el-input
          v-model="form.fileName"
          placeholder="请输入文件名"
          clearable
          class="w-[180px]!"
        />
      </el-form-item>
      <el-form-item label="后缀名：" prop="fileExt">
        <el-input
          v-model="form.fileExt"
          placeholder="请输入后缀名"
          clearable
          class="w-[180px]!"
        />
      </el-form-item>
      <el-form-item label="类别：" prop="isImage">
        <el-select
          v-model="form.isImage"
          placeholder="请选择类别"
          clearable
          class="w-[180px]!"
        >
          <el-option label="图片" :value="true" />
          <el-option label="文件" :value="false" />
        </el-select>
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

    <PureTableBar title="附件管理" :columns="columns" @refresh="onSearch">
      <template #buttons>
        <!-- 文件上传区域 - 简化版本 -->
        <div class="upload-container">
          <el-upload
            ref="uploadRef"
            v-bind="uploadConfig"
            :http-request="handleUpload"
            class="inline-block"
          >
            <el-button type="primary" :icon="useRenderIcon(Upload)">
              选择文件
            </el-button>
          </el-upload>
        </div>

        <!-- 批量删除按钮 -->
        <el-popconfirm
          v-if="selectedNum > 0"
          :title="`是否确认删除这${selectedNum}项`"
          @confirm="onbatchDel"
        >
          <template #reference>
            <el-button type="danger" :icon="useRenderIcon(Delete)">
              批量删除({{ selectedNum }})
            </el-button>
          </template>
        </el-popconfirm>
      </template>

      <template v-slot="{ size, dynamicColumns }">
        <pure-table
          ref="tableRef"
          align-whole="center"
          showOverflowTooltip
          table-layout="auto"
          :loading="loading"
          :size="size"
          adaptive
          :adaptiveConfig="{ offsetBottom: 108 }"
          :data="dataList"
          :columns="dynamicColumns"
          :pagination="{ ...pagination, size }"
          :header-cell-style="{
            background: 'var(--el-fill-color-light)',
            color: 'var(--el-text-color-primary)'
          }"
          row-key="id"
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
              @click="openDialog('编辑', row)"
            >
              编辑
            </el-button>
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(Download)"
              @click="handleDownload(row)"
            >
              下载
            </el-button>
            <el-popconfirm
              :title="`是否确认删除文件名为${row.fileName}的这条数据`"
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
  </div>
</template>

<style scoped lang="scss">
.upload-container {
  display: flex;
  gap: 8px;
  align-items: center;

  :deep(.el-upload-list) {
    max-height: 200px;
    overflow-y: auto;
  }
}
</style>
