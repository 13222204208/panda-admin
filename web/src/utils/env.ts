// 环境变量工具函数
const { VITE_API_BASE_URL } = import.meta.env;

/**
 * 获取完整的API URL
 * @param url 相对路径
 * @returns 完整的API URL
 */
export const getApiUrl = (url: string): string => {
  // 如果已经是完整URL，直接返回
  if (url.startsWith('http')) {  
    return url;
  }
   
  // 确保URL以/开头
  const normalizedUrl = url.startsWith('/') ? url : `/${url}`;
  
  return `${VITE_API_BASE_URL}${normalizedUrl}`; 
};

/**
 * 获取API基础URL
 */
export const getApiBaseUrl = (): string => {
  return VITE_API_BASE_URL || "";  
};

/**
 * 获取完整的文件URL（用于图片、附件等）
 * @param fileUrl 文件相对路径
 * @returns 完整的文件URL
 */
export const getFileUrl = (fileUrl: string): string => {
  if (!fileUrl) return "";
  return getApiUrl(fileUrl);
};