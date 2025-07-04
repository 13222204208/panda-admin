interface FormItemProps {
  /** ID */
  id?: string | number;
  /** 字典标题 */
  title: string;
  /** 字典类型 */
  dictType: string;
  /** 字典标签 */
  dictLabel: string;
  /** 字典值 */
  dictValue: string;
  /** 排序值 */
  sort: number;
  /** 状态（1启用，0禁用） */
  status: number;
  /** 备注说明 */
  remark?: string;
  /** 创建时间 */
  createTime?: string;
  /** 更新时间 */
  updateTime?: string;
}

// 批量字典项
interface BatchDictItem {
  dictLabel: string;
  dictValue: string;
  sort: number;
  remark?: string;
}

// 批量表单数据结构（统一定义）
interface BatchFormItemProps {
  /** 字典类型 */
  dictType: string;
  /** 字典标题 */
  title: string;
  /** 字典项列表 */
  dictItems: BatchDictItem[];
}

interface FormProps {
  formInline: FormItemProps;
}

interface BatchFormProps {
  formInline: BatchFormItemProps;
}

interface TableItemProps extends FormItemProps {
  /** 字典ID */
  id: number;
  /** 创建时间 */
  createTime: string;
  /** 更新时间 */
  updateTime: string;
}

export type {
  FormItemProps,
  FormProps,
  BatchFormProps,
  BatchFormItemProps,
  BatchDictItem,
  TableItemProps
};
