import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse, PageResponse, PageParams } from "./common/types";

// ==================== 数据表相关类型定义 ====================

/** 数据表信息 */
export interface TableInfo {
  /** 表名称 */
  tableName: string;
  /** 表描述/注释 */
  tableComment: string;
  /** 创建时间 */
  createTime: string;
  /** 更新时间 */
  updateTime: string;
}

/** 获取数据表列表参数 */
export interface GetTablesParams extends PageParams {
  /** 表名称（模糊搜索） */
  tableName?: string;
  /** 表描述（模糊搜索） */
  tableComment?: string;
  /** 创建时间范围开始 */
  createTimeStart?: string;
  /** 创建时间范围结束 */
  createTimeEnd?: string;
}

/** 表字段信息 */
export interface TableColumn {
  /** 字段名 */
  columnName: string;
  /** 字段类型 */
  columnType: string;
  /** 字段注释 */
  columnComment: string;
  /** 是否主键 */
  isPrimaryKey: boolean;
  /** 是否可为空 */
  isNullable: boolean;
  /** 默认值 */
  defaultValue?: string;
  /** 字段长度 */
  columnLength?: number;
}

/** 表详细信息 */
export interface TableDetail extends TableInfo {
  /** 表字段列表 */
  columns: TableColumn[];
  /** 表索引信息 */
  indexes?: any[];
}

/** SQL生成参数 */
export interface GenerateSqlParams {
  /** 需求描述 */
  prompt: string;
  /** 目标表名（可选） */
  tableName?: string;
}

/** SQL生成响应 */
export interface GenerateSqlResponse {
  /** 生成的SQL语句 */
  sql: string;
  /** 生成说明 */
  description?: string;
}

/** SQL执行参数 */
export interface ExecuteSqlParams {
  /** SQL语句 */
  sql: string;
}

/** SQL执行响应 */
export interface ExecuteSqlResponse {
  /** 执行结果 */
  success: boolean;
  /** 影响行数 */
  affectedRows?: number;
  /** 查询结果（如果是查询语句） */
  data?: any[];
  /** 错误信息 */
  error?: string;
}

// ==================== API 接口函数 ====================

/**
 * 获取数据库表列表
 * @param params 查询参数
 * @returns 表列表数据
 */
export const getTableList = (params?: GetTablesParams) => {
  return http.request<PageResponse<TableInfo[]>>("get", baseUrlApi("table"), {
    params
  });
};

/**
 * AI生成SQL语句
 * @param params 生成参数
 * @returns 生成的SQL
 */
export const generateSql = (params: GenerateSqlParams) => {
  return http.request<BaseResponse<GenerateSqlResponse>>(
    "post",
    baseUrlApi("sql/generate"),
    { data: params }
  );
};

/**
 * 执行SQL语句
 * @param params 执行参数
 * @returns 执行结果
 */
export const executeSql = (params: ExecuteSqlParams) => {
  return http.request<BaseResponse<ExecuteSqlResponse>>(
    "post",
    baseUrlApi("sql/execute"),
    { data: params }
  );
};

/**
 * 生成代码
 * @param tableName 表名
 * @param options 生成选项
 * @returns 生成结果
 */
export const generateCode = (tableName: string, options?: any) => {
  return http.request<BaseResponse<{ files: any[] }>>(
    "post",
    baseUrlApi(`generate/code/${tableName}`),
    { data: options }
  );
};

/** 导入表信息（简化版，只包含表名和注释） */
export interface ImportTableInfo {
  /** 表名称 */
  tableName: string;
  /** 表描述/注释 */
  tableComment: string;
}

/** 导入表结构请求参数 */
export interface ImportTablesParams {
  /** 要导入的表信息列表 */
  tables: ImportTableInfo[];
}

/** 导入表结构响应参数 */
export interface ImportTablesResponse {
  /** 成功导入的表数量 */
  count: number;
  /** 导入的表名列表 */
  tables?: string[];
  /** 导入失败的错误信息 */
  errors?: string[];
  /** 导入是否成功 */
  success: boolean;
}

/**
 * 导入表结构
 * @param params 导入参数
 * @returns 导入结果
 */
export const importTables = (params: ImportTablesParams) => {
  return http.request<BaseResponse<ImportTablesResponse>>(
    "post",
    baseUrlApi("table/import"),
    { data: params }
  );
};

/**
 * 获取代码生成记录列表
 * @param params 查询参数
 * @returns 代码生成记录列表数据
 */
