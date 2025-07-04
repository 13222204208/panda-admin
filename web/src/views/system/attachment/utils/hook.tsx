import dayjs from "dayjs";
import editForm from "../form.vue";
import { message } from "@/utils/message";
import { ElImage } from "element-plus";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps } from "../utils/types";
import type { PaginationProps } from "@pureadmin/table";
import { deviceDetection } from "@pureadmin/utils";
import { reactive, ref, h } from "vue";
import {
  getAttachmentList,
  updateAttachment,
  deleteAttachment,
  batchDeleteAttachments,
  downloadAttachment,
  type AttachmentInfo
} from "@/api/attachment";

export function useAttachment() {
  const form = reactive({
    fileName: undefined,
    fileExt: undefined,
    isImage: undefined
  });

  const formRef = ref();
  const dataList = ref<AttachmentInfo[]>([]);
  const loading = ref(true);
  const selectedNum = ref(0);
  // 添加选中的附件数据存储
  const selectedAttachments = ref<AttachmentInfo[]>([]);

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });

  const columns: TableColumnList = [
    {
      label: "勾选列",
      type: "selection",
      fixed: "left",
      reserveSelection: true
    },
    {
      label: "附件ID",
      prop: "id",
      width: 90
    },
    {
      label: "预览",
      prop: "thumbnailUrl",
      width: 80,
      cellRenderer: ({ row }) => {
        if (row.isImage && row.fileUrl) {
          return (
            <ElImage
              src={row.fileUrl}
              preview-src-list={[row.fileUrl]}
              fit="cover"
              class="w-[40px] h-[40px] rounded cursor-pointer"
              preview-teleported={true}
            />
          );
        } else {
          return (
            <div class="w-[40px] h-[40px] bg-gray-100 rounded flex items-center justify-center">
              <span class="text-xs text-gray-500">
                {row.fileExt?.toUpperCase()}
              </span>
            </div>
          );
        }
      }
    },
    {
      label: "文件名",
      prop: "fileName",
      minWidth: 200,
      showOverflowTooltip: true
    },
    {
      label: "原始文件名",
      prop: "originalName",
      minWidth: 200,
      showOverflowTooltip: true
    },
    {
      label: "文件后缀名",
      prop: "fileExt",
      width: 120
    },
    {
      label: "文件大小",
      prop: "fileSize",
      width: 100,
      formatter: ({ fileSize }) => formatFileSize(fileSize)
    },
    {
      label: "类别",
      prop: "isImage",
      width: 80,
      cellRenderer: ({ row, props }) => (
        <el-tag
          size={props.size}
          type={row.isImage ? "success" : "info"}
          effect="plain"
        >
          {row.isImage ? "图片" : "文件"}
        </el-tag>
      )
    },
    {
      label: "上传者",
      prop: "uploaderName",
      width: 120
    },
    {
      label: "状态",
      prop: "status",
      width: 100,
      cellRenderer: ({ row, props }) => (
        <el-tag
          size={props.size}
          type={row.status === 1 ? "success" : "danger"}
          effect="plain"
        >
          {row.status === 1 ? "正常" : "已删除"}
        </el-tag>
      )
    },
    {
      label: "创建时间",
      prop: "createTime",
      minWidth: 160,
      formatter: ({ createTime }) =>
        dayjs(createTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 200,
      slot: "operation"
    }
  ];

  // 格式化文件大小
  function formatFileSize(bytes: number): string {
    if (bytes === 0) return "0 B";
    const k = 1024;
    const sizes = ["B", "KB", "MB", "GB", "TB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
  }

  function handleSelectionChange(val: AttachmentInfo[]) {
    selectedNum.value = val.length;
    selectedAttachments.value = val; // 存储选中的附件数据
  }

  function resetForm(formEl) {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    try {
      const { data } = await getAttachmentList({
        currentPage: pagination.currentPage,
        pageSize: pagination.pageSize,
        fileName: form.fileName,
        fileExt: form.fileExt,
        isImage: form.isImage
      });
      dataList.value = data.list || [];
      pagination.total = data.total;
    } catch (error) {
      console.error("获取附件列表失败:", error);
      message("获取附件列表失败", { type: "error" });
    } finally {
      loading.value = false;
    }
  }

  function openDialog(title = "编辑", row?: AttachmentInfo) {
    addDialog({
      title: `${title}附件`,
      props: {
        formInline: {
          id: row?.id,
          fileName: row?.fileName ?? "",
          originalName: row?.originalName ?? "",
          fileType: row?.fileType ?? "",
          remark: row?.remark ?? ""
        }
      },
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        FormRef.validate(async valid => {
          if (valid) {
            try {
              const response = await updateAttachment({
                id: curData.id!,
                fileName: curData.fileName,
                remark: curData.remark
              });
              if (response.code === 0) {
                message(`${title}附件成功`, { type: "success" });
                done();
                onSearch();
              } else {
                message(`${title}附件失败：${response.message}`, {
                  type: "error"
                });
              }
            } catch (error) {
              console.error(`${title}附件异常:`, error);
              message(`${title}附件失败，请稍后重试`, { type: "error" });
            }
          }
        });
      }
    });
  }

  async function handleDelete(row: AttachmentInfo) {
    try {
      const result = await deleteAttachment(row.id);
      if (result.code === 0) {
        message(`删除附件"${row.fileName}"成功`, { type: "success" });
        onSearch();
      } else {
        message(result.message || "删除附件失败", { type: "error" });
      }
    } catch (error) {
      console.error("删除附件异常:", error);
      message("删除附件失败，请稍后重试", { type: "error" });
    }
  }

  async function handleDownload(row: AttachmentInfo) {
    try {
      const blob = await downloadAttachment(row.id);
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = url;
      link.download = row.originalName;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
      message("下载成功", { type: "success" });
    } catch (error) {
      console.error("下载附件异常:", error);
      message("下载失败，请稍后重试", { type: "error" });
    }
  }

  async function onbatchDel() {
    if (selectedAttachments.value.length === 0) {
      message("请先选择要删除的附件", { type: "warning" });
      return;
    }

    try {
      // 提取选中附件的ID
      const ids = selectedAttachments.value.map(item => item.id);

      const result = await batchDeleteAttachments({ ids });

      if (result.code === 0) {
        message(`成功删除 ${selectedAttachments.value.length} 个附件`, {
          type: "success"
        });

        // 清空选中状态
        selectedAttachments.value = [];
        selectedNum.value = 0;

        // 刷新列表
        await onSearch();
      } else {
        message(result.message || "批量删除失败", { type: "error" });
      }
    } catch (error) {
      console.error("批量删除附件异常:", error);
      message("批量删除失败，请稍后重试", { type: "error" });
    }
  }

  function handleSizeChange(val: number) {
    pagination.pageSize = val;
    onSearch();
  }

  function handleCurrentChange(val: number) {
    pagination.currentPage = val;
    onSearch();
  }

  return {
    form,
    loading,
    columns,
    dataList,
    selectedNum,
    pagination,
    onSearch,
    resetForm,
    openDialog,
    handleDelete,
    handleDownload,
    onbatchDel,
    handleSizeChange,
    handleCurrentChange,
    handleSelectionChange
  };
}
