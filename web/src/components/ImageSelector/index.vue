<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { message } from "@/utils/message";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { addDialog } from "@/components/ReDialog";
import { getAttachmentList } from "@/api/attachment";
import type { AttachmentInfo } from "@/api/attachment";
import { getFileUrl } from "@/utils/env";

import Search from "~icons/ep/search";
import Picture from "~icons/ep/picture";
import Document from "~icons/ep/document";
import Check from "~icons/ep/check";
import Close from "~icons/ep/close";
import Refresh from "~icons/ep/refresh";

interface FileImageSelectorProps {
  /** 是否支持多选，默认false */
  multiple?: boolean;
  /** 最大选择数量，仅在multiple为true时生效 */
  maxCount?: number;
  /** 已选择的文件列表 - 修改为支持字符串和字符串数组 */
  modelValue?: string[] | string | null;
  /** 文件类型：'image' | 'file' | 'all'，默认'all' */
  fileType?: "image" | "file" | "all";
  /** 允许的文件扩展名，如 ['jpg', 'png', 'pdf'] */
  allowedExtensions?: string[];
}

interface FileImageSelectorEmits {
  (e: "update:modelValue", value: string[] | string | null): void;
  (e: "confirm", value: string[] | string | null): void;
}

const props = withDefaults(defineProps<FileImageSelectorProps>(), {
  multiple: false,
  maxCount: 9,
  fileType: "all"
});

const emit = defineEmits<FileImageSelectorEmits>();

// 弹窗状态
const dialogVisible = ref(false);
const loading = ref(false);

// 搜索表单 - 参考 attachment 页面
const searchForm = ref({
  fileName: "",
  fileExt: "", // 文件扩展名
  isImage:
    props.fileType === "image"
      ? true
      : props.fileType === "file"
        ? false
        : undefined // 根据props设置默认值
});

// 文件类型选项
const fileTypeOptions = computed(() => {
  const options = [];
  if (props.fileType === "all" || props.fileType === "image") {
    options.push({ label: "图片", value: true });
  }
  if (props.fileType === "all" || props.fileType === "file") {
    options.push({ label: "文件", value: false });
  }
  return options;
});

// 文件扩展名选项
const extensionOptions = computed(() => {
  if (props.allowedExtensions && props.allowedExtensions.length > 0) {
    return props.allowedExtensions.map(ext => ({
      label: ext.toUpperCase(),
      value: ext.toLowerCase()
    }));
  }

  // 默认选项
  const imageExts = ["jpg", "jpeg", "png", "gif", "webp", "svg"];
  const fileExts = [
    "pdf",
    "doc",
    "docx",
    "xls",
    "xlsx",
    "ppt",
    "pptx",
    "txt",
    "zip",
    "rar"
  ];

  let exts = [];
  if (props.fileType === "image") {
    exts = imageExts;
  } else if (props.fileType === "file") {
    exts = fileExts;
  } else {
    exts = [...imageExts, ...fileExts];
  }

  return exts.map(ext => ({
    label: ext.toUpperCase(),
    value: ext.toLowerCase()
  }));
});

// 格式化文件大小
const formatBytes = (bytes: number): string => {
  if (bytes === 0) return "0 B";
  const k = 1024;
  const sizes = ["B", "KB", "MB", "GB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
};

// 重置搜索 - 参考 attachment 页面的 resetForm
const handleReset = () => {
  searchForm.value = {
    fileName: "",
    fileExt: "",
    isImage:
      props.fileType === "image"
        ? true
        : props.fileType === "file"
          ? false
          : undefined
  };
  pagination.value.currentPage = 1;
  getFileList();
};

// 分页配置
const pagination = ref({
  currentPage: 1,
  pageSize: 12,
  total: 0
});

// 文件列表
const fileList = ref<AttachmentInfo[]>([]);

// 选中的文件
const selectedFiles = ref<AttachmentInfo[]>([]);

// 计算属性
const isMaxSelected = computed(() => {
  return props.multiple && selectedFiles.length >= props.maxCount;
});

// 判断是否为图片文件
const isImageFile = (file: AttachmentInfo): boolean => {
  const imageExts = ["jpg", "jpeg", "png", "gif", "webp", "svg", "bmp", "ico"];
  const ext = file.fileName.split(".").pop()?.toLowerCase();
  return imageExts.includes(ext || "");
};

// 获取文件图标
const getFileIcon = (file: AttachmentInfo) => {
  if (isImageFile(file)) {
    return Picture;
  }
  return Document;
};

// 获取文件列表
const getFileList = async () => {
  loading.value = true;
  try {
    const params = {
      ...searchForm.value,
      currentPage: pagination.value.currentPage,
      pageSize: pagination.value.pageSize
    };

    const { data } = await getAttachmentList(params);
    fileList.value = data.list;
    pagination.value.total = data.total;
  } catch (error) {
    console.error("获取文件列表失败:", error);
    message("获取文件列表失败", { type: "error" });
  } finally {
    loading.value = false;
  }
};

// 搜索
const handleSearch = () => {
  pagination.value.currentPage = 1;
  getFileList();
};

// 分页处理
const handleSizeChange = (val: number) => {
  pagination.value.pageSize = val;
  getFileList();
};

const handleCurrentChange = (val: number) => {
  pagination.value.currentPage = val;
  getFileList();
};

// 选择文件
const handleSelectFile = (file: AttachmentInfo) => {
  if (props.multiple) {
    const index = selectedFiles.value.findIndex(item => item.id === file.id);
    if (index > -1) {
      selectedFiles.value.splice(index, 1);
    } else {
      if (selectedFiles.value.length < props.maxCount) {
        selectedFiles.value.push(file);
      } else {
        message(`最多只能选择${props.maxCount}个文件`, { type: "warning" });
      }
    }
  } else {
    selectedFiles.value = [file];
  }
};

// 判断文件是否被选中
const isFileSelected = (file: AttachmentInfo) => {
  return selectedFiles.value.some(item => item.id === file.id);
};

// 确认选择 - 修改返回值为fileUrl
const handleConfirm = () => {
  if (selectedFiles.value.length === 0) {
    message("请选择文件", { type: "warning" });
    return;
  }

  // 只返回fileUrl，单个返回字符串，多个返回字符串数组
  const result = props.multiple
    ? selectedFiles.value.map(file => file.fileUrl)
    : selectedFiles.value[0].fileUrl;
  console.log("result", result);
  emit("update:modelValue", result);
  emit("confirm", result);
  dialogVisible.value = false;
};

// 打开弹窗 - 修改初始化逻辑以支持字符串格式
const openDialog = () => {
  // 初始化选中状态 - 需要根据fileUrl匹配
  selectedFiles.value = [];

  if (props.modelValue) {
    // 延迟初始化，等文件列表加载完成后再匹配
    dialogVisible.value = true;
    getFileList().then(() => {
      if (Array.isArray(props.modelValue)) {
        // 多选模式：根据fileUrl匹配文件
        selectedFiles.value = fileList.value.filter(file =>
          props.modelValue.includes(file.fileUrl)
        );
      } else if (typeof props.modelValue === "string") {
        // 单选模式：根据fileUrl匹配文件
        const matchedFile = fileList.value.find(
          file => file.fileUrl === props.modelValue
        );
        if (matchedFile) {
          selectedFiles.value = [matchedFile];
        }
      }
    });
  } else {
    dialogVisible.value = true;
    getFileList();
  }
};

// 取消选择
const handleCancel = () => {
  dialogVisible.value = false;
  selectedFiles.value = [];
};

// 删除这个重复的 openDialog 函数定义（第283-295行）
// const openDialog = () => {
//   // 初始化选中状态
//   if (props.modelValue) {
//     selectedFiles.value = Array.isArray(props.modelValue)
//       ? [...props.modelValue]
//       : [props.modelValue];
//   } else {
//     selectedFiles.value = [];
//   }
//
//   dialogVisible.value = true;
//   getFileList();
// };

// 暴露方法
defineExpose({
  openDialog
});

// 组件挂载时获取数据
onMounted(() => {
  // 不在挂载时获取数据，只在打开弹窗时获取
});
</script>

<template>
  <div class="file-image-selector">
    <!-- 触发按钮插槽 -->
    <slot name="trigger" :open="openDialog">
      <el-button
        type="primary"
        :icon="useRenderIcon(props.fileType === 'file' ? Document : Picture)"
        @click="openDialog"
      >
        {{
          props.fileType === "file"
            ? "选择文件"
            : props.fileType === "image"
              ? "选择图片"
              : "选择文件/图片"
        }}
      </el-button>
    </slot>

    <!-- 文件选择弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      :title="
        props.fileType === 'file'
          ? '选择文件'
          : props.fileType === 'image'
            ? '选择图片'
            : '选择文件/图片'
      "
      width="80%"
      :close-on-click-modal="false"
      class="file-image-selector-dialog"
    >
      <!-- 搜索区域 - 参考 attachment 页面样式 -->
      <el-form
        :inline="true"
        :model="searchForm"
        class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] mb-4"
      >
        <el-form-item label="文件名：" prop="fileName">
          <el-input
            v-model="searchForm.fileName"
            placeholder="请输入文件名"
            clearable
            class="w-[180px]!"
            @keyup.enter="handleSearch"
          />
        </el-form-item>

        <el-form-item label="文件格式：" prop="fileExt">
          <el-select
            v-model="searchForm.fileExt"
            placeholder="请选择格式"
            clearable
            class="w-[180px]!"
          >
            <el-option
              v-for="option in extensionOptions"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item
          v-if="fileTypeOptions.length > 1"
          label="类别："
          prop="isImage"
        >
          <el-select
            v-model="searchForm.isImage"
            placeholder="请选择类别"
            clearable
            class="w-[180px]!"
          >
            <el-option
              v-for="option in fileTypeOptions"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            :icon="useRenderIcon(Search)"
            :loading="loading"
            @click="handleSearch"
          >
            搜索
          </el-button>
          <el-button :icon="useRenderIcon(Refresh)" @click="handleReset">
            重置
          </el-button>
        </el-form-item>
      </el-form>

      <!-- 选择提示 -->
      <div class="selection-info mb-4">
        <span class="text-[var(--el-text-color-regular)]">
          {{
            props.multiple
              ? `已选择 ${selectedFiles.length}/${props.maxCount} 个文件`
              : "单选模式"
          }}
        </span>
      </div>

      <!-- 文件网格 -->
      <div v-loading="loading" class="file-grid">
        <div
          v-for="file in fileList"
          :key="file.id"
          class="file-item"
          :class="{
            selected: isFileSelected(file),
            disabled:
              !props.multiple &&
              selectedFiles.length > 0 &&
              !isFileSelected(file)
          }"
          @click="handleSelectFile(file)"
        >
          <div class="file-wrapper">
            <!-- 图片预览 -->
            <el-image
              v-if="isImageFile(file)"
              :src="getFileUrl(file.fileUrl)"
              :alt="file.fileName"
              fit="cover"
              class="file-preview"
            >
              <template #error>
                <div class="file-error">
                  <el-icon><Picture /></el-icon>
                  <span>加载失败</span>
                </div>
              </template>
            </el-image>

            <!-- 文件图标 -->
            <div v-else class="file-icon">
              <el-icon :size="48"
                ><component :is="getFileIcon(file)"
              /></el-icon>
              <div class="file-ext">
                {{ file.fileName.split(".").pop()?.toUpperCase() }}
              </div>
            </div>

            <!-- 选中状态遮罩 -->
            <div v-if="isFileSelected(file)" class="selected-overlay">
              <el-icon class="check-icon"><Check /></el-icon>
            </div>

            <!-- 禁用状态遮罩 -->
            <div v-else-if="isMaxSelected" class="disabled-overlay"></div>
          </div>

          <!-- 文件信息 -->
          <div class="file-info">
            <div class="file-name" :title="file.fileName">
              {{ file.fileName }}
            </div>
            <div class="file-size">{{ formatBytes(file.fileSize) }}</div>
          </div>
        </div>

        <!-- 空状态 -->
        <el-empty
          v-if="!loading && fileList.length === 0"
          description="暂无文件"
        />
      </div>

      <!-- 分页 -->
      <div class="pagination-wrapper mt-4" v-if="pagination.total > 0">
        <el-pagination
          v-model:current-page="pagination.currentPage"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[12, 24, 48, 96]"
          layout="total, sizes, prev, pager, next, jumper"
          background
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>

      <!-- 底部按钮 -->
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleCancel">取消</el-button>
          <el-button
            type="primary"
            :disabled="selectedFiles.length === 0"
            @click="handleConfirm"
          >
            确定选择
            {{ selectedFiles.length > 0 ? `(${selectedFiles.length})` : "" }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.file-image-selector-dialog {
  // 搜索表单样式 - 参考 attachment 页面
  .search-form {
    border-radius: 6px;
    padding-bottom: 12px;

    .el-form-item {
      margin-bottom: 0;
    }
  }

  .selection-info {
    padding: 8px 0;
    border-bottom: 1px solid var(--el-border-color-lighter);
  }

  .file-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 16px;
    min-height: 400px;
    max-height: 500px;
    overflow-y: auto;
    padding: 16px 0;

    .file-item {
      position: relative;
      cursor: pointer;
      border: 2px solid transparent;
      border-radius: 8px;
      transition: all 0.3s ease;
      background: var(--el-fill-color-lighter);

      &:hover {
        border-color: var(--el-color-primary-light-7);
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
      }

      &.selected {
        border-color: var(--el-color-primary);
        box-shadow: 0 0 0 2px var(--el-color-primary-light-8);
      }

      &.disabled {
        cursor: not-allowed;
        opacity: 0.5;
      }

      .file-wrapper {
        position: relative;
        width: 100%;
        height: 140px;
        border-radius: 6px 6px 0 0;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;

        .file-preview {
          width: 100%;
          height: 100%;

          :deep(.el-image__inner) {
            transition: transform 0.3s ease;
          }
        }

        .file-icon {
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          height: 100%;
          color: var(--el-text-color-regular);
          background: var(--el-fill-color-light);

          .el-icon {
            margin-bottom: 8px;
          }

          .file-ext {
            font-size: 12px;
            font-weight: 500;
            color: var(--el-text-color-placeholder);
          }
        }

        .file-error {
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          height: 100%;
          color: var(--el-text-color-placeholder);
          background: var(--el-fill-color-light);

          .el-icon {
            font-size: 32px;
            margin-bottom: 8px;
          }

          span {
            font-size: 12px;
          }
        }

        .selected-overlay {
          position: absolute;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          background: rgba(64, 145, 247, 0.3);
          display: flex;
          align-items: center;
          justify-content: center;

          .check-icon {
            font-size: 32px;
            color: white;
            background: var(--el-color-primary);
            border-radius: 50%;
            padding: 8px;
          }
        }

        .disabled-overlay {
          position: absolute;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          background: rgba(0, 0, 0, 0.3);
        }
      }

      .file-info {
        padding: 12px;

        .file-name {
          font-size: 14px;
          font-weight: 500;
          color: var(--el-text-color-primary);
          margin-bottom: 4px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }

        .file-size {
          font-size: 12px;
          color: var(--el-text-color-regular);
        }
      }
    }
  }

  .pagination-wrapper {
    display: flex;
    justify-content: center;
    padding: 16px 0;
    border-top: 1px solid var(--el-border-color-lighter);
  }

  .dialog-footer {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
  }
}
</style>
