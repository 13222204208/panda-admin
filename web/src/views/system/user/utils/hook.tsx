import "./reset.css";
import dayjs from "dayjs";
import roleForm from "../form/role.vue";
import editForm from "../form/index.vue";
import { zxcvbn } from "@zxcvbn-ts/core";
import { handleTree } from "@/utils/tree";
import { message } from "@/utils/message";
import userAvatar from "@/assets/user.jpg";
import { usePublicHooks } from "../../hooks";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import ReCropperPreview from "@/components/ReCropperPreview";
import type { FormItemProps, RoleFormItemProps } from "../utils/types";
import {
  getKeyList,
  isAllEmpty,
  hideTextAtIndex,
  deviceDetection
} from "@pureadmin/utils";
import { getDepartmentList } from "@/api/department";
import { getAllRoles } from "@/api/role";
import {
  getUserList,
  createUser,
  updateUser,
  deleteUser,
  batchDeleteUsers,
  resetUserPassword,
  getUserRoleIds,
  assignUserRoles, // 分配用户角色
  uploadUserAvatar
} from "@/api/user";
import {
  ElForm,
  ElInput,
  ElFormItem,
  ElProgress,
  ElMessageBox
} from "element-plus";
import {
  type Ref,
  h,
  ref,
  toRaw,
  watch,
  computed,
  reactive,
  onMounted
} from "vue";

