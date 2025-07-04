export interface CodeConfigForm {
  tableName: string;
  tableComment: string;
  className: string;
  packageName: string;
  moduleName: string;
  businessName: string;
  author: string;
  email: string;
  genType: string;
  genPath: string;
  tableOptions: {
    treeCode: boolean;
    treeParentCode: boolean;
    treeName: boolean;
    parentMenuId: boolean;
    sync: boolean;
    subTable: boolean;
    crud: boolean;
    rest: boolean;
  };
  functionName: string;
  functionAuthor: string;
  remark: string;
}

export interface CodeConfigItem {
  id: number;
  tableName: string;
  tableComment: string;
  moduleName: string;
  packageName: string;
  author: string;
  createTime: Date;
}
