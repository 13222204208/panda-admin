<script setup lang="ts">
import { ref } from "vue";
import { ElMessage } from "element-plus";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { generateSql, type GenerateSqlParams } from "@/api/generate";
import { executeSql } from "@/api/generate";

interface Props {
  visible: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(["update:visible"]);

const loading = ref(false);
const generating = ref(false);
const prompt = ref("");
const sqlResult = ref("");

// 生成 SQL
const onGenerate = async () => {
  if (!prompt.value.trim()) {
    ElMessage.warning("请输入要求描述");
    return;
  }

  loading.value = true;
  generating.value = true;
  sqlResult.value = "";

  try {
    // 显示生成中状态
    const dots = [".", "..", "..."];
    let i = 0;
    const interval = setInterval(() => {
      sqlResult.value = `正在生成${dots[i]}`;
      i = (i + 1) % dots.length;
    }, 500);

    // 调用真实的API接口
    const params: GenerateSqlParams = {
      prompt: prompt.value.trim()
    };

    const response = await generateSql(params);

    // 清除加载动画
    clearInterval(interval);

    if (response.code === 0 && response.data) {
      sqlResult.value = response.data.sql;
      ElMessage.success("SQL生成成功");
    } else {
      throw new Error(response.message || "生成失败");
    }
  } catch (error: any) {
    console.error("生成SQL错误:", error);
    sqlResult.value = "";

    // 根据错误类型显示不同的提示
    if (error.message?.includes("timeout") || error.message?.includes("超时")) {
      ElMessage.error("请求超时，请稍后重试");
    } else if (
      error.message?.includes("network") ||
      error.message?.includes("网络")
    ) {
      ElMessage.error("网络连接失败，请检查网络");
    } else {
      ElMessage.error(error.message || "生成失败，请稍后重试");
    }
  } finally {
    loading.value = false;
    generating.value = false;
  }
};

// 复制 SQL
const onCopy = async () => {
  if (!sqlResult.value) {
    ElMessage.warning("暂无可复制的内容");
    return;
  }

  try {
    await navigator.clipboard.writeText(sqlResult.value);
    ElMessage.success("已复制到剪贴板");
  } catch (error) {
    // 降级方案
    const textArea = document.createElement("textarea");
    textArea.value = sqlResult.value;
    document.body.appendChild(textArea);
    textArea.select();
    document.execCommand("copy");
    document.body.removeChild(textArea);
    ElMessage.success("已复制到剪贴板");
  }
};

// 执行 SQL
const onExecute = async () => {
  if (!sqlResult.value) {
    ElMessage.warning("请先生成 SQL");
    return;
  }

  loading.value = true;
  try {
    const response = await executeSql({ sql: sqlResult.value });

    if (response.code === 0) {
      ElMessage.success("执行成功");
      handleClose();
    } else {
      throw new Error(response.message || "执行失败");
    }
  } catch (error: any) {
    console.error("执行SQL错误:", error);
    ElMessage.error(error.message || "执行失败，请稍后重试");
  } finally {
    loading.value = false;
  }
};

// 清空内容
const onClear = () => {
  prompt.value = "";
  sqlResult.value = "";
};

const handleClose = () => {
  onClear();
  emit("update:visible", false);
};

// 示例提示
const examplePrompts = ref([
  "创建一个用户表，包含用户名、邮箱、手机号等字段",
  "创建一个订单表，包含订单号、用户ID、商品信息等",
  "创建一个商品表，包含商品名称、价格、库存等字段"
]);

const useExample = (example: string) => {
  prompt.value = example;
};
</script>

<template>
  <el-dialog
    :model-value="props.visible"
    title="SQL 预览"
    width="800px"
    :close-on-click-modal="false"
    @update:model-value="emit('update:visible', $event)"
    @close="handleClose"
  >
    <el-form>
      <el-form-item label="需求描述">
        <el-input
          v-model="prompt"
          type="textarea"
          :rows="4"
          placeholder="请详细描述您要创建的表结构需求，例如：创建一个用户表，包含用户名、邮箱等字段"
          maxlength="500"
          show-word-limit
          clearable
        />
      </el-form-item>

      <!-- 示例提示 -->
      <el-form-item>
        <div class="example-section">
          <div class="example-label">
            <el-icon :size="14">
              <component :is="useRenderIcon('ep:lightbulb')" />
            </el-icon>
            <span>示例提示：</span>
          </div>
          <div class="example-tags">
            <el-tag
              v-for="(example, index) in examplePrompts"
              :key="index"
              size="small"
              type="info"
              effect="plain"
              class="example-tag"
              @click="useExample(example)"
            >
              {{ example }}
            </el-tag>
          </div>
        </div>
      </el-form-item>

      <el-form-item>
        <div class="flex gap-2">
          <el-button
            type="primary"
            :loading="loading"
            :disabled="!prompt.trim()"
            :icon="useRenderIcon('ep:magic-stick')"
            @click="onGenerate"
          >
            {{ loading ? "AI 生成中..." : "生成 SQL" }}
          </el-button>
          <el-button
            v-if="prompt"
            :icon="useRenderIcon('ep:refresh')"
            @click="onClear"
          >
            清空
          </el-button>
        </div>
      </el-form-item>
    </el-form>

    <el-form-item v-if="sqlResult || generating" label="SQL 语句">
      <div class="sql-result-header" v-if="sqlResult && !generating">
        <span class="text-[var(--el-text-color-regular)]">生成的 SQL 语句</span>
        <el-button
          size="small"
          type="primary"
          :icon="useRenderIcon('ep:document-copy')"
          @click="onCopy"
        >
          复制
        </el-button>
      </div>
      <div
        class="sql-editor-container"
        :class="{ 'is-generating': generating }"
      >
        <el-input
          v-model="sqlResult"
          type="textarea"
          :rows="12"
          class="sql-editor"
          resize="none"
          spellcheck="false"
          :readonly="generating"
          placeholder="生成的 SQL 语句将在这里显示..."
        />
      </div>
    </el-form-item>

    <template #footer>
      <div class="flex items-center justify-between">
        <span class="text-[var(--el-text-color-secondary)]">
          提示：SQL 语句可以手动修改
        </span>
        <div class="flex gap-2">
          <el-button @click="handleClose">取 消</el-button>
          <el-button
            type="primary"
            :loading="loading"
            :disabled="!sqlResult || generating"
            :icon="useRenderIcon('ep:position')"
            @click="onExecute"
          >
            {{ loading ? "执行中..." : "执行 SQL" }}
          </el-button>
        </div>
      </div>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
:deep(.el-form-item__content) {
  flex: 1;
}

.example-section {
  width: 100%;

  .example-label {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: var(--el-text-color-regular);
    margin-bottom: 8px;
  }

  .example-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;

    .example-tag {
      cursor: pointer;
      transition: all 0.2s;

      &:hover {
        transform: translateY(-1px);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      }
    }
  }
}

.sql-result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  padding-bottom: 8px;
  border-bottom: 1px solid var(--el-border-color-lighter);
}

.sql-editor-container {
  width: 100%;
  height: 300px;
  border-radius: 4px;
  overflow: hidden;
  border: 1px solid var(--el-border-color);
  transition: border-color 0.3s;

  &.is-generating {
    border-color: var(--el-color-primary);
  }

  .sql-editor {
    height: 100%;

    :deep(.el-textarea__inner) {
      height: 100%;
      font-family: "Menlo", "Monaco", "Consolas", "Courier New", monospace;
      font-size: 14px;
      line-height: 1.6;
      padding: 12px;
      background-color: var(--el-fill-color-light);
      color: var(--el-text-color-primary);
      border: none;
      border-radius: 0;
      resize: none;
      box-shadow: none;

      &:focus {
        outline: none;
        box-shadow: none;
        background-color: var(--el-color-white);
      }

      &[readonly] {
        background-color: var(--el-fill-color-lighter);
        color: var(--el-color-primary);
      }
    }
  }
}
</style>
