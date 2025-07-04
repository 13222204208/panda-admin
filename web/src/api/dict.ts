import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type {
  BaseResponse,
  PageResponse,
  PageParams
} from "@/api/common/types";

// 字典项接口
export interface DictItem {
  id?: number;
  title: string;
  dictType: string;
  dictLabel: string;
  dictValue: string;
  sort: number;
  status: number;
  remark?: string;
  createTime?: string;
  updateTime?: string;
}

// 获取字典列表参数
export interface GetDictListParams extends PageParams {
  dictType?: string;
  dictLabel?: string;
  status?: number;
  createTime?: string[];
}

// 创建字典参数
export interface CreateDictParams {
  title: string;
  dictType: string;
  dictLabel: string;
  dictValue: string;
  sort?: number;
  status?: number;
  remark?: string;
}

// 更新字典参数
export interface UpdateDictParams extends CreateDictParams {
  id: number;
}

// 批量删除参数
export interface BatchDeleteParams {
  ids: number[];
}

// 字典列表响应
export interface DictListResponse extends BaseResponse {
  data: PageResponse<DictItem>;
}

/** 获取字典列表 */
export const getDictList = (params?: GetDictListParams) => {
  return http.request<DictListResponse>("get", baseUrlApi("dict"), {
    params
  });
};

/** 更新字典 */
export const updateDict = (data: UpdateDictParams) => {
  return http.request<BaseResponse>("put", baseUrlApi(`dict/${data.id}`), {
    data
  });
};

/** 删除字典 */
export const deleteDict = (id: number) => {
  return http.request<BaseResponse>("delete", baseUrlApi(`dict/${id}`));
};

/** 批量删除字典 */
export const batchDeleteDict = (data: BatchDeleteParams) => {
  return http.request<BaseResponse>("post", baseUrlApi("dict/batch-delete"), {
    data
  });
};

/** 根据字典类型获取字典选项 */
export const getDictOptions = (dictType: string) => {
  return http.request<BaseResponse>(
    "get",
    baseUrlApi(`dict/options/${dictType}`)
  );
};

// 批量创建字典参数
export interface BatchCreateDictParams {
  dictType: string;
  title: string;
  dictItems: {
    dictLabel: string;
    dictValue: string;
    sort?: number;
    status?: number;
    remark?: string;
  }[];
}

/** 批量创建字典 */
export const batchCreateDict = (data: BatchCreateDictParams) => {
  return http.request<BaseResponse>("post", baseUrlApi("dict/batch"), { data });
};

// 字典类型信息
export interface DictTypeInfo {
  title: string;
  dictType: string;
}

// 获取不重复的字典类型和标题响应
export interface GetDistinctTypesResponse extends BaseResponse {
  data: {
    types: DictTypeInfo[];
  };
}

/** 获取不重复的字典类型和标题 */
export const getDistinctDictTypes = () => {
  return http.request<GetDistinctTypesResponse>(
    "get",
    baseUrlApi("dict/types")
  );
};