export function useUser(tableRef: Ref, treeRef: Ref) {
  const form = reactive({
    // 左侧部门树的id
    departmentId: undefined,
    username: undefined,
    phone: undefined,
    status: undefined
  });
  const formRef = ref();
  const ruleFormRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  // 上传头像信息
  const avatarInfo = ref();
  const switchLoadMap = ref({});
  const { switchStyle } = usePublicHooks();
  const higherDeptOptions = ref();
  const treeData = ref([]);
  const treeLoading = ref(true);
  const selectedNum = ref(0);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "勾选列", // 如果需要表格多选，此处label必须设置
      type: "selection",
      fixed: "left",
      reserveSelection: true // 数据刷新后保留选项
    },
    {
      label: "用户编号",
      prop: "id",
      width: 90
    },
    {
      label: "用户头像",
      prop: "avatar",
      cellRenderer: ({ row }) => (
        <el-image
          fit="cover"
          preview-teleported={true}
          src={row.avatar || userAvatar}
          preview-src-list={Array.of(row.avatar || userAvatar)}
          class="w-[24px] h-[24px] rounded-full align-middle"
        />
      ),
      width: 90
    },
    {
      label: "用户名称",
      prop: "username",
      minWidth: 130
    },
    {
      label: "用户昵称",
      prop: "nickname",
      minWidth: 130
    },
    {
      label: "性别",
      prop: "sex",
      minWidth: 90,
      cellRenderer: ({ row, props }) => (
        <el-tag
          size={props.size}
          type={row.sex === 2 ? "danger" : null}
          effect="plain"
        >
          {row.sex === 2 ? "女" : "男"}
        </el-tag>
      )
    },
    {
      label: "部门",
      prop: "dept.name",
      minWidth: 90
    },
    {
      label: "手机号码",
      prop: "phone",
      minWidth: 90,
      formatter: ({ phone }) => hideTextAtIndex(phone, { start: 3, end: 6 })
    },
    {
      label: "状态",
      prop: "status",
      minWidth: 90,
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
      )
    },
    {
      label: "创建时间",
      minWidth: 90,
      prop: "createTime",
      formatter: ({ createTime }) =>
        dayjs(createTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 180,
      slot: "operation"
    }
  ];
  const buttonClass = computed(() => {
    return [
      "h-[20px]!",
      "reset-margin",
      "text-gray-500!",
      "dark:text-white!",
      "dark:hover:text-primary!"
    ];
  });
  // 重置的新密码
  const pwdForm = reactive({
    newPwd: ""
  });
  const pwdProgress = [
    { color: "#e74242", text: "非常弱" },
    { color: "#EFBD47", text: "弱" },
    { color: "#ffa500", text: "一般" },
    { color: "#1bbf1b", text: "强" },
    { color: "#008000", text: "非常强" }
  ];
  // 当前密码强度（0-4）
  const curScore = ref();
  const roleOptions = ref([]);

  function onChange({ row, index }) {
    ElMessageBox.confirm(
      `确认要<strong>${
        row.status === 0 ? "停用" : "启用"
      }</strong><strong style='color:var(--el-color-primary)'>${
        row.username
      }</strong>用户吗?`,
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
        switchLoadMap.value[index] = Object.assign(
          {},
          switchLoadMap.value[index],
          {
            loading: true
          }
        );
        try {
          // 调用更新用户API
          const result = await updateUser({
            id: row.id,
            status: row.status
          });

          // 检查响应code是否为0
          if (result.code !== 0) {
            throw new Error(result.message || "修改用户状态失败");
          }

          switchLoadMap.value[index] = Object.assign(
            {},
            switchLoadMap.value[index],
            {
              loading: false
            }
          );

          message("已成功修改用户状态", {
            type: "success"
          });

          // 可选：刷新表格数据以确保数据同步
          // onSearch();
        } catch (error) {
          console.error("修改用户状态失败:", error);
          // 错误处理：恢复原状态
          row.status === 0 ? (row.status = 1) : (row.status = 0);

          switchLoadMap.value[index] = Object.assign(
            {},
            switchLoadMap.value[index],
            {
              loading: false
            }
          );

          // 显示具体的错误信息
          const errorMessage = error.message || "修改用户状态失败，请重试";
          message(errorMessage, {
            type: "error"
          });
        }
      })
      .catch(() => {
        // 用户取消操作，恢复原状态
        row.status === 0 ? (row.status = 1) : (row.status = 0);
      });
  }

  function handleUpdate(row) {
    console.log(row);
  }

  async function handleDelete(row) {
    const result = await deleteUser(row.id);

    if (result?.code === 0) {
      message(`成功删除用户：${row.username}`, {
        type: "success"
      });
      // 刷新表格数据
      onSearch();
    } else {
      message(result?.message || "删除用户失败", {
        type: "error"
      });
    }
  }

  function handleSizeChange(val: number) {
    console.log(`${val} items per page`);
  }

  function handleCurrentChange(val: number) {
    console.log(`current page: ${val}`);
  }

  /** 当CheckBox选择项发生变化时会触发该事件 */
  function handleSelectionChange(val) {
    selectedNum.value = val.length;
    // 重置表格高度
    tableRef.value.setAdaptive();
  }

  /** 取消选择 */
  function onSelectionCancel() {
    selectedNum.value = 0;
    // 用于多选表格，清空用户的选择
    tableRef.value.getTableRef().clearSelection();
  }

  /** 批量删除 */
  async function onbatchDel() {
    // 返回当前选中的行
    const curSelected = tableRef.value.getTableRef().getSelectionRows();

    if (curSelected.length === 0) {
      message("请先选择要删除的用户", { type: "warning" });
      return;
    }

    const userIds = getKeyList(curSelected, "id");
    // 调用批量删除API
    const result = await batchDeleteUsers({ ids: userIds });

    // 检查API响应
    if (result.code !== 0) {
      message(result?.message || "删除用户失败", {
        type: "error"
      });
      return;
    }

    // 显示成功消息
    message(`成功删除 ${userIds.length} 个用户`, {
      type: "success"
    });

    // 清除选中状态
    tableRef.value.getTableRef().clearSelection();

    // 刷新列表
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getUserList(toRaw(form));
    console.log("用户列表", data);
    dataList.value = data.list || [];
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
    form.departmentId = undefined;
    treeRef.value.onTreeReset();
    onSearch();
  };

  function onTreeSelect({ id, selected }) {
    form.departmentId = selected ? id : "";
    onSearch();
  }

  function formatHigherDeptOptions(treeList) {
    // 根据返回数据的status字段值判断追加是否禁用disabled字段，返回处理后的树结构，用于上级部门级联选择器的展示（实际开发中也是如此，不可能前端需要的每个字段后端都会返回，这时需要前端自行根据后端返回的某些字段做逻辑处理）
    if (!treeList || !treeList.length) return;
    const newTreeList = [];
    for (let i = 0; i < treeList.length; i++) {
      treeList[i].disabled = treeList[i].status === 0 ? true : false;
      formatHigherDeptOptions(treeList[i].children);
      newTreeList.push(treeList[i]);
    }
    return newTreeList;
  }

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}用户`,
      props: {
        formInline: {
          title,
          higherDeptOptions: formatHigherDeptOptions(higherDeptOptions.value),
          departmentId: row?.dept?.id ?? 0,
          nickname: row?.nickname ?? "",
          username: row?.username ?? "",
          password: row?.password ?? "",
          phone: row?.phone ?? "",
          email: row?.email ?? "",
          sex: row?.sex ?? "",
          status: row?.status ?? 1,
          remark: row?.remark ?? ""
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;

        // 修复：移除递归调用，正确实现关闭弹框和刷新数据的逻辑
        function handleSuccess(username: string) {
          message(`您${title}了用户名称为${username}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        // 参数验证函数
        function validateParams(
          data: FormItemProps,
          isCreate: boolean
        ): string | null {
          if (!data.username?.trim()) {
            return "用户名不能为空";
          }
          if (isCreate && !data.password?.trim()) {
            return "密码不能为空";
          }

          return null;
        }

        // 创建用户参数
        function buildCreateParams(data: FormItemProps) {
          return {
            title: data.title,
            departmentId: data.departmentId,
            nickname: data.nickname,
            username: data.username,
            password: data.password,
            phone: data.phone,
            email: data.email,
            sex: data.sex,
            status: data.status,
            remark: data.remark || ""
          };
        }

        // 更新用户参数
        function buildUpdateParams(data: FormItemProps) {
          const params: any = {
            id: data.id,
            title: data.title,
            departmentId: data.departmentId,
            nickname: data.nickname,
            username: data.username,
            phone: data.phone,
            email: data.email,
            sex: data.sex,
            status: data.status,
            remark: data.remark || ""
          };

          // 只有在密码不为空时才包含密码字段
          if (data.password?.trim()) {
            params.password = data.password;
          }

          return params;
        }

        FormRef.validate(async (valid: boolean) => {
          if (!valid) {
            message("请完善表单信息", { type: "warning" });
            return;
          }

          try {
            console.log("表单数据:", curData);

            const isCreate = title === "新增";

            // 参数验证
            const validationError = validateParams(curData, isCreate);
            if (validationError) {
              message(validationError, { type: "warning" });
              return;
            }

            let result;

            if (isCreate) {
              // 新增用户
              const createParams = buildCreateParams(curData);
              console.log("创建用户参数:", createParams);
              result = await createUser(createParams);
            } else {
              // 更新用户
              curData.id = row?.id;
              if (!curData.id) {
                message("用户ID不能为空", { type: "warning" });
                return;
              }
              const updateParams = buildUpdateParams(curData);
              console.log("更新用户参数:", updateParams);
              result = await updateUser(updateParams);
            }

            // 处理API响应
            if (result?.code === 0) {
              const action = isCreate ? "新增" : "更新";
              message(`成功${action}用户：${curData.username}`, {
                type: "success"
              });
              handleSuccess(curData.username);
            } else {
              const action = isCreate ? "新增" : "更新";
              const errorMsg = result?.message || `${action}用户失败`;
              console.error(`${action}用户失败:`, result);
              message(errorMsg, { type: "error" });
            }
          } catch (error) {
            console.error("用户操作异常:", {
              title,
              curData,
              error: error?.message || error,
              stack: error?.stack
            });

            // 根据错误类型提供更具体的错误信息
            let errorMessage = `${title}用户失败，请重试`;

            if (error?.message?.includes("Network")) {
              errorMessage = "网络连接失败，请检查网络后重试";
            } else if (error?.message?.includes("timeout")) {
              errorMessage = "请求超时，请重试";
            }

            message(errorMessage, { type: "error" });
          }
        });
      }
    });
  }

  const cropRef = ref();
  /** 上传头像 */
  function handleUpload(row) {
    addDialog({
      title: "裁剪、上传头像",
      width: "40%",
      closeOnClickModal: false,
      fullscreen: deviceDetection(),
      contentRenderer: () =>
        h(ReCropperPreview, {
          ref: cropRef,
          imgSrc: row.avatar || userAvatar,
          onCropper: info => (avatarInfo.value = info)
        }),
      beforeSure: async done => {
        try {
          if (!avatarInfo.value?.base64) {
            message("请先裁剪头像", { type: "warning" });
            return;
          }

          // 调用头像上传接口
          const result = await uploadUserAvatar({
            id: row.id,
            avatar: avatarInfo.value.base64
          });

          if (result.code === 0) {
            message("头像上传成功", { type: "success" });
            done(); // 关闭弹框
            onSearch(); // 刷新表格数据
          } else {
            message("头像上传失败", { type: "error" });
          }
        } catch (error) {
          console.error("头像上传失败：", error);
          message("头像上传失败，请重试", { type: "error" });
        }
      },
      closeCallBack: () => cropRef.value.hidePopover()
    });
  }

  watch(
    pwdForm,
    ({ newPwd }) =>
      (curScore.value = isAllEmpty(newPwd) ? -1 : zxcvbn(newPwd).score)
  );

  /** 重置密码 */
  function handleReset(row) {
    addDialog({
      title: `重置 ${row.username} 用户的密码`,
      width: "30%",
      draggable: true,
      closeOnClickModal: false,
      fullscreen: deviceDetection(),
      contentRenderer: () => (
        <>
          <ElForm ref={ruleFormRef} model={pwdForm}>
            <ElFormItem
              prop="newPwd"
              rules={[
                {
                  required: true,
                  message: "请输入新密码",
                  trigger: "blur"
                }
              ]}
            >
              <ElInput
                clearable
                show-password
                type="password"
                v-model={pwdForm.newPwd}
                placeholder="请输入新密码"
              />
            </ElFormItem>
          </ElForm>
          <div class="my-4 flex">
            {pwdProgress.map(({ color, text }, idx) => (
              <div
                class="w-[19vw]"
                style={{ marginLeft: idx !== 0 ? "4px" : 0 }}
              >
                <ElProgress
                  striped
                  striped-flow
                  duration={curScore.value === idx ? 6 : 0}
                  percentage={curScore.value >= idx ? 100 : 0}
                  color={color}
                  stroke-width={10}
                  show-text={false}
                />
                <p
                  class="text-center"
                  style={{ color: curScore.value === idx ? color : "" }}
                >
                  {text}
                </p>
              </div>
            ))}
          </div>
        </>
      ),
      closeCallBack: () => (pwdForm.newPwd = ""),
      beforeSure: done => {
        ruleFormRef.value.validate(async valid => {
          if (valid) {
            try {
              // 调用重置密码接口
              const result = await resetUserPassword({
                id: row.id,
                password: pwdForm.newPwd
              });

              // 检查响应码是否为0（成功）
              if (result.code === 0) {
                // 表单规则校验通过
                message(`已成功重置 ${row.username} 用户的密码`, {
                  type: "success"
                });

                done(); // 关闭弹框
                onSearch(); // 刷新表格数据
              } else {
                // 响应码不为0，显示错误信息
                message(`重置密码失败: ${result.message || "操作失败"}`, {
                  type: "error"
                });
              }
            } catch (error) {
              message(`重置密码失败: ${error.message || "未知错误"}`, {
                type: "error"
              });
            }
          }
        });
      }
    });
  }

  /** 分配角色 */
  async function handleRole(row) {
    // 选中的角色列表
    const ids = (await getUserRoleIds(row.id)).data?.roleIds ?? [];
    addDialog({
      title: `分配 ${row.username} 用户的角色`,
      props: {
        formInline: {
          username: row?.username ?? "",
          nickname: row?.nickname ?? "",
          roleOptions: roleOptions.value ?? [],
          ids
        }
      },
      width: "400px",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(roleForm),
      // 修改 beforeSure 回调函数
      beforeSure: async (done, { options }) => {
        const curData = options.props.formInline as RoleFormItemProps;
        console.log("curIds", curData.ids);

        try {
          // 调用角色分配接口
          const result = await assignUserRoles({
            userId: row.id,
            roleIds: curData.ids.map(item => Number(item))
          });

          if (result.code === 0) {
            message("角色分配成功", { type: "success" });
            done(); // 关闭弹框
          } else {
            message(`角色分配失败: ${result.message || "操作失败"}`, {
              type: "error"
            });
          }
        } catch (error) {
          message(`角色分配失败: ${error.message || "未知错误"}`, {
            type: "error"
          });
        }
      }
    });
  }

  onMounted(async () => {
    treeLoading.value = true;
    onSearch();

    // 归属部门
    const { data } = await getDepartmentList();
    higherDeptOptions.value = handleTree(data.list);
    treeData.value = handleTree(data.list);
    treeLoading.value = false;

    // 角色列表
    roleOptions.value = await getAllRoles();
  });

  return {
    form,
    loading,
    columns,
    dataList,
    treeData,
    treeLoading,
    selectedNum,
    pagination,
    buttonClass,
    deviceDetection,
    onSearch,
    resetForm,
    onbatchDel,
    openDialog,
    onTreeSelect,
    handleUpdate,
    handleDelete,
    handleUpload,
    handleReset,
    handleRole,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
