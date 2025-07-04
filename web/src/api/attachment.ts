import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse, PageResponse, PageParams } from "./common/types";

/** 附件信息 */
export interface AttachmentInfo {
  /** 附件ID */
  id: number;
  /** 文件名 */
  fileName: string;
  /** 原始文件名 */
  originalName: string;
  /** 文件大小（字节） */
  fileSize: number;
  /** 文件类型 */
  fileType: string;
  /** 文件扩展名 */
  fileExt: string;
  /** 文件路径 */
  filePath: string;
  /** 文件URL */
  fileUrl: string;
  /** 是否为图片 */
  isImage: boolean;
  /** 缩略图URL */
  thumbnailUrl?: string;
  /** 上传者ID */
  uploaderId: number;
  /** 上传者名称 */
  uploaderName: string;
  /** 状态 */
  status: number;
  /** 备注 */
  remark: string;
  /** 创建时间 */
  createTime: string;
  /** 更新时间 */
  updateTime: string;
}

/** 附件列表查询参数 */
export interface AttachmentListParams extends PageParams {
  /** 文件名（模糊查询） */
  fileName?: string;
  /** 文件扩展名 */
  fileExt?: string;
  /** 是否为图片 */
  isImage?: boolean;
  /** 状态筛选 */
  status?: number;
}

/** 上传文件参数 - 简化为只接收文件 */
export interface UploadFileParams {
  /** 文件对象 */
  file: File;
}

/** 上传文件 - 修改为只接收文件 */
export const uploadFile = (file: File) => {
  const formData = new FormData();
  formData.append("file", file);

  return http.request<BaseResponse<AttachmentInfo>>(
    "post",
    baseUrlApi("attachment/upload"),
    { data: formData },
    {
      headers: {
        "Content-Type": "multipart/form-data"
      }
    }
  );
};

/** 更新附件参数 */
export interface UpdateAttachmentParams {
  /** 附件ID */
  id: number;
  /** 文件名 */
  fileName?: string;
  /** 备注 */
  remark?: string;
  /** 状态 */
  status?: number;
}

/** 批量删除参数 */
export interface BatchDeleteParams {
  /** 附件ID数组 */
  ids: number[];
}

/** 获取附件列表 */
export const getAttachmentList = (params?: AttachmentListParams) => {
  return http.request<PageResponse<AttachmentInfo>>(
    "get",
    baseUrlApi("attachment"),
    { params }
  );
};

/** 更新附件信息 */
export const updateAttachment = (data: UpdateAttachmentParams) => {
  return http.request<BaseResponse<AttachmentInfo>>(
    "put",
    baseUrlApi(`attachment/${data.id}`),
    { data }
  );
};

/** 删除附件 */
export const deleteAttachment = (id: number) => {
  return http.request<BaseResponse<null>>(
    "delete",
    baseUrlApi(`attachment/${id}`)
  );
};

/** 批量删除附件 */
export const batchDeleteAttachments = (data: BatchDeleteParams) => {
  return http.request<BaseResponse<null>>(
    "post",
    baseUrlApi("attachment/batch"),
    { data }
  );
};

/** 下载附件 */
export const downloadAttachment = (id: number) => {
  return http.request<Blob>(
    "get",
    baseUrlApi(`attachment/download/${id}`),
    {},
    {
      responseType: "blob"
    }
  );
};
