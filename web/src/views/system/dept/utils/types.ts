interface FormItemProps {
  higherDeptOptions: Record<string, unknown>[];
  id?: number; // 编辑时传入id
  parentId: number;
  name: string;
  principal: string;
  phone: string;
  email: string;
  sort: number;
  status: number;
  remark: string;
}
interface FormProps {
  formInline: FormItemProps;
}

export type { FormItemProps, FormProps };
