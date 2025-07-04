import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type {
  BaseResponse,
  PageResponse,
  PageParams,
  PageData
} from "./common/types";

/** 角色信息 */
export interface RoleInfo {
  /** 角色ID */
  id: number;
  /** 角色名称 */
  name: string;
  /** 角色编码 */
  code: string;
  /** 状态 */
  status: number;
  /** 备注 */
  remark: string;
  /** 创建时间 */
  createTime: string;
  /** 更新时间（可选） */
  updateTime?: string;
}

/** 角色列表查询参数 */
export interface RoleListParams extends PageParams {
  /** 角色名称（模糊查询） */
  name?: string;
  /** 角色编码（模糊查询） */
  code?: string;
  /** 状态筛选 */
  status?: number;
}

/** 创建角色参数 */
export type CreateRoleParams = Omit<
  RoleInfo,
  "id" | "createTime" | "updateTime"
>;

/** 更新角色参数 */
export type UpdateRoleParams = Partial<CreateRoleParams> & { id: number };

/** 批量删除参数 */
export interface BatchDeleteParams {
  /** 角色ID数组 */
  ids: number[];
}

/** 获取角色列表 */
export const getRoleList = (params?: RoleListParams) => {
  console.log("请求的参数", params);
  return http.request<PageResponse<RoleInfo>>("get", baseUrlApi("role"), {
    params
  });
};

/** 获取角色详情 */
export const getRoleDetail = (id: number) => {
  return http.request<BaseResponse<RoleInfo>>("get", baseUrlApi(`role/${id}`));
};

/** 创建角色 */
export const createRole = (data: CreateRoleParams) => {
  return http.request<BaseResponse<RoleInfo>>("post", baseUrlApi("role"), {
    data
  });
};

/** 更新角色 */
export const updateRole = (data: UpdateRoleParams) => {
  return http.request<BaseResponse<RoleInfo>>(
    "put",
    baseUrlApi(`role/${data.id}`),
    { data }
  );
};

/** 删除角色 */
export const deleteRole = (id: number) => {
  return http.request<BaseResponse<null>>("delete", baseUrlApi(`role/${id}`));
};

/** 获取所有角色列表（不分页，用于下拉选择等场景） */
export const getAllRoles = async (): Promise<RoleInfo[]> => {
  const response = await http.request<BaseResponse<PageData<RoleInfo>>>(
    "get",
    baseUrlApi("all-role")
  );
  return response.data.list;
};

/** 分配角色菜单权限参数 */
export interface AssignRoleMenusParams {
  /** 角色ID */
  id: number;
  /** 菜单ID列表 */
  menuIds: number[];
}

/** 获取角色菜单ID列表响应 */
export interface RoleMenuIdsResponse {
  /** 菜单ID列表 */
  menuIds: number[];
}

/** 分配角色菜单权限 */
export const assignRoleMenus = (data: AssignRoleMenusParams) => {
  return http.request<BaseResponse<null>>(
    "post",
    baseUrlApi(`role/${data.id}/menus`),
    { data: { menuIds: data.menuIds } }
  );
};

/** 获取角色菜单ID列表 */
export const getRoleMenuIds = (id: number) => {
  return http.request<BaseResponse<RoleMenuIdsResponse>>(
    "get",
    baseUrlApi(`role/${id}/menu-ids`)
  );
};
