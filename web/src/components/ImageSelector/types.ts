export interface ImageSelectorProps {
  /** 是否支持多选，默认false */
  multiple?: boolean;
  /** 最大选择数量，仅在multiple为true时生效 */
  maxCount?: number;
  /** 已选择的图片列表 */
  modelValue?: AttachmentInfo[] | AttachmentInfo | null;
}

export interface ImageSelectorEmits {
  (
    e: "update:modelValue",
    value: AttachmentInfo[] | AttachmentInfo | null
  ): void;
  (e: "confirm", value: AttachmentInfo[] | AttachmentInfo | null): void;
}

export interface ImageSelectorInstance {
  openDialog: () => void;
}
