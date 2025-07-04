import dayjs from "dayjs";
import editForm from "../form/edit.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import type { FormItemProps } from "../utils/types";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import { type Ref, h, ref, reactive, onMounted } from "vue";
import batchForm from "../form/batch.vue";
import type { BatchFormItemProps } from "../utils/types";
// 导入API接口
import {
  getDictList,
  updateDict,
  deleteDict,
  batchDeleteDict,
  batchCreateDict,
  type GetDictListParams
} from "@/api/dict";

export function useDict(tableRef: Ref) {
  const form = reactive({
    dictType: "",
    dictLabel: "",
    status: "",
    createTime: []
  });
  const formRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const selectedNum = ref(0);
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
      label: "序号",
      type: "index",
      width: 70,
      index: (index: number) => {
        return (pagination.currentPage - 1) * pagination.pageSize + index + 1;
      }
    },
    {
      label: "字典标题",
      prop: "title",
      minWidth: 120
    },
    {
      label: "字典类型",
      prop: "dictType",
      minWidth: 120
    },
    {
      label: "字典标签",
      prop: "dictLabel",
      minWidth: 120
    },
    {
      label: "字典值",
      prop: "dictValue",
      minWidth: 120
    },
    {
      label: "排序",
      prop: "sort",
      width: 80
    },
    {
      label: "状态",
      prop: "status",
      width: 80,
      cellRenderer: ({ row, props }) => (
        <el-tag
          size={props.size}
          type={row.status === 1 ? "success" : "danger"}
          effect="plain"
        >
          {row.status === 1 ? "启用" : "禁用"}
        </el-tag>
      )
    },
    {
      label: "备注",
      prop: "remark",
      minWidth: 150,
      showOverflowTooltip: true
    },
    {
      label: "创建时间",
      prop: "createTime",
      minWidth: 180,
      formatter: ({ createTime }) =>
        dayjs(createTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "更新时间",
      prop: "updateTime",
      minWidth: 180,
      formatter: ({ updateTime }) =>
        dayjs(updateTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 180,
      slot: "operation"
    }
  ];

  async function onSearch() {
    loading.value = true;
    try {
      const params: GetDictListParams = {
        currentPage: pagination.currentPage,
        pageSize: pagination.pageSize
      };

      // 添加搜索条件
      if (form.dictType) {
        params.dictType = form.dictType;
      }
      if (form.dictLabel) {
        params.dictLabel = form.dictLabel;
      }
      if (form.status !== "") {
        params.status = Number(form.status);
      }

      const res = await getDictList(params);
      console.log("字典数据列表", res);
      if (res.code === 0) {
        dataList.value = res.data.list;
        pagination.total = res.data.total;
      } else {
        message(res.message || "获取数据失败", { type: "error" });
      }
    } catch (error) {
      console.error("获取字典列表失败:", error);
      message("获取数据失败", { type: "error" });
    } finally {
      loading.value = false;
    }
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function handleUpdate(row) {
    openDialog("编辑", row);
  }

  async function handleDelete(row) {
    try {
      const data = await deleteDict(row.id);
      if (data.code === 0) {
        message(`您删除了字典标签为${row.dictLabel}的这条数据`, {
          type: "success"
        });
        onSearch();
      } else {
        message(data.message || "删除失败", { type: "error" });
      }
    } catch (error) {
      if (error !== "cancel") {
        console.error("删除字典失败:", error);
        message("删除失败", { type: "error" });
      }
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

  function handleSelectionChange(val) {
    selectedNum.value = val.length;
  }

  function onSelectionCancel() {
    selectedNum.value = 0;
    tableRef.value.getTableRef().clearSelection();
  }

  async function onbatchDel() {
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    try {
      const ids = getKeyList(curSelected, "id");
      const data = await batchDeleteDict({ ids });
      if (data.code === 0) {
        tableRef.value.getTableRef().clearSelection();
        message(`成功删除了${curSelected.length}条数据`, { type: "success" });
        onSearch();
      } else {
        message(data.message || "批量删除失败", { type: "error" });
      }
    } catch (error) {
      if (error !== "cancel") {
        console.error("批量删除失败:", error);
        message("批量删除失败", { type: "error" });
      }
    }
  }

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}字典`,
      props: {
        formInline: {
          id: row?.id ?? undefined,
          dictType: row?.dictType ?? "",
          dictLabel: row?.dictLabel ?? "",
          dictValue: row?.dictValue ?? "",
          sort: row?.sort ?? 1,
          status: row?.status ?? 1,
          remark: row?.remark ?? ""
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef }),
      beforeSure: async (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;

        async function chores() {
          message(`您${title}了字典标签为${curData.dictLabel}的这条数据`, {
            type: "success"
          });
          done();
          onSearch();
        }

        FormRef.validate(async valid => {
          if (valid) {
            try {
              let result;
              if (curData.id) {
                // 更新
                result = await updateDict(curData);
              } else {
                message("无法获取字典数据", { type: "warning" });
                return;
              }
              console.log("操作结果", result);
              if (result.code === 0) {
                chores();
              } else {
                message(result.message || "操作失败", { type: "error" });
              }
            } catch (error) {
              console.error("操作失败:", error);
              message("操作失败", { type: "error" });
            }
          }
        });
      }
    });
  }

  function openBatchDialog() {
    addDialog({
      title: "批量新增字典",
      props: {
        formInline: {
          dictType: "",
          title: "", // 添加这个字段
          dictItems: [
            {
              dictLabel: "",
              dictValue: "",
              sort: 1,
              remark: ""
            }
          ]
        }
      },
      width: "85%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(batchForm, { ref: formRef }),
      beforeSure: async (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as BatchFormItemProps;

        async function chores() {
          message(`成功批量添加了${curData.dictItems.length}个字典项`, {
            type: "success"
          });
          done();
          onSearch();
        }

        FormRef.validate(async valid => {
          if (valid) {
            try {
              // 检查字典值是否重复
              const dictValues = curData.dictItems.map(item => item.dictValue);
              const uniqueValues = new Set(dictValues);
              if (dictValues.length !== uniqueValues.size) {
                message("字典值不能重复，请检查后重试", { type: "warning" });
                return;
              }

              const data = await batchCreateDict({
                dictType: curData.dictType,
                title: curData.title,
                dictItems: curData.dictItems.map(item => ({
                  ...item,
                  status: 1 // 默认启用状态
                }))
              });

              if (data.code === 0) {
                chores();
              } else {
                message(data.message || "批量创建失败", { type: "error" });
              }
            } catch (error) {
              console.error("批量创建失败:", error);
              message("批量创建失败", { type: "error" });
            }
          }
        });
      }
    });
  }

  onMounted(() => {
    onSearch();
  });

  return {
    form,
    loading,
    columns,
    dataList,
    selectedNum,
    pagination,
    deviceDetection,
    onSearch,
    resetForm,
    onbatchDel,
    openDialog,
    openBatchDialog,
    handleUpdate,
    handleDelete,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
