import dayjs from "dayjs";
import editForm from "../form/index.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import type { FormItemProps } from "../utils/types";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import {
  getProductList,
  createProduct,
  updateProduct,
  deleteProduct,
  batchDeleteProducts
} from "@/api/product";

import {
  type Ref,
  h,
  ref,
  reactive,
  onMounted
} from "vue";

export function useProduct(tableRef: Ref) {
  const form = reactive({
    name: "",
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
      label: "商品ID",
      prop: "id",
      minWidth: 90
    },
    {
      label: "商品名称",
      prop: "name",
      minWidth: 90
    },
    {
      label: "商品描述",
      prop: "description",
      minWidth: 90
    },
    {
      label: "商品价格",
      prop: "price",
      minWidth: 90
    },
    {
      label: "库存数量",
      prop: "stock",
      minWidth: 90
    },
    {
      label: "商品主图URL",
      prop: "main_image",
      minWidth: 90
    },
    {
      label: "商品轮播图URL数组(JSON格式)",
      prop: "carousel_images",
      minWidth: 90
    },
    {
      label: "附件文件URL",
      prop: "attachment",
      minWidth: 90
    },
    {
      label: "商品状态(1:上架,0:下架)",
      prop: "status",
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
    deleteProduct(row.id).then(() => {
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
    batchDeleteProducts(ids).then(() => {
      message(`已删除${ids.length}条数据`, { type: "success" });
      onSelectionCancel();
      onSearch();
    });
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getProductList({
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
      title: `${title}商品信息表`,
      props: {
        formInline: {
          title,
          // 如果是编辑模式，传递完整的row数据；如果是新增模式，使用默认值
          ...(title === "新增" ? {
            name: "",
            description: "",
            price: "",
            stock: "",
            main_image: "",
            carousel_images: "",
            attachment: "",
            status: "",
          } : {
            id: row?.id,
            name: row?.name ?? "",
            description: row?.description ?? "",
            price: row?.price ?? "",
            stock: row?.stock ?? "",
            main_image: row?.main_image ?? "",
            carousel_images: row?.carousel_images ?? "",
            attachment: row?.attachment ?? "",
            status: row?.status ?? "",
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
              createProduct(curData).then((res) => {
                if (res.code === 0) {
                  chores();
                } else {
                  message(res.message || "操作失败", { type: "error" });
                }
              }).catch((error) => {
                message(error.message || "网络错误", { type: "error" });
              });
            } else {
              updateProduct(curData).then((res) => {
                if (res.code === 0) {
                  chores();
                } else {
                  message(res.message || "操作失败", { type: "error" });
                }
              }).catch((error) => {
                message(error.message || "网络错误", { type: "error" });
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