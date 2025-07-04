interface FormItemProps {
  /** ID */
  id?: string | number;
  /** 表名称 */
  tableName: string;
  /** 表描述 */
  tableComment: string;
  /** 实体类名 */
  className: string;
  /** 创建时间 */
  createTime?: string;
  /** 更新时间 */
  updateTime?: string;
}

interface FormProps {
  formInline: FormItemProps;
}

interface TableItemProps extends FormItemProps {
  /** 表ID */
  id: number;
  /** 创建时间 */
  createTime: string;
  /** 更新时间 */
  updateTime: string;
}

export type { FormItemProps, FormProps, TableItemProps };
