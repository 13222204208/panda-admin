export interface FileImageSelectorProps {
  /** 是否支持多选，默认false */
  multiple?: boolean;
  /** 最大选择数量，仅在multiple为true时生效 */
  maxCount?: number;
  /** 已选择的文件列表 */
  modelValue?: AttachmentInfo[] | AttachmentInfo | null;
  /** 文件类型：'image' | 'file' | 'all'，默认'all' */
  fileType?: "image" | "file" | "all";
  /** 允许的文件扩展名，如 ['jpg', 'png', 'pdf'] */
  allowedExtensions?: string[];
}

export interface FileImageSelectorEmits {
  (
    e: "update:modelValue",
    value: AttachmentInfo[] | AttachmentInfo | null
  ): void;
  (e: "confirm", value: AttachmentInfo[] | AttachmentInfo | null): void;
}

export interface FileImageSelectorInstance {
  openDialog: () => void;
}
