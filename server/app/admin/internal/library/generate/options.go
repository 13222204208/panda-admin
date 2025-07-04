package generate

import (
	"encoding/json"
	"fmt"
	"server/app/admin/internal/consts"
	"strings"
)

// GenerateOptionsToString 将生成选项转换为字符串
func GenerateOptionsToString(options *consts.GenerateOptions) string {
	var optionList []string

	if options.Create {
		optionList = append(optionList, consts.OptionCreate)
	}
	if options.Update {
		optionList = append(optionList, consts.OptionUpdate)
	}
	if options.Delete {
		optionList = append(optionList, consts.OptionDelete)
	}
	if options.BatchDelete {
		optionList = append(optionList, consts.OptionBatchDelete)
	}
	if options.List {
		optionList = append(optionList, consts.OptionList)
	}

	return strings.Join(optionList, ",")
}

// ParseGenerateOptionsFromConfig 从配置对象解析生成选项
func ParseGenerateOptionsFromConfig(config map[string]interface{}) *consts.GenerateOptions {
	options := &consts.GenerateOptions{}

	// 从配置中读取各个选项
	if hasAdd, ok := config["hasAdd"].(bool); ok {
		options.Create = hasAdd
	}

	if hasEdit, ok := config["hasEdit"].(bool); ok {
		options.Update = hasEdit
	}

	if hasQuery, ok := config["hasQuery"].(bool); ok {
		options.List = hasQuery // 查询对应列表功能
	}

	if hasDelete, ok := config["hasDelete"].(bool); ok {
		options.Delete = hasDelete
	}

	if hasBatchDelete, ok := config["hasBatchDelete"].(bool); ok {
		options.BatchDelete = hasBatchDelete
	}

	// 从配置中读取菜单相关选项
	if menuIcon, ok := config["menuIcon"].(string); ok {
		options.MenuIcon = menuIcon
	}

	if menuName, ok := config["menuName"].(string); ok {
		options.MenuName = menuName
	}

	if parentMenuId, ok := config["parentMenuId"].(int); ok {
		options.ParentMenuId = parentMenuId
	}

	return options
}

// ParseGenerateOptionsFromJSON 从JSON字符串解析生成选项（支持你的数据格式）
func ParseGenerateOptionsFromJSON(jsonStr string) (*consts.GenerateOptions, error) {
	var config map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &config); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	return ParseGenerateOptionsFromConfig(config), nil
}
