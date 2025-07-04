<script setup lang="ts">
import { ref, reactive } from "vue";
import { message } from "@/utils/message";
import { deviceDetection } from "@pureadmin/utils";
import { resetUserPassword } from "@/api/user";
import { useUserStoreHook } from "@/store/modules/user";
import type { FormInstance, FormRules } from "element-plus";

defineOptions({
  name: "AccountManagement"
});

const showPasswordDialog = ref(false);
const passwordFormRef = ref<FormInstance>();
const loading = ref(false);

const passwordForm = reactive({
  currentPassword: "",
  newPassword: "",
  confirmPassword: ""
});

const passwordRules = reactive<FormRules>({
  currentPassword: [
    { required: true, message: "请输入当前密码", trigger: "blur" }
  ],
  newPassword: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 6, max: 20, message: "密码长度应在6-20位之间", trigger: "blur" }
  ],
  confirmPassword: [
    { required: true, message: "请确认新密码", trigger: "blur" },
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error("两次输入的密码不一致"));
        } else {
          callback();
        }
      },
      trigger: "blur"
    }
  ]
});

const list = ref([
  {
    title: "账户密码",
    // illustrate: "当前密码强度：强",
    button: "修改"
  }
]);

function onClick(item) {
  if (item.title === "账户密码") {
    showPasswordDialog.value = true;
  }
}

// 重置密码表单
const resetPasswordForm = () => {
  passwordForm.currentPassword = "";
  passwordForm.newPassword = "";
  passwordForm.confirmPassword = "";
  passwordFormRef.value?.clearValidate();
};

// 关闭密码对话框
const handleClosePasswordDialog = () => {
  showPasswordDialog.value = false;
  resetPasswordForm();
};

// 提交密码修改
const handleSubmitPassword = async () => {
  if (!passwordFormRef.value) return;

  await passwordFormRef.value.validate(async valid => {
    if (valid) {
      try {
        loading.value = true;
        const userStore = useUserStoreHook();
        const userId = userStore?.id;
        console.log("userInfo用户id", userStore.id);
        if (!userId) {
          message("获取用户信息失败", { type: "error" });
          return;
        }

        const res = await resetUserPassword({
          id: userId,
          oldPassword: passwordForm.currentPassword,
          password: passwordForm.newPassword
        });

        if (res.code === 0) {
          message("密码修改成功", { type: "success" });
          handleClosePasswordDialog();
        } else {
          message(res.message || "密码修改失败", { type: "error" });
        }
      } catch (error) {
        console.error("修改密码失败:", error);
        message("修改密码失败，请稍后重试", { type: "error" });
      } finally {
        loading.value = false;
      }
    }
  });
};
</script>

<template>
  <div
    :class="[
      'min-w-[180px]',
      deviceDetection() ? 'max-w-[100%]' : 'max-w-[70%]'
    ]"
  >
    <h3 class="my-8!">账户管理</h3>
    <div v-for="(item, index) in list" :key="index">
      <div class="flex items-center">
        <div class="flex-1">
          <p>{{ item.title }}</p>
          <!-- <el-text class="mx-1" type="info">{{ item.illustrate }}</el-text> -->
        </div>
        <el-button type="primary" text @click="onClick(item)">
          {{ item.button }}
        </el-button>
      </div>
      <el-divider />
    </div>

    <!-- 修改密码对话框 -->
    <el-dialog
      v-model="showPasswordDialog"
      title="修改密码"
      width="400px"
      :before-close="handleClosePasswordDialog"
    >
      <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="100px"
      >
        <el-form-item label="当前密码" prop="currentPassword">
          <el-input
            v-model="passwordForm.currentPassword"
            type="password"
            placeholder="请输入当前密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="passwordForm.newPassword"
            type="password"
            placeholder="请输入新密码(6-20位)"
            show-password
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            show-password
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleClosePasswordDialog">取消</el-button>
          <el-button
            type="primary"
            :loading="loading"
            @click="handleSubmitPassword"
          >
            {{ loading ? "修改中..." : "确定修改" }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style lang="scss" scoped>
.el-divider--horizontal {
  border-top: 0.1px var(--el-border-color) var(--el-border-style);
}
</style>
