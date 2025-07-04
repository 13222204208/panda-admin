// 通用API响应类型定义

/** 基础响应结构 */
export interface BaseResponse<T = any> {
  /** 响应码，0表示成功 */
  code: number;
  /** 响应消息 */
  message: string;
  /** 响应数据 */
  data: T;
}

/** 分页响应数据结构 */
export interface PageData<T = any> {
  /** 数据列表 */
  list: T[];
  /** 总条目数 */
  total: number;
  /** 当前页码 */
  currentPage: number;
  /** 每页条目数（可选） */
  pageSize?: number;
}

/** 分页响应类型 */
export type PageResponse<T = any> = BaseResponse<PageData<T>>;

/** 简单响应类型（无数据或简单数据） */
export type SimpleResponse<T = any> = BaseResponse<T>;

/** 分页查询参数 */
export interface PageParams {
  /** 当前页码 */
  currentPage?: number;
  /** 每页条目数 */
  pageSize?: number;
}
