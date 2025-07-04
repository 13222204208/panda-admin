<script setup lang="ts">
import { useI18n } from "vue-i18n";
import Motion from "./utils/motion";
import { useRouter } from "vue-router";
import { message } from "@/utils/message";
import { loginRules } from "./utils/rule";
import { debounce } from "@pureadmin/utils";
import { useNav } from "@/layout/hooks/useNav";
import type { FormInstance } from "element-plus";
import { $t, transformI18n } from "@/plugins/i18n";
import { useLayout } from "@/layout/hooks/useLayout";
import { useUserStoreHook } from "@/store/modules/user";
import { initRouter, getTopMenu } from "@/router/utils";
import { bg, avatar, illustration } from "./utils/static";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { ref, toRaw, reactive, watch, onMounted, computed } from "vue";
import { useTranslationLang } from "@/layout/hooks/useTranslationLang";
import { useDataThemeChange } from "@/layout/hooks/useDataThemeChange";
import { useEventListener } from "@vueuse/core";

import dayIcon from "@/assets/svg/day.svg?component";
import darkIcon from "@/assets/svg/dark.svg?component";
import globalization from "@/assets/svg/globalization.svg?component";
import Lock from "~icons/ri/lock-fill";
import Check from "~icons/ep/check";
import User from "~icons/ri/user-3-fill";
import Info from "~icons/ri/information-line";

defineOptions({
  name: "Login"
});

// ===== 响应式数据 =====
const router = useRouter();
const loading = ref(false);
const checked = ref(false);
const disabled = ref(false);
const ruleFormRef = ref<FormInstance>();
const captchaImg = ref("");
const loginDay = ref(7);

// 表单数据
const ruleForm = reactive({
  username: "developer",
  password: "123456",
  captchaId: "",
  verifyCode: ""
});

// ===== Hooks =====
const { initStorage } = useLayout();
const { t } = useI18n();
const { dataTheme, overallStyle, dataThemeChange } = useDataThemeChange();
const { title, getDropdownItemStyle, getDropdownItemClass } = useNav();
const { locale, translationCh, translationEn } = useTranslationLang();

// ===== 计算属性 =====
const isFormValid = computed(() => {
  return ruleForm.username && ruleForm.password && ruleForm.verifyCode;
});

// ===== 初始化 =====
initStorage();
dataThemeChange(overallStyle.value);

// ===== 方法 =====
/**
 * 获取验证码
 */
const getCaptcha = async (): Promise<void> => {
  try {
    const res = await useUserStoreHook().getCaptcha();
    if (res.code === 0) {
      captchaImg.value = res.data.captchaImg;
      ruleForm.captchaId = res.data.captchaId;
      // 清空验证码输入
      ruleForm.verifyCode = "";
    } else {
      message(res.message || "获取验证码失败", { type: "error" });
    }
  } catch (error) {
    console.error("获取验证码失败:", error);
    message("获取验证码失败，请重试", { type: "error" });
  }
};

/**
 * 用户登录
 */
const onLogin = async (formEl: FormInstance | undefined): Promise<void> => {
  if (!formEl) return;

  try {
    const valid = await formEl.validate();
    if (!valid) return;

    loading.value = true;

    const res = await useUserStoreHook().loginByUsername({
      username: ruleForm.username,
      password: ruleForm.password,
      captchaId: ruleForm.captchaId,
      verifyCode: ruleForm.verifyCode
    });

    if (res.code === 0) {
      // 初始化路由并跳转
      await initRouter();
      await router.push(getTopMenu(true).path);
      message(t("login.pureLoginSuccess"), { type: "success" });
    } else {
      message(res.message || "登录失败", { type: "error" });
      // 登录失败后刷新验证码
      await getCaptcha();
    }
  } catch (error) {
    console.error("登录失败:", error);
    message("登录失败，请重试", { type: "error" });
    await getCaptcha();
  } finally {
    loading.value = false;
  }
};

// ===== 防抖处理 =====
const debouncedLogin = debounce(
  (formRef: FormInstance | undefined) => onLogin(formRef),
  1000,
  true
);

// ===== 事件监听 =====
useEventListener(document, "keydown", ({ code }) => {
  if (
    ["Enter", "NumpadEnter"].includes(code) &&
    !disabled.value &&
    !loading.value &&
    isFormValid.value
  ) {
    debouncedLogin(ruleFormRef.value);
  }
});

// ===== 监听器 =====
watch(checked, bool => {
  useUserStoreHook().SET_ISREMEMBERED(bool);
});

watch(loginDay, value => {
  useUserStoreHook().SET_LOGINDAY(value);
});

// ===== 生命周期 =====
onMounted(() => {
  getCaptcha();
});
</script>

