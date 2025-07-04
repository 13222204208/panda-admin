interface FormItemProps {
  /** 用于判断是新增还是修改 */
  title: string;
  id?: number;
  username?: string;
  email?: string;
  mobile?: string;
}

interface FormProps {
  formInline: FormItemProps;
}

export type { FormItemProps, FormProps };