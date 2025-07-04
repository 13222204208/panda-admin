<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { message } from "@/utils/message";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { addDialog } from "@/components/ReDialog";
import { getAttachmentList } from "@/api/attachment";
import type { AttachmentInfo } from "@/api/attachment";

import Search from "~icons/ep/search";
import Picture from "~icons/ep/picture";
import Check from "~icons/ep/check";
import Close from "~icons/ep/close";

interface ImageSelectorProps {
  /** 是否支持多选，默认false */
  multiple?: boolean;
  /** 最大选择数量，仅在multiple为true时生效 */
  maxCount?: number;
  /** 已选择的图片列表 */
  modelValue?: AttachmentInfo[] | AttachmentInfo | null;
}

interface ImageSelectorEmits {
  (
    e: "update:modelValue",
    value: AttachmentInfo[] | AttachmentInfo | null
  ): void;
  (e: "confirm", value: AttachmentInfo[] | AttachmentInfo | null): void;
}

const props = withDefaults(defineProps<ImageSelectorProps>(), {
  multiple: false,
  maxCount: 9
});

const emit = defineEmits<ImageSelectorEmits>();

// 弹窗状态
const dialogVisible = ref(false);
const loading = ref(false);

// 搜索表单
const searchForm = ref({
  fileName: "",
  isImage: true // 只显示图片
});

// 分页配置
const pagination = ref({
  currentPage: 1,
  pageSize: 12,
  total: 0
});

// 图片列表
const imageList = ref<AttachmentInfo[]>([]);

// 选中的图片
const selectedImages = ref<AttachmentInfo[]>([]);

// 计算属性
const isMaxSelected = computed(() => {
  return props.multiple && selectedImages.value.length >= props.maxCount;
});

// 删除原有的 formatFileSize 函数（第66-72行）
// function formatFileSize(bytes: number): string {
//   if (bytes === 0) return "0 B";
//   const k = 1024;
//   const sizes = ["B", "KB", "MB", "GB"];
//   const i = Math.floor(Math.log(bytes) / Math.log(k));
//   return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
// }
// 获取图片列表
const getImageList = async () => {
  loading.value = true;
  try {
    const params = {
      ...searchForm.value,
      currentPage: pagination.value.currentPage,
      pageSize: pagination.value.pageSize
    };

    const { data } = await getAttachmentList(params);
    imageList.value = data.list;
    pagination.value.total = data.total;
  } catch (error) {
    console.error("获取图片列表失败:", error);
    message("获取图片列表失败", { type: "error" });
  } finally {
    loading.value = false;
  }
};

// 搜索
const handleSearch = () => {
  pagination.value.currentPage = 1;
  getImageList();
};

// 重置搜索
const handleReset = () => {
  searchForm.value.fileName = "";
  pagination.value.currentPage = 1;
  getImageList();
};

// 分页处理
const handleSizeChange = (val: number) => {
  pagination.value.pageSize = val;
  getImageList();
};

const handleCurrentChange = (val: number) => {
  pagination.value.currentPage = val;
  getImageList();
};

// 选择图片
const handleSelectImage = (image: AttachmentInfo) => {
  if (props.multiple) {
    const index = selectedImages.value.findIndex(item => item.id === image.id);
    if (index > -1) {
      selectedImages.value.splice(index, 1);
    } else {
      if (selectedImages.value.length < props.maxCount) {
        selectedImages.value.push(image);
      } else {
        message(`最多只能选择${props.maxCount}张图片`, { type: "warning" });
      }
    }
  } else {
    selectedImages.value = [image];
  }
};

// 判断图片是否被选中
const isImageSelected = (image: AttachmentInfo) => {
  return selectedImages.value.some(item => item.id === image.id);
};

// 确认选择
const handleConfirm = () => {
  if (selectedImages.value.length === 0) {
    message("请选择图片", { type: "warning" });
    return;
  }

  const result = props.multiple
    ? selectedImages.value
    : selectedImages.value[0];
  emit("update:modelValue", result);
  emit("confirm", result);
  dialogVisible.value = false;
};

// 取消选择
const handleCancel = () => {
  dialogVisible.value = false;
  selectedImages.value = [];
};

// 打开弹窗
const openDialog = () => {
  // 初始化选中状态
  if (props.modelValue) {
    selectedImages.value = Array.isArray(props.modelValue)
      ? [...props.modelValue]
      : [props.modelValue];
  } else {
    selectedImages.value = [];
  }

  dialogVisible.value = true;
  getImageList();
};

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
  <div class="image-selector">
    <!-- 触发按钮插槽 -->
    <slot name="trigger" :open="openDialog">
      <el-button
        type="primary"
        :icon="useRenderIcon(Picture)"
        @click="openDialog"
      >
        选择图片
      </el-button>
    </slot>

    <!-- 图片选择弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      title="选择图片"
      width="80%"
      :close-on-click-modal="false"
      class="image-selector-dialog"
    >
      <!-- 搜索区域 -->
      <div class="search-area mb-4">
        <el-form :inline="true" :model="searchForm" class="search-form">
          <el-form-item label="图片名称：">
            <el-input
              v-model="searchForm.fileName"
              placeholder="请输入图片名称"
              clearable
              class="w-[200px]!"
              @keyup.enter="handleSearch"
            />
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
            <el-button @click="handleReset"> 重置 </el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 选择提示 -->
      <div class="selection-info mb-4">
        <span class="text-[var(--el-text-color-regular)]">
          {{
            props.multiple
              ? `已选择 ${selectedImages.length}/${props.maxCount} 张图片`
              : "单选模式"
          }}
        </span>
      </div>

      <!-- 图片网格 -->
      <div v-loading="loading" class="image-grid">
        <div
          v-for="image in imageList"
          :key="image.id"
          class="image-item"
          :class="{
            selected: isImageSelected(image),
            disabled:
              !props.multiple &&
              selectedImages.length > 0 &&
              !isImageSelected(image)
          }"
          @click="handleSelectImage(image)"
        >
          <div class="image-wrapper">
            <el-image
              :src="image.url"
              :alt="image.fileName"
              fit="cover"
              class="image"
              :preview-src-list="[image.url]"
              :initial-index="0"
              preview-teleported
            >
              <template #error>
                <div class="image-error">
                  <el-icon><Picture /></el-icon>
                  <span>加载失败</span>
                </div>
              </template>
            </el-image>

            <!-- 选中状态遮罩 -->
            <div v-if="isImageSelected(image)" class="selected-overlay">
              <el-icon class="check-icon"><Check /></el-icon>
            </div>

            <!-- 禁用状态遮罩 -->
            <div v-else-if="isMaxSelected" class="disabled-overlay"></div>
          </div>

          <!-- 图片信息 -->
          <div class="image-info">
            <div class="image-name" :title="image.fileName">
              {{ image.fileName }}
            </div>
            <div class="image-size">{{ formatBytes(image.fileSize) }}</div>
          </div>
        </div>

        <!-- 空状态 -->
        <el-empty
          v-if="!loading && imageList.length === 0"
          description="暂无图片"
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
            :disabled="selectedImages.length === 0"
            @click="handleConfirm"
          >
            确定选择
            {{ selectedImages.length > 0 ? `(${selectedImages.length})` : "" }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.image-selector-dialog {
  .search-area {
    padding: 16px;
    background: var(--el-fill-color-lighter);
    border-radius: 6px;

    .search-form {
      margin: 0;

      .el-form-item {
        margin-bottom: 0;
      }
    }
  }

  .selection-info {
    padding: 8px 0;
    border-bottom: 1px solid var(--el-border-color-lighter);
  }

  .image-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 16px;
    min-height: 400px;
    max-height: 500px;
    overflow-y: auto;
    padding: 16px 0;

    .image-item {
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

      .image-wrapper {
        position: relative;
        width: 100%;
        height: 140px;
        border-radius: 6px 6px 0 0;
        overflow: hidden;

        .image {
          width: 100%;
          height: 100%;

          :deep(.el-image__inner) {
            transition: transform 0.3s ease;
          }
        }

        .image-error {
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

      .image-info {
        padding: 12px;

        .image-name {
          font-size: 14px;
          font-weight: 500;
          color: var(--el-text-color-primary);
          margin-bottom: 4px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }

        .image-size {
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

// 响应式设计
@media (max-width: 768px) {
  .image-selector-dialog {
    .image-grid {
      grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
      gap: 12px;

      .image-item {
        .image-wrapper {
          height: 120px;
        }

        .image-info {
          padding: 8px;

          .image-name {
            font-size: 12px;
          }

          .image-size {
            font-size: 11px;
          }
        }
      }
    }
  }
}
</style>
