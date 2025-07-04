// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Attachment is the golang structure for table attachment.
type Attachment struct {
	Id           uint64      `json:"id"           orm:"id"            description:"主键ID"`           // 主键ID
	FileName     string      `json:"fileName"     orm:"file_name"     description:"文件名"`            // 文件名
	OriginalName string      `json:"originalName" orm:"original_name" description:"原始文件名"`          // 原始文件名
	FileSize     uint64      `json:"fileSize"     orm:"file_size"     description:"文件大小（字节）"`       // 文件大小（字节）
	FileType     string      `json:"fileType"     orm:"file_type"     description:"文件类型（MIME）"`     // 文件类型（MIME）
	FileExt      string      `json:"fileExt"      orm:"file_ext"      description:"文件扩展名"`          // 文件扩展名
	FilePath     string      `json:"filePath"     orm:"file_path"     description:"文件路径"`           // 文件路径
	FileUrl      string      `json:"fileUrl"      orm:"file_url"      description:"文件URL"`          // 文件URL
	IsImage      int         `json:"isImage"      orm:"is_image"      description:"是否为图片"`          // 是否为图片
	ThumbnailUrl string      `json:"thumbnailUrl" orm:"thumbnail_url" description:"缩略图URL"`         // 缩略图URL
	UploaderId   uint64      `json:"uploaderId"   orm:"uploader_id"   description:"上传者ID"`          // 上传者ID
	UploaderName string      `json:"uploaderName" orm:"uploader_name" description:"上传者名称"`          // 上传者名称
	Status       int         `json:"status"       orm:"status"        description:"状态（1正常，0禁用/删除）"` // 状态（1正常，0禁用/删除）
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`             // 备注
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"上传时间"`           // 上传时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`           // 更新时间
}
