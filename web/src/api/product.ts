import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";

type Result = {
  success: boolean;
  data?: any;
};

type ResultTable = {
  success: boolean;
  data?: {
    /** 列表数据 */
    list: Array<any>;
    /** 总条目数 */
    total?: number;
    /** 每页显示条目个数 */
    pageSize?: number;
    /** 当前页数 */
    currentPage?: number;
  };
};

/** 商品信息表数据类型 */
type ProductData = {
  name?: string;
  description?: string;
  price?: string;
  stock?: string;
  main_image?: string;
  carousel_images?: string;
  attachment?: string;
  status?: string;
};

/** 商品信息表更新数据类型 */
type ProductUpdateData = ProductData & {
  id: number;
};

/** 商品信息表查询参数类型 */
type ProductQueryParams = {
  name?: string;
  pageSize?: number;
  currentPage?: number;
};
/** 获取商品信息表列表 */
export const getProductList = (data?: ProductQueryParams) => {
  return http.request<ResultTable>("get", baseUrlApi("product"), { params: data });
};
/** 创建商品信息表 */
export const createProduct = (data: ProductData) => {
  return http.request<Result>("post", baseUrlApi("product"), { data });
};
/** 更新商品信息表 */
export const updateProduct = (data: ProductUpdateData) => {
  return http.request<Result>("put", baseUrlApi(`product/${data.id}`), { data });
};
/** 删除商品信息表 */
export const deleteProduct = (id: number) => {
  return http.request<Result>("delete", baseUrlApi(`product/${id}`));
};
/** 批量删除商品信息表 */
export const batchDeleteProducts = (ids: number[]) => {
  return http.request<Result>("delete", baseUrlApi("product/batch"), { data: { ids } });
};

export type { ProductData, ProductUpdateData, ProductQueryParams };