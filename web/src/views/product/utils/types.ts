interface FormItemProps {
  /** 用于判断是新增还是修改 */
  title: string;
  id?: number;
  name?: string;
  description?: string;
  price?: string;
  stock?: string;
  main_image?: string;
  carousel_images?: string;
  attachment?: string;
  status?: string;
}

interface FormProps {
  formInline: FormItemProps;
}

export type { FormItemProps, FormProps };