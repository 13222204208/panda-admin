import dayjs from "dayjs";
import { message } from "@/utils/message";
import type { PaginationProps } from "@pureadmin/table";
import { getKeyList } from "@pureadmin/utils";
import { ElMessageBox } from "element-plus";
import { type Ref, ref, reactive, onMounted } from "vue";
import { getGenerateRecordList, deleteCodeGenRecord } from "@/api/generate";
import { useRouter } from "vue-router";

export function useGenerate(tableRef: Ref) {
  const form = reactive({
    tableName: "",
    tableComment: "",
    createTime: []
  });
  const dataList = ref([]);
  const loading = ref(true);
  const selectedNum = ref(0);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const router = useRouter();
  const columns: TableColumnList = [
    {
      label: "序号",
      type: "index",
      width: 70,
      index: (index: number) => {
        return (pagination.currentPage - 1) * pagination.pageSize + index + 1;
      }
    },
    {
      label: "表名称",
      prop: "tableName",
      minWidth: 120
    },
    {
      label: "表描述",
      prop: "tableComment",
      minWidth: 120
    },
    {
      label: "模块名",
      prop: "moduleName",
      minWidth: 120
    },
    {
      label: "状态",
      prop: "status",
      minWidth: 100,
      cellRenderer: ({ row }) => {
        const statusMap = {
          0: { text: "待生成", type: "info" },
          1: { text: "已生成", type: "success" }
        };
        const status = statusMap[row.status] || statusMap[0];
        return (
          <el-tag type={status.type} size="small">
            {status.text}
          </el-tag>
        );
      }
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
      width: 280,
      slot: "operation"
    }
  ];

  function handleCode(row) {
    // 跳转到代码配置页面，传递表信息
    router.push({
      path: "/develop/code",
      query: {
        tableId: row.id
      }
    });
  }

  async function handleDelete(row) {
    try {
      const res = await deleteCodeGenRecord(row.id);
      if (res.code !== 0) {
        message(res.message, { type: "error" });
        return;
      }
      message("删除成功", { type: "success" });
      onSearch(); // 重新加载数据
    } catch (error) {
      if (error !== "cancel") {
        console.error("删除失败:", error);
        message("删除失败，请稍后重试", { type: "error" });
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

  function onbatchDel() {
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    const ids = getKeyList(curSelected, "id");

    dataList.value = dataList.value.filter(item => !ids.includes(item.id));
    tableRef.value.getTableRef().clearSelection();
    onSearch();
    message(`已删除序号为 ${ids.join(",")} 的数据`, { type: "success" });
  }

  async function onSearch() {
    loading.value = true;

    try {
      const params = {
        currentPage: pagination.currentPage,
        pageSize: pagination.pageSize,
        tableName: form.tableName || undefined,
        tableComment: form.tableComment || undefined,
        createTimeStart: form.createTime?.[0]
          ? dayjs(form.createTime[0]).format("YYYY-MM-DD HH:mm:ss")
          : undefined,
        createTimeEnd: form.createTime?.[1]
          ? dayjs(form.createTime[1]).format("YYYY-MM-DD HH:mm:ss")
          : undefined
      };

      const { data } = await getGenerateRecordList(params);

      dataList.value = data.list || [];
      pagination.total = data.total || 0;
    } catch (error) {
      console.error("获取代码生成记录列表失败:", error);
      message("获取数据失败，请稍后重试", { type: "error" });
      dataList.value = [];
      pagination.total = 0;
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
    console.log("更新配置", row);
  }

  function handleDownload(row) {
    ElMessageBox.confirm(
      `确认要生成并下载表 "${row.tableName}" 的代码吗？`,
      "代码生成确认",
      {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning"
      }
    )
      .then(() => {
        message(`正在生成 ${row.tableName} 的代码...`, { type: "info" });

        setTimeout(() => {
          const blob = new Blob(
            [
              `// 生成的代码文件\n// 表名: ${row.tableName}\n// 描述: ${row.tableComment}\n// 模块名: ${row.moduleName}`
            ],
            {
              type: "text/plain"
            }
          );
          const url = window.URL.createObjectURL(blob);
          const link = document.createElement("a");
          link.href = url;
          link.download = `${row.moduleName}.zip`;
          link.click();
          window.URL.revokeObjectURL(url);

          message(`${row.tableName} 代码生成完成！`, { type: "success" });
        }, 2000);
      })
      .catch(() => {
        message("已取消代码生成", { type: "info" });
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
    onSearch,
    resetForm,
    onbatchDel,
    handleUpdate,
    handleDelete,
    handleDownload,
    handleCode, // 添加这行
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
