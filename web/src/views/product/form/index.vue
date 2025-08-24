<script setup lang="ts">
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "../utils/rule";
import { FormProps } from "../utils/types";
import ImageSelector from "@/components/ImageSelector";
import type { AttachmentInfo } from "@/api/attachment";
import { getFileUrl } from "@/utils/env";

// 导入图标组件
import Plus from "~icons/ep/plus";
import ZoomIn from "~icons/ep/zoom-in";
import Delete from "~icons/ep/delete";
import Upload from "~icons/ep/upload";
import Document from "~icons/ep/document";
import Close from "~icons/ep/close";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    title: "新增",
    name: "",
    description: "",
    price: "",
    stock: "",
    main_image: "",
    carousel_images: "",
    attachment: "",
    status: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });

// 图片预览
const previewImage = (image: AttachmentInfo | string) => {
  const url =
    image && image.fileUrl ? getFileUrl(image.fileUrl) : getFileUrl(image);
  // 可以使用 el-image-viewer 或其他预览组件
  window.open(url, "_blank");
};

// 移除图片
const removeImage = (fieldName: string, index: number) => {
  if (Array.isArray(newFormInline.value[fieldName])) {
    newFormInline.value[fieldName].splice(index, 1);
  }
};

// 动态键值对相关方法
const addDynamicItem = (fieldName: string) => {
  if (!newFormInline.value[fieldName]) {
    newFormInline.value[fieldName] = {};
  }
  const key = `key_${Date.now()}`;
  newFormInline.value[fieldName][key] = "";
};

const removeDynamicItem = (fieldName: string, key: string) => {
  if (
    newFormInline.value[fieldName] &&
    newFormInline.value[fieldName][key] !== undefined
  ) {
    delete newFormInline.value[fieldName][key];
  }
};
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-row :gutter="30">
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="商品名称" prop="name">
          <!-- 默认使用普通输入框 -->
          <el-input
            v-model="newFormInline.name"
            clearable
            placeholder="请输入商品名称"
          />
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="商品描述" prop="description">
          <!-- 默认使用普通输入框 -->
          <el-input
            v-model="newFormInline.description"
            clearable
            placeholder="请输入商品描述"
          />
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="商品价格" prop="price">
          <!-- 数字输入 -->
          <el-input-number
            v-model="newFormInline.price"
            placeholder="请输入商品价格"
            class="w-full"
          />
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="库存数量" prop="stock">
          <!-- 数字输入 -->
          <el-input-number
            v-model="newFormInline.stock"
            placeholder="请输入库存数量"
            class="w-full"
          />
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="商品主图URL" prop="main_image">
          <!-- 单图上传 -->
          <ImageSelector
            v-model="newFormInline.main_image"
            :multiple="false"
            file-type="image"
          >
            <template #trigger="{ open }">
              <div class="image-upload-trigger" @click="open">
                <img
                  v-if="newFormInline.main_image"
                  :src="getFileUrl(newFormInline.main_image)"
                  class="avatar"
                />
                <div v-else class="upload-placeholder">
                  <el-icon class="upload-icon"><Plus /></el-icon>
                  <div class="upload-text">点击选择商品主图URL</div>
                </div>
              </div>
            </template>
          </ImageSelector>
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item
          label="商品轮播图URL数组(JSON格式)"
          prop="carousel_images"
        >
          <!-- 多图上传 -->
          <ImageSelector
            v-model="newFormInline.carousel_images"
            :multiple="true"
            :max-count="9"
            file-type="image"
          >
            <template #trigger="{ open }">
              <div class="images-upload-container">
                <!-- 已选择的图片预览 -->
                <div class="selected-images">
                  <div
                    v-for="(image, index) in Array.isArray(
                      newFormInline.carousel_images
                    )
                      ? newFormInline.carousel_images
                      : newFormInline.carousel_images
                        ? JSON.parse(newFormInline.carousel_images)
                        : []"
                    :key="index"
                    class="image-preview"
                  >
                    <img :src="getFileUrl(image)" class="preview-img" />
                    <div class="image-overlay">
                      <el-icon
                        class="preview-icon"
                        @click.stop="previewImage(image)"
                      >
                        <ZoomIn />
                      </el-icon>
                      <el-icon
                        class="remove-icon"
                        @click.stop="removeImage('carousel_images', index)"
                      >
                        <Delete />
                      </el-icon>
                    </div>
                  </div>
                </div>
                <!-- 添加图片按钮 -->
                <div class="add-image-btn" @click="open">
                  <el-icon class="add-icon"><Plus /></el-icon>
                  <div class="add-text">添加商品轮播图URL数组(JSON格式)</div>
                </div>
              </div>
            </template>
          </ImageSelector>
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="附件文件URL" prop="attachment">
          <!-- 单文件上传 -->
          <ImageSelector
            v-model="newFormInline.attachment"
            :multiple="false"
            file-type="file"
          >
            <template #trigger="{ open }">
              <div class="file-upload-container">
                <el-button type="primary" @click="open">
                  <el-icon><Upload /></el-icon>
                  点击选择附件文件URL
                </el-button>
                <div v-if="newFormInline.attachment" class="selected-file">
                  <el-icon><Document /></el-icon>
                  <span class="file-name">
                    {{
                      (newFormInline.attachment || "").split("/").pop() ||
                      "已选择文件"
                    }}
                  </span>
                  <el-icon
                    class="remove-file"
                    @click.stop="newFormInline.attachment = null"
                  >
                    <Close />
                  </el-icon>
                </div>
              </div>
            </template>
          </ImageSelector>
        </el-form-item>
      </re-col>
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="商品状态(1:上架,0:下架)" prop="status">
          <!-- 单选按钮 -->
          <el-radio-group v-model="newFormInline.status">
            <!-- 这里可以根据字典类型动态加载选项 -->
            <el-radio label="1">选项1</el-radio>
            <el-radio label="2">选项2</el-radio>
          </el-radio-group>
        </el-form-item>
      </re-col>
    </el-row>
  </el-form>
