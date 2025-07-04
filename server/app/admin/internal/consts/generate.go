package consts

// 代码生成选项常量
const (
	// CRUD操作选项
	OptionCreate      = "create"      // 生成创建接口
	OptionUpdate      = "update"      // 生成更新接口
	OptionDelete      = "delete"      // 生成删除接口
	OptionBatchDelete = "batchDelete" // 生成批量删除接口
	OptionList        = "list"        // 生成列表接口

	// 默认选项（全部生成）
	DefaultOptions = "create,update,delete,batchDelete,list"
)

// GenerateOptions 生成选项结构
type GenerateOptions struct {
	Create      bool `json:"create"`      // 是否生成创建接口
	Update      bool `json:"update"`      // 是否生成更新接口
	Delete      bool `json:"delete"`      // 是否生成删除接口
	BatchDelete bool `json:"batchDelete"` // 是否生成批量删除接口
	List        bool `json:"list"`        // 是否生成列表接口
	//menuIcon
	MenuIcon string `json:"menuIcon"` // 菜单图标
	//menuName
	MenuName string `json:"menuName"` // 菜单名称
	//parentMenuId
	ParentMenuId int `json:"parentMenuId"` // 父菜单ID
}
