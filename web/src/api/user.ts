import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse, PageResponse, PageParams } from "./common/types";

export type UserResult = {
  success: boolean;
  data: {
    /** 头像 */
    avatar: string;
    /** 用户名 */
    username: string;
    /** 昵称 */
    nickname: string;
    /** 当前登录用户的角色 */
    roles: Array<string>;
    /** 按钮级别权限 */
    permissions: Array<string>;
    /** `token` */
    accessToken: string;
    /** 用于调用刷新`accessToken`的接口时所需的`token` */
    refreshToken: string;
    /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
    expires: Date;
  };
};

export type RefreshTokenResult = {
  success: boolean;
  data: {
    /** `token` */
    accessToken: string;
    /** 用于调用刷新`accessToken`的接口时所需的`token` */
    refreshToken: string;
    /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
    expires: Date;
  };
};

export type UserInfo = {
  /** 头像 */
  avatar: string;
  /** 用户名 */
  username: string;
  /** 昵称 */
  nickname: string;
  /** 邮箱 */
  email: string;
  /** 联系电话 */
  phone: string;
  /** 简介 */
  description: string;
};

export type UserInfoResult = {
  success: boolean;
  data: UserInfo;
};

// 用户管理相关接口类型定义

// 用户信息接口
export interface UserManageInfo {
  id: number;
  title: string;
  departmentId: number;
  dept?: {
    id: number;
    name: string;
  };
  nickname: string;
  username: string;
  phone: string | number;
  email: string;
  sex: string | number;
  status: number;
  remark: string;
  createTime?: string;
}

// 用户列表查询参数
export interface UserListParams extends PageParams {
  title?: string;
  departmentId?: number;
  nickname?: string;
  username?: string;
  phone?: string;
  email?: string;
  sex?: string | number;
  status?: number;
}

// 创建用户参数
export type CreateUserParams = {
  title: string;
  departmentId: number;
  nickname: string;
  username: string;
  password: string;
  phone: string | number;
  email: string;
  sex: string | number;
  status: number;
  remark: string;
};

// 更新用户参数
export type UpdateUserParams = Partial<Omit<CreateUserParams, "password">> & {
  id: number;
  password?: string; // 密码可选，不传则不更新
};

// 重置密码参数
export interface ResetPasswordParams {
  id: number;
  password: string;
  oldPassword?: string;
}

// 批量删除参数
export interface BatchDeleteParams {
  ids: number[];
}

// 用户详情响应
export interface UserDetailResponse extends UserManageInfo {
  roles: Array<{
    id: number;
    name: string;
    code: string;
  }>;
}

/** 刷新`token` */
export const refreshTokenApi = (data?: object) => {
  return http.request<RefreshTokenResult>("post", "/refresh-token", { data });
};

// ==================== 用户管理 CRUD 接口 ====================

// 获取用户列表（支持分页）
export const getUserList = (params?: UserListParams) => {
  return http.request<PageResponse<UserManageInfo[]>>(
    "get",
    baseUrlApi("user"),
    { params }
  );
};

// 获取用户详情
export const getUserDetail = () => {
  return http.request<BaseResponse<UserDetailResponse>>(
    "get",
    baseUrlApi(`user/detail`)
  );
};

// 创建用户
export const createUser = (data: CreateUserParams) => {
  return http.request<BaseResponse<{ id: number }>>(
    "post",
    baseUrlApi("user"),
    { data }
  );
};

// 更新用户
export const updateUser = (data: UpdateUserParams) => {
  const { id, ...updateData } = data;
  return http.request<BaseResponse<null>>("put", baseUrlApi(`user/${id}`), {
    data: updateData
  });
};

// 删除用户
export const deleteUser = (id: number) => {
  return http.request<BaseResponse<null>>("delete", baseUrlApi(`user/${id}`));
};

// 重置用户密码
export const resetUserPassword = (data: ResetPasswordParams) => {
  const { id, password, oldPassword } = data;

  // 参数验证
  if (!id || id <= 0) {
    throw new Error("用户ID不能为空且必须大于0");
  }

  if (!password || password.trim().length === 0) {
    throw new Error("新密码不能为空");
  }

  // 密码强度验证（可选）
  if (password.length < 6) {
    throw new Error("密码长度不能少于6位");
  }

  const requestData: { password: string; oldPassword?: string } = { password };

  // 如果提供了旧密码，则包含在请求中
  if (oldPassword) {
    requestData.oldPassword = oldPassword;
  }

  return http.request<BaseResponse<null>>(
    "put",
    baseUrlApi(`user/${id}/reset-password`),
    { data: requestData }
  );
};

// 批量删除用户
export const batchDeleteUsers = (data: BatchDeleteParams) => {
  return http.request<BaseResponse<null>>("delete", baseUrlApi("user/batch"), {
    data
  });
};
// 获取用户所有角色ID
export interface GetUserRoleIdsResponse {
  roleIds: number[];
}

// 获取用户所有角色ID
export const getUserRoleIds = (userId: number) => {
  return http.request<BaseResponse<GetUserRoleIdsResponse>>(
    "post",
    baseUrlApi("user-role-ids"),
    {
      data: { userId }
    }
  );
};

// 用户角色分配参数
export interface AssignUserRolesParams {
  userId: number;
  roleIds: number[];
}

// 分配用户角色
export const assignUserRoles = (data: AssignUserRolesParams) => {
  return http.request<BaseResponse<null>>(
    "post",
    baseUrlApi("user/assign-roles"),
    { data }
  );
};

// 用户头像上传参数
export interface UploadAvatarParams {
  id: number;
  avatar: string; // base64格式的图片数据
}

// 用户头像上传响应
export interface UploadAvatarResponse {
  avatarUrl: string;
}

// 上传用户头像
export const uploadUserAvatar = (data: UploadAvatarParams) => {
  const { id, avatar } = data;
  return http.request<BaseResponse<UploadAvatarResponse>>(
    "post",
    baseUrlApi(`user/${id}/avatar`),
    { data: { avatar } }
  );
};