</template>

<style scoped>
/* 单图上传样式 */
.image-upload-trigger {
  cursor: pointer;
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.image-upload-trigger:hover {
  border-color: var(--el-color-primary);
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
  object-fit: cover;
}

.upload-placeholder {
  width: 178px;
  height: 178px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: var(--el-fill-color-lighter);
}

.upload-icon {
  font-size: 28px;
  color: #8c939d;
  margin-bottom: 8px;
}

.upload-text {
  font-size: 14px;
  color: var(--el-text-color-regular);
}

/* 多图上传样式 */
.images-upload-container {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.selected-images {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.image-preview {
  position: relative;
  width: 120px;
  height: 120px;
  border-radius: 6px;
  overflow: hidden;
  border: 1px solid var(--el-border-color);
}

.preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  opacity: 0;
  transition: opacity 0.3s;
}

.image-preview:hover .image-overlay {
  opacity: 1;
}

.preview-icon,
.remove-icon {
  color: white;
  font-size: 18px;
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.preview-icon:hover {
  background: rgba(255, 255, 255, 0.2);
}

.remove-icon:hover {
  background: var(--el-color-danger);
}

.add-image-btn {
  width: 120px;
  height: 120px;
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: var(--el-transition-duration-fast);
  background: var(--el-fill-color-lighter);
}

.add-image-btn:hover {
  border-color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
}

.add-icon {
  font-size: 24px;
  color: var(--el-text-color-placeholder);
  margin-bottom: 4px;
}

.add-text {
  font-size: 12px;
  color: var(--el-text-color-regular);
}

/* 文件上传样式 */
.file-upload-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.selected-file {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: var(--el-fill-color-light);
  border-radius: 4px;
  border: 1px solid var(--el-border-color);
}

.file-name {
  flex: 1;
  font-size: 14px;
  color: var(--el-text-color-primary);
}

.remove-file {
  cursor: pointer;
  color: var(--el-text-color-placeholder);
  transition: color 0.3s;
}

.remove-file:hover {
  color: var(--el-color-danger);
}

.files-upload-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.selected-files {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.file-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: var(--el-fill-color-light);
  border-radius: 4px;
  border: 1px solid var(--el-border-color);
}

/* 动态键值对样式 */
.dynamic-input {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.dynamic-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.dynamic-item .el-input {
  flex: 1;
}

.dynamic-item .el-button {
  flex-shrink: 0;
}
</style>
