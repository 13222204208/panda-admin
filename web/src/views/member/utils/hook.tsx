import dayjs from "dayjs";
import editForm from "../form/index.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import type { FormItemProps } from "../utils/types";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import {
  getMemberList,
  createMember,
  updateMember,
  deleteMember,
  batchDeleteMembers
} from "@/api/member";
import {
  ElMessageBox
} from "element-plus";
import {
  type Ref,
  h,
  ref,
  toRaw,
  reactive,
  onMounted
} from "vue";

export function useMember(tableRef: Ref) {
  const form = reactive({
    username: "",
    email: "",
  });
  const formRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const selectedNum = ref(0);
  const multipleSelection = ref([]);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      type: "selection",
      width: 55,
      align: "left",
      reserveSelection: true
    },
    {
      label: "会员ID",
      prop: "id",
      minWidth: 90
    },
    {
      label: "用户名",
      prop: "username",
      minWidth: 90
    },
    {
      label: "邮箱",
      prop: "email",
      minWidth: 90
    },
    {
      label: "手机号",
      prop: "mobile",
      minWidth: 90
    },
    {
      label: "创建时间",
      prop: "created_at",
      minWidth: 90,
      formatter: ({ created_at }) =>
        dayjs(created_at).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "更新时间",
      prop: "updated_at",
      minWidth: 90,
      formatter: ({ updated_at }) =>
        dayjs(updated_at).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 180,
      slot: "operation"
    }
  ];

  function handleDelete(row) {
    deleteMember(row.id).then(() => {
      message(`您删除了编号为${row.id}的这条数据`, { type: "success" });
      onSearch();
    });
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
    multipleSelection.value = val;
  }

  function onSelectionCancel() {
    selectedNum.value = 0;
    multipleSelection.value = [];
    tableRef.value.getTableRef().clearSelection();
  }

  function onbatchDel() {
    if (multipleSelection.value.length === 0) {
      message("请选择要删除的数据", { type: "warning" });
      return;
    }
    const ids = getKeyList(multipleSelection.value, "id");
    batchDeleteMembers(ids).then(() => {
      message(`已删除${ids.length}条数据`, { type: "success" });
      onSelectionCancel();
      onSearch();
    });
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getMemberList({
      currentPage: pagination.currentPage,
      pageSize: pagination.pageSize,
      ...form
    });
    dataList.value = data.list;
    pagination.total = data.total;
    loading.value = false;
  }

  const resetForm = (formEl) => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}会员信息表`,
      props: {
        formInline: {
          title,
          // 如果是编辑模式，传递完整的row数据；如果是新增模式，使用默认值
          ...(title === "新增" ? {
            username: "",
            email: "",
            mobile: "",
          } : {
            id: row?.id,
            username: row?.username ?? "",
            email: row?.email ?? "",
            mobile: row?.mobile ?? "",
          })
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        function chores() {
          message(`您${title}了编号为${curData.id}的这条数据`, {
            type: "success"
          });
          done();
          onSearch();
        }
        FormRef.validate(valid => {
          if (valid) {
            if (title === "新增") {
              createMember(curData).then(() => {
                chores();
              });
            } else {
              updateMember(curData).then(() => {
                chores();
              });
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
    onSearch,
    resetForm,
    onbatchDel,
    openDialog,
    handleDelete,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}