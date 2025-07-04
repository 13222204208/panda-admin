import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse, PageResponse, PageParams } from "./common/types";

// 部门信息接口
export interface DepartmentInfo {
  id: number;
  parentId?: number;
  name: string;
  principal?: string;
  phone?: string;
  email?: string;
  sort?: number;
  status?: number;
  remark?: string;
  createTime?: string;
}

// 部门列表查询参数
export interface DepartmentListParams extends PageParams {
  parentId?: number;
  name?: string;
  status?: number;
}

// 创建部门参数
export type CreateDepartmentParams = Omit<DepartmentInfo, "id" | "createTime">;

// 更新部门参数
export type UpdateDepartmentParams = Partial<CreateDepartmentParams> & {
  id: number;
};

// 批量删除参数
export interface BatchDeleteParams {
  ids: number[];
}

// 获取部门列表（支持分页）
export const getDepartmentList = (params?: DepartmentListParams) => {
  return http.request<PageResponse<DepartmentInfo[]>>(
    "get",
    baseUrlApi("department"),
    { params }
  );
};

// 获取部门详情
export const getDepartmentDetail = (id: number) => {
  return http.request<BaseResponse<DepartmentInfo>>(
    "get",
    baseUrlApi(`department/${id}`)
  );
};

// 创建部门
export const createDepartment = (data: CreateDepartmentParams) => {
  return http.request<BaseResponse<DepartmentInfo>>(
    "post",
    baseUrlApi("department"),
    { data }
  );
};

// 更新部门
export const updateDepartment = (data: UpdateDepartmentParams) => {
  const { id, ...updateData } = data;
  return http.request<BaseResponse<DepartmentInfo>>(
    "put",
    baseUrlApi(`department/${id}`),
    { data: updateData }
  );
};

// 删除部门
export const deleteDepartment = (id: number) => {
  return http.request<BaseResponse<null>>(
    "delete",
    baseUrlApi(`department/${id}`)
  );
};

// 获取所有部门（不分页，用于下拉选择等场景）
export const getAllDepartments = () => {
  return http.request<BaseResponse<DepartmentInfo[]>>(
    "get",
    baseUrlApi("department"),
    { params: { pageSize: -1 } } // 使用特殊参数表示不分页
  );
};

// 获取部门树形结构
export const getDepartmentTree = () => {
  return http.request<BaseResponse<DepartmentInfo[]>>(
    "get",
    baseUrlApi("department/tree")
  );
};