<template>
  <div class="select-none">
    <img :src="bg" class="wave" />
    <!-- 顶部工具栏 -->
    <div class="flex-c absolute right-5 top-3">
      <!-- 主题切换 -->
      <el-switch
        v-model="dataTheme"
        inline-prompt
        :active-icon="dayIcon"
        :inactive-icon="darkIcon"
        @change="dataThemeChange"
      />

      <!-- 国际化 -->
      <el-dropdown trigger="click">
        <globalization
          class="hover:text-primary hover:!bg-[transparent] w-[20px] h-[20px] ml-1.5 cursor-pointer outline-none duration-300"
        />
        <template #dropdown>
          <el-dropdown-menu class="translation">
            <el-dropdown-item
              :style="getDropdownItemStyle(locale, 'zh')"
              :class="['dark:!text-white', getDropdownItemClass(locale, 'zh')]"
              @click="translationCh"
            >
              <IconifyIconOffline
                v-show="locale === 'zh'"
                class="check-zh"
                :icon="Check"
              />
              简体中文
            </el-dropdown-item>
            <el-dropdown-item
              :style="getDropdownItemStyle(locale, 'en')"
              :class="['dark:!text-white', getDropdownItemClass(locale, 'en')]"
              @click="translationEn"
            >
              <span v-show="locale === 'en'" class="check-en">
                <IconifyIconOffline :icon="Check" />
              </span>
              English
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>

    <!-- 登录容器 -->
    <div class="login-container">
      <!-- 插图 -->
      <div class="img">
        <component :is="toRaw(illustration)" />
      </div>

      <!-- 登录表单 -->
      <div class="login-box">
        <div class="login-form">
          <avatar class="avatar" />

          <Motion>
            <h2 class="outline-none">{{ title }}</h2>
          </Motion>

          <el-form
            ref="ruleFormRef"
            :model="ruleForm"
            :rules="loginRules"
            size="large"
          >
            <!-- 用户名 -->
            <Motion :delay="100">
              <el-form-item
                :rules="[
                  {
                    required: true,
                    message: transformI18n($t('login.pureUsernameReg')),
                    trigger: 'blur'
                  }
                ]"
                prop="username"
              >
                <el-input
                  v-model="ruleForm.username"
                  clearable
                  :placeholder="t('login.pureUsername')"
                  :prefix-icon="useRenderIcon(User)"
                />
              </el-form-item>
            </Motion>

            <!-- 密码 -->
            <Motion :delay="150">
              <el-form-item prop="password">
                <el-input
                  v-model="ruleForm.password"
                  clearable
                  show-password
                  :placeholder="t('login.purePassword')"
                  :prefix-icon="useRenderIcon(Lock)"
                />
              </el-form-item>
            </Motion>

            <!-- 验证码 -->
            <Motion :delay="200">
              <el-form-item
                prop="verifyCode"
                :rules="[
                  {
                    required: true,
                    message: '请输入验证码',
                    trigger: 'blur'
                  }
                ]"
              >
                <el-input
                  v-model="ruleForm.verifyCode"
                  clearable
                  :placeholder="t('login.pureVerifyCode')"
                  :prefix-icon="useRenderIcon('ri:shield-keyhole-line')"
                >
                  <template #append>
                    <img
                      v-if="captchaImg"
                      :src="captchaImg"
                      class="cursor-pointer h-[32px] w-[100px] object-contain"
                      alt="验证码"
                      title="点击刷新验证码"
                      @click="getCaptcha"
                    />
                    <div
                      v-else
                      class="h-[32px] w-[100px] flex items-center justify-center text-gray-400 cursor-pointer"
                      @click="getCaptcha"
                    >
                      点击获取
                    </div>
                  </template>
                </el-input>
              </el-form-item>
            </Motion>

            <!-- 记住密码和登录按钮 -->
            <Motion :delay="250">
              <el-form-item>
                <div class="w-full h-[20px] flex justify-between items-center">
                  <el-checkbox v-model="checked">
                    <span class="flex items-center">
                      <select
                        v-model="loginDay"
                        :style="{
                          width: loginDay < 10 ? '10px' : '16px',
                          outline: 'none',
                          background: 'none',
                          appearance: 'none'
                        }"
                      >
                        <option value="1">1</option>
                        <option value="7">7</option>
                        <option value="30">30</option>
                      </select>
                      {{ t("login.pureRemember") }}
                      <IconifyIconOffline
                        v-tippy="{
                          content: t('login.pureRememberInfo'),
                          placement: 'top'
                        }"
                        :icon="Info"
                        class="ml-1"
                      />
                    </span>
                  </el-checkbox>
                </div>

                <el-button
                  class="w-full mt-4"
                  size="default"
                  type="primary"
                  :loading="loading"
                  :disabled="disabled || !isFormValid"
                  @click="onLogin(ruleFormRef)"
                >
                  {{ t("login.pureLogin") }}
                </el-button>
              </el-form-item>
            </Motion>
          </el-form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url("@/style/login.css");
</style>

<style lang="scss" scoped>
:deep(.el-input-group__append, .el-input-group__prepend) {
  padding: 0;
}

.translation {
  :deep(.el-dropdown-menu__item) {
    padding: 5px 40px;
  }

  .check-zh {
    position: absolute;
    left: 20px;
  }

  .check-en {
    position: absolute;
    left: 20px;
  }
}
</style>
