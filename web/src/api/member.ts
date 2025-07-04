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

/** 会员信息表数据类型 */
type MemberData = {
  username?: string;
  email?: string;
  mobile?: string;
};

/** 会员信息表更新数据类型 */
type MemberUpdateData = MemberData & {
  id: number;
};

/** 会员信息表查询参数类型 */
type MemberQueryParams = {
  username?: string;
  email?: string;
  pageSize?: number;
  currentPage?: number;
};
/** 获取会员信息表列表 */
export const getMemberList = (data?: MemberQueryParams) => {
  return http.request<ResultTable>("get", baseUrlApi("member"), { params: data });
};
/** 创建会员信息表 */
export const createMember = (data: MemberData) => {
  return http.request<Result>("post", baseUrlApi("member"), { data });
};
/** 更新会员信息表 */
export const updateMember = (data: MemberUpdateData) => {
  return http.request<Result>("put", baseUrlApi(`member/${data.id}`), { data });
};
/** 删除会员信息表 */
export const deleteMember = (id: number) => {
  return http.request<Result>("delete", baseUrlApi(`member/${id}`));
};
/** 批量删除会员信息表 */
export const batchDeleteMembers = (ids: number[]) => {
  return http.request<Result>("delete", baseUrlApi("member/batch"), { data: { ids } });
};

export type { MemberData, MemberUpdateData, MemberQueryParams };