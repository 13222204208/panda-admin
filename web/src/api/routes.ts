import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse } from "./common/types";

// 路由元信息接口
export interface RouteMeta {
  icon?: string;
  title: string;
  rank?: number;
  roles?: string[];
  keepAlive?: boolean;
  frameSrc?: string;
  hiddenTag?: boolean;
  fixedTag?: boolean;
  showLink: boolean;
  // showParent?: boolean;
  enterTransition?: string;
  leaveTransition?: string;
  activePath?: string;
}

// 路由信息接口
export interface RouteInfo {
  path: string;
  name: string;
  component: string;
  meta: RouteMeta;
  children?: RouteInfo[];
}

// 获取用户路由权限响应
export interface GetUserRoutesResponse {
  routes: RouteInfo[];
}

// 获取用户路由权限接口
export const getUserRoutes = () => {
  return http.request<BaseResponse<GetUserRoutesResponse>>(
    "get",
    baseUrlApi("user/routes")
  );
};
