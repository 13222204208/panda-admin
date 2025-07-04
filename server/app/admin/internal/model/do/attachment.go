// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Attachment is the golang structure of table attachment for DAO operations like Where/Data.
type Attachment struct {
	g.Meta       `orm:"table:attachment, do:true"`
	Id           interface{} // 主键ID
	FileName     interface{} // 文件名
	OriginalName interface{} // 原始文件名
	FileSize     interface{} // 文件大小（字节）
	FileType     interface{} // 文件类型（MIME）
	FileExt      interface{} // 文件扩展名
	FilePath     interface{} // 文件路径
	FileUrl      interface{} // 文件URL
	IsImage      interface{} // 是否为图片
	ThumbnailUrl interface{} // 缩略图URL
	UploaderId   interface{} // 上传者ID
	UploaderName interface{} // 上传者名称
	Status       interface{} // 状态（1正常，0禁用/删除）
	Remark       interface{} // 备注
	CreatedAt    *gtime.Time // 上传时间
	UpdatedAt    *gtime.Time // 更新时间
}