export const getGenerateRecordList = (params?: PageParams) => {
  return http.request<PageResponse<any[]>>(
    "get",
    baseUrlApi("generate/record"),
    {
      params
    }
  );
};

/** 字段信息 */
export interface ColumnInfo {
  /** 字段名 */
  columnName: string;
  /** 数据类型 */
  dataType: string;
  /** 完整字段类型 */
  columnType: string;
  /** 是否可为空 */
  isNullable: string;
  /** 默认值 */
  columnDefault: string;
  /** 字段注释 */
  columnComment: string;
  /** 键类型 */
  columnKey: string;
  /** 额外信息 */
  extra: string;
}

/** 表及其字段信息 */
export interface TableWithColumnsInfo {
  /** 表名 */
  tableName: string;
  /** 表描述 */
  tableComment: string;
  /** 字段列表 */
  columns: ColumnInfo[];
}

/** 获取表和字段信息响应 */
export interface GetTablesWithColumnsResponse {
  /** 表和字段信息列表 */
  list: TableWithColumnsInfo[];
}

/**
 * 获取所有表和字段信息（过滤时间字段）
 * @returns 表和字段信息列表
 */
export const getTablesWithColumns = () => {
  return http.request<BaseResponse<GetTablesWithColumnsResponse>>(
    "get",
    baseUrlApi("table/columns")
  );
};

export interface CodeGenRecordInfo {
  /** 主键ID */
  id: number;
  /** 表名 */
  tableName: string;
  /** 表描述 */
  tableComment: string;
  /** 字段列表 */
  columns: string;
  /** 生成参数 */
  options: string;
}
// 获取代码生成记录详情
export const getCodeGenRecordDetail = (id: number) => {
  return http.request<BaseResponse<CodeGenRecordInfo>>(
    "get",
    baseUrlApi(`generate/record/${id}`)
  );
};

/** 表单模式选项 */
export interface FormModeOption {
  /** 显示标签 */
  label: string;
  /** 选项值 */
  value: string;
}

/** 表单验证选项 */
export interface FormValidationOption {
  /** 显示标签 */
  label: string;
  /** 选项值 */
  value: string;
}

/** 查询条件选项 */
export interface WhereModeOption {
  /** 显示标签 */
  label: string;
  /** 选项值 */
  value: string;
}

/** 字段配置选项响应数据 */
export interface ColumnConfigOptionsData {
  /** 表单模式选项 */
  formModes: FormModeOption[];
  /** 表单验证选项 */
  formValidations: FormValidationOption[];
  /** 查询条件选项 */
  whereModes: WhereModeOption[];
}

// 获取字段配置选项
export const getColumnConfigOptions = () => {
  return http.request<BaseResponse<ColumnConfigOptionsData>>(
    "get",
    baseUrlApi("generate/column/options")
  );
};

/**
 * 删除代码生成记录
 * @param id 记录ID
 * @returns 删除结果
 */
export const deleteCodeGenRecord = (id: number) => {
  return http.request<BaseResponse<null>>(
    "delete",
    baseUrlApi(`generate/record/${id}`)
  );
};

/** 代码生成记录参数 */
export interface CodeGenRecordParams {
  /** 表名 */
  tableName: string;
  /** 表注释 */
  tableComment?: string;
  /** 包名 */
  packageName: string;
  /** 模块名 */
  moduleName: string;
  /** 配置选项 */
  options?: string;
  /** 字段信息 */
  columns?: string;
}

/**
 * 更新代码生成记录
 * @param id 记录ID
 * @param data 更新数据
 * @returns 更新结果
 */
export const updateCodeGenRecord = (id: number, data: CodeGenRecordParams) => {
  return http.request<BaseResponse<null>>(
    "put",
    baseUrlApi(`generate/record/${id}`),
    { data }
  );
};

/**
 * 代码生成
 * @param id 记录ID
 * @param data 代码生成参数
 * @returns 代码生成结果
 */
export const CodeGenRecord = (id: number, data: CodeGenRecordParams) => {
  return http.request<BaseResponse<null>>(
    "post",
    baseUrlApi(`generate/record/${id}`),
    { data }
  );
};

/** 根据表名获取字段信息响应 */
export interface GetTableColumnsResponse {
  /** 表名 */
  tableName: string;
  /** 字段列表 */
  columns: ColumnInfo[];
}

/** 根据表名获取字段信息 */
export const getTableColumns = (tableName: string) => {
  return http.request<BaseResponse<GetTableColumnsResponse>>(
    "get",
    baseUrlApi(`table/${tableName}/columns`)
  );
};
