<script setup lang="ts">
import { ref, reactive } from "vue";
import type { FormRules } from "element-plus";
import type { FormProps } from "../utils/types";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    id: "",
    dictType: "",
    dictLabel: "",
    dictValue: "",
    sort: 0,
    status: 1,
    remark: ""
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
  dictLabel: [{ required: true, message: "字典标签为必填项", trigger: "blur" }],
  dictValue: [{ required: true, message: "字典值为必填项", trigger: "blur" }],
  sort: [
    {
      type: "number",
      message: "排序值必须为数字",
      trigger: "blur",
      transform: (value: string) => Number(value)
    }
  ],
  status: [{ required: true, message: "状态为必填项", trigger: "change" }]
});

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
    label-width="100px"
  >
    <el-row :gutter="30">
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-form-item label="字典类型" prop="dictType">
          <el-input
            v-model="newFormInline.dictType"
            clearable
            placeholder="请输入字典类型（如 sex、status）"
          />
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-form-item label="字典标签" prop="dictLabel">
          <el-input
            v-model="newFormInline.dictLabel"
            clearable
            placeholder="请输入字典标签（如 男、启用）"
          />
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-form-item label="字典值" prop="dictValue">
          <el-input
            v-model="newFormInline.dictValue"
            clearable
            placeholder="请输入字典值（如 1、enabled）"
          />
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-form-item label="排序值" prop="sort">
          <el-input-number
            v-model="newFormInline.sort"
            :min="0"
            :max="9999"
            controls-position="right"
            placeholder="请输入排序值"
            class="w-full"
          />
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="newFormInline.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-col>
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <el-form-item label="备注说明" prop="remark">
          <el-input
            v-model="newFormInline.remark"
            type="textarea"
            :rows="3"
            placeholder="请输入备注说明"
            maxlength="255"
            show-word-limit
          />
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
