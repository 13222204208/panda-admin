<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "./utils/rule";
import { FormProps } from "./utils/types";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    fileName: "",
    originalName: "",
    fileType: "",
    remark: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-form-item label="文件名" prop="fileName">
      <el-input
        v-model="newFormInline.fileName"
        clearable
        placeholder="请输入文件名"
      />
    </el-form-item>

    <el-form-item label="原始名" prop="originalName">
      <el-input
        v-model="newFormInline.originalName"
        clearable
        placeholder="原始文件名"
        disabled
      />
    </el-form-item>

    <el-form-item label="文件类型" prop="fileType">
      <el-input
        v-model="newFormInline.fileType"
        clearable
        placeholder="文件类型"
        disabled
      />
    </el-form-item>

    <el-form-item label="备注">
      <el-input
        v-model="newFormInline.remark"
        placeholder="请输入备注信息"
        type="textarea"
        :rows="3"
      />
    </el-form-item>
  </el-form>
</template>
