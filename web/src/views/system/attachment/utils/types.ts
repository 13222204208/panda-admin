/** 附件信息 */
export interface AttachmentInfo {
  /** 附件ID */
  id?: number;
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
  uploaderId?: number;
  /** 上传者名称 */
  uploaderName?: string;
  /** 状态 */
  status: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  createTime?: string;
  /** 更新时间 */
  updateTime?: string;
}

/** 附件表单属性 */
export interface FormItemProps {
  id?: number;
  /** 文件名 */
  fileName: string;
  /** 原始文件名 */
  originalName: string;
  /** 文件类型 */
  fileType: string;
  /** 备注 */
  remark: string;
}

/** 表单组件属性 */
export interface FormProps {
  formInline: FormItemProps;
}

/** 上传文件参数 */
export interface UploadFileParams {
  /** 文件对象 */
  file: File;
  /** 文件类型分类 */
  category?: string;
  /** 备注 */
  remark?: string;
}