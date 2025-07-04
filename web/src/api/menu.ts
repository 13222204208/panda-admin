import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse } from "./common/types";

/** 菜单信息 */
export interface MenuInfo {
  /** 主键ID */
  id: number;
  /** 菜单类型（0菜单、1 iframe、2外链、3按钮） */
  menuType: number;
  /** 父级菜单ID */
  parentId: number;
  /** 菜单名称 */
  title: string;
  /** 路由名称（必须唯一） */
  name: string;
  /** 路由路径 */
  path: string;
  /** 组件路径 */
  component: string;
  /** 菜单排序（home 的 rank 应为 0） */
  rank: number;
  /** 重定向地址 */
  redirect: string;
  /** 菜单图标 */
  icon: string;
  /** 右侧额外图标 */
  extraIcon: string;
  /** 进场动画 */
  enterTransition: string;
  /** 离场动画 */
  leaveTransition: string;
  /** 激活菜单的 path */
  activePath: string;
  /** 权限标识（按钮级别权限） */
  auths: string;
  /** iframe 链接地址 */
  frameSrc: string;
  /** iframe 页面是否首次加载显示动画 */
  frameLoading: number;
  /** 是否缓存该页面 */
  keepAlive: number;
  /** 是否禁止添加到标签页 */
  hiddenTag: number;
  /** 是否固定在标签页中 */
  fixedTag: number;
  /** 是否在菜单中显示该项 */
  showLink: number;
  /** 是否显示父级菜单 */
  showParent: number;
  /** 创建时间 */
  createdAt: string;
}

/** 菜单树节点 */
export interface MenuTreeNode {
  /** 菜单ID */
  id: number;
  /** 菜单标题 */
  title: string;
  /** 子菜单 */
  children?: MenuTreeNode[];
}

/** 菜单列表查询参数 */
export interface MenuListParams {
  /** 菜单名称 */
  title?: string;
  /** 菜单类型（0菜单、1 iframe、2外链、3按钮） */
  menuType?: number;
  /** 父级菜单ID */
  parentId?: number;
}

/** 菜单公共字段（用于创建和更新） */
export interface MenuCommon {
  /** 菜单类型（0菜单、1 iframe、2外链、3按钮） */
  menuType?: number;
  /** 父级菜单ID */
  parentId?: number;
  /** 菜单名称 */
  title?: string;
  /** 路由名称（必须唯一） */
  name?: string;
  /** 路由路径 */
  path?: string;
  /** 组件路径 */
  component?: string;
  /** 菜单排序（home 的 rank 应为 0） */
  rank?: number;
  /** 重定向地址 */
  redirect?: string;
  /** 菜单图标 */
  icon?: string;
  /** 右侧额外图标 */
  extraIcon?: string;
  /** 进场动画 */
  enterTransition?: string;
  /** 离场动画 */
  leaveTransition?: string;
  /** 激活菜单的 path */
  activePath?: string;
  /** 权限标识（按钮级别权限） */
  auths?: string;
  /** iframe 链接地址 */
  frameSrc?: string;
  /** iframe 页面是否首次加载显示动画 */
  frameLoading?: number | boolean;
  /** 是否缓存该页面 */
  keepAlive?: number | boolean;
  /** 是否禁止添加到标签页 */
  hiddenTag?: number | boolean;
  /** 是否固定在标签页中 */
  fixedTag?: number | boolean;
  /** 是否在菜单中显示该项 */
  showLink?: number | boolean;
  /** 是否显示父级菜单 */
  showParent?: number | boolean;
}

/** 创建菜单参数 */
export type CreateMenuParams = MenuCommon;

/** 更新菜单参数 */
export type UpdateMenuParams = MenuCommon;

// ==================== API 函数 ====================

/** 获取菜单列表（不分页，获取全部菜单） */
export const getMenuList = (params?: MenuListParams) => {
  return http.request<BaseResponse<{ list: MenuInfo[] }>>(
    "get",
    baseUrlApi("menu"),
    { params }
  );
};

/** 获取菜单树（只包含id和title字段） */
export const getMenuTree = () => {
  return http.request<BaseResponse<{ tree: MenuTreeNode[] }>>(
    "get",
    baseUrlApi("menu/tree")
  );
};

/** 创建菜单 */
export const createMenu = (data: CreateMenuParams) => {
  return http.request<BaseResponse<{ id: number }>>(
    "post",
    baseUrlApi("menu"),
    { data }
  );
};

/** 更新菜单 */
export const updateMenu = (data: UpdateMenuParams, id: number) => {
  return http.request<BaseResponse<null>>("put", baseUrlApi(`menu/${id}`), {
    data
  });
};

/** 删除菜单 */
export const deleteMenu = (id: number) => {
  return http.request<BaseResponse<null>>("delete", baseUrlApi(`menu/${id}`));
};
