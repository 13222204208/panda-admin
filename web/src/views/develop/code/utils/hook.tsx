import { ref } from "vue";
import { getMenuTree } from "@/api/menu";
import { getTablesWithColumns } from "@/api/generate";

export function useCodeConfig() {
  // 创建响应式的菜单树数据
  const menuTreeData = ref([]);
  // 创建响应式的表字段数据
  const tableColumnsData = ref([]);

  // 获取表字段数据
  const fetchTableColumns = async () => {
    try {
      const response = await getTablesWithColumns();
      console.log("表字段数据:", response.data.list);
      // 更新响应式数据
      tableColumnsData.value = response.data.list;
      return response.data;
    } catch (error) {
      console.error("获取表字段数据失败:", error);
    }
  };

  // 获取菜单树数据
  const fetchMenuTree = async () => {
    try {
      const response = await getMenuTree();
      console.log("菜单树数据:", response.data.tree);
      // 更新响应式数据
      menuTreeData.value = response.data.tree;
      return response.data.tree;
    } catch (error) {
      console.error("获取菜单树失败:", error);
    }
  };

  return {
    menuTreeData,
    tableColumnsData,
    fetchTableColumns,
    fetchMenuTree
  };
}
