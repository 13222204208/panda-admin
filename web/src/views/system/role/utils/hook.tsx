import dayjs from "dayjs";
import editForm from "../form.vue";
import { handleTree } from "@/utils/tree";
import { message } from "@/utils/message";
import { ElMessageBox } from "element-plus";
import { usePublicHooks } from "../../hooks";
import { transformI18n } from "@/plugins/i18n";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps } from "../utils/types";
import type { PaginationProps } from "@pureadmin/table";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import { getMenuList } from "@/api/menu";
import { getRoleMenuIds, assignRoleMenus } from "@/api/role";
import { type Ref, reactive, ref, onMounted, h, toRaw, watch } from "vue";
import { getRoleList, createRole, updateRole, deleteRole } from "@/api/role";
export function useRole(treeRef: Ref) {
  const form = reactive({
    name: undefined,
    code: undefined,
    status: undefined
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const treeIds = ref([]);
  const treeData = ref([]);
  const isShow = ref(false);
  const loading = ref(true);
  const isLinkage = ref(false);
  const treeSearchValue = ref();
  const switchLoadMap = ref({});
  const isExpandAll = ref(false);
  const isSelectAll = ref(false);
  const { switchStyle } = usePublicHooks();
  const treeProps = {
    value: "id",
    label: "title",
    children: "children"
  };
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "角色编号",
      prop: "id"
    },
    {
      label: "角色名称",
      prop: "name"
    },
    {
      label: "角色标识",
      prop: "code"
    },
    {
      label: "状态",
      cellRenderer: scope => (
        <el-switch
          size={scope.props.size === "small" ? "small" : "default"}
          loading={switchLoadMap.value[scope.index]?.loading}
          v-model={scope.row.status}
          active-value={1}
          inactive-value={0}
          active-text="已启用"
          inactive-text="已停用"
          inline-prompt
          style={switchStyle.value}
          onChange={() => onChange(scope as any)}
        />
      ),
      minWidth: 90
    },
    {
      label: "备注",
      prop: "remark",
      minWidth: 160
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
      width: 210,
      slot: "operation"
    }
  ];
  // const buttonClass = computed(() => {
  //   return [
  //     "h-[20px]!",
  //     "reset-margin",
  //     "text-gray-500!",
  //     "dark:text-white!",
  //     "dark:hover:text-primary!"
  //   ];
  // });

  function onChange({ row, index }) {
    ElMessageBox.confirm(
      `确认要<strong>${
        row.status === 0 ? "停用" : "启用"
      }</strong><strong style='color:var(--el-color-primary)'>${
        row.name
      }</strong>吗?`,
      "系统提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        dangerouslyUseHTMLString: true,
        draggable: true
      }
    )
      .then(async () => {
        // 设置加载状态
        switchLoadMap.value[index] = Object.assign(
          {},
          switchLoadMap.value[index],
          {
            loading: true
          }
        );

        try {
          // 调用更新接口
          const result = await updateRole({
            id: row.id,
            name: row.name,
            code: row.code,
            status: row.status // 使用当前状态（已经被开关组件改变）
          });

          // 检查更新结果
          if (result.code === 0) {
            message(`已${row.status === 1 ? "启用" : "停用"}${row.name}`, {
              type: "success"
            });
          } else {
            // 更新失败，恢复原状态
            row.status = row.status === 0 ? 1 : 0;
            message(result.message || "状态更新失败", { type: "error" });
          }
        } catch (error) {
          console.error("更新角色状态异常:", error);
          // 更新失败，恢复原状态
          row.status = row.status === 0 ? 1 : 0;
          message("状态更新失败，请稍后重试", { type: "error" });
        } finally {
          // 关闭加载状态
          switchLoadMap.value[index] = Object.assign(
            {},
            switchLoadMap.value[index],
            {
              loading: false
            }
          );
        }
      })
      .catch(() => {
        // 用户取消操作，恢复原状态
        row.status = row.status === 0 ? 1 : 0;
      });
  }

  async function handleDelete(row) {
    try {
      // 调用删除API
      const result = await deleteRole(row.id);

      if (result.code === 0) {
        message(`删除角色"${row.name}"成功`, { type: "success" });
        onSearch(); // 刷新列表
      } else {
        message(result.message || "删除角色失败", { type: "error" });
      }
    } catch (error) {
      console.error("删除角色异常:", error);
      message("删除角色失败，请稍后重试", { type: "error" });
    }
  }

  function handleSizeChange(val: number) {
    console.log(`${val} items per page`);
  }

  function handleCurrentChange(val: number) {
    console.log(`current page: ${val}`);
  }

  function handleSelectionChange(val) {
    console.log("handleSelectionChange", val);
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getRoleList(toRaw(form));
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}角色`,
      props: {
        formInline: {
          name: row?.name ?? "",
          code: row?.code ?? "",
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
        function chores() {
          message(`您${title}了角色名称为${curData.name}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        // 修改表单提交逻辑
        FormRef.validate(async valid => {
          if (valid) {
            console.log("curData", curData);
            try {
              let result;

              if (title === "新增") {
                result = await createRole({
                  name: curData.name,
                  code: curData.code,
                  status: 1,
                  remark: curData.remark || ""
                });
              } else {
                result = await updateRole({
                  id: row.id,
                  name: curData.name,
                  code: curData.code,
                  remark: curData.remark || ""
                });
              }
              // 统一处理返回结果
              if (result.code === 0) {
                message(`${title}角色成功`, { type: "success" });
                chores(); // 成功后执行后续操作
              } else {
                message(result.message || `${title}角色失败`, {
                  type: "error"
                });
              }
            } catch (error) {
              console.error(`${title}角色异常:`, error);
              message(`${title}角色失败，请稍后重试`, { type: "error" });
            }
          }
        });
      }
    });
  }

  /** 菜单权限 */
  async function handleMenu(row?: any) {
    const { id } = row;
    if (id) {
      curRow.value = row;
      isShow.value = true;
      const { data } = await getRoleMenuIds(id);
      treeRef.value.setCheckedKeys(data.menuIds);
    } else {
      curRow.value = null;
      isShow.value = false;
    }
  }

  /** 高亮当前权限选中行 */
  function rowStyle({ row: { id } }) {
    return {
      cursor: "pointer",
      background: id === curRow.value?.id ? "var(--el-fill-color-light)" : ""
    };
  }

  /** 菜单权限-保存 */
  async function handleSave() {
    const { id, name } = curRow.value;
    const menuIds = treeRef.value.getCheckedKeys();

    try {
      const res = await assignRoleMenus({ id, menuIds });
      if (res.code === 0) {
        message(`角色名称为${name}的菜单权限修改成功`, {
          type: "success"
        });
        isShow.value = false;
      } else {
        message(res.message || "分配菜单权限失败", { type: "error" });
      }
    } catch (error) {
      console.error("分配菜单权限失败:", error);
      message("分配菜单权限失败", { type: "error" });
    }
  }

  /** 数据权限 可自行开发 */
  // function handleDatabase() {}

  const onQueryChanged = (query: string) => {
    treeRef.value!.filter(query);
  };

  const filterMethod = (query: string, node) => {
    return transformI18n(node.title)!.includes(query);
  };

  onMounted(async () => {
    onSearch();
    const { data } = await getMenuList();
    treeIds.value = getKeyList(data.list, "id");
    treeData.value = handleTree(data.list);
  });

  watch(isExpandAll, val => {
    val
      ? treeRef.value.setExpandedKeys(treeIds.value)
      : treeRef.value.setExpandedKeys([]);
  });

  watch(isSelectAll, val => {
    val
      ? treeRef.value.setCheckedKeys(treeIds.value)
      : treeRef.value.setCheckedKeys([]);
  });

  return {
    form,
    isShow,
    curRow,
    loading,
    columns,
    rowStyle,
    dataList,
    treeData,
    treeProps,
    isLinkage,
    pagination,
    isExpandAll,
    isSelectAll,
    treeSearchValue,
    // buttonClass,
    onSearch,
    resetForm,
    openDialog,
    handleMenu,
    handleSave,
    handleDelete,
    filterMethod,
    transformI18n,
    onQueryChanged,
    // handleDatabase,
    handleSizeChange,
    handleCurrentChange,
    handleSelectionChange
  };
}
