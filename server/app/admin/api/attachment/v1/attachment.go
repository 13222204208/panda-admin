package v1

import (
	"server/app/admin/api/common/page"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AttachmentCommon 附件公共字段
type AttachmentCommon struct {
	FileName *string `json:"fileName,omitempty" v:"required#请输入附件名称" dc:"附件名称"`
	Remark   *string `json:"remark,omitempty" dc:"备注"`
}

// GetListReq 查询附件列表请求参数
type GetListReq struct {
	g.Meta `path:"/attachment" method:"get" tags:"附件管理" summary:"获取附件列表"`
	page.ReqPage
	FileName *string `json:"fileName,omitempty" dc:"附件名称"`
	FileExt  *string `json:"fileExt,omitempty" dc:"文件扩展名"`
	IsImage  *bool   `json:"isImage,omitempty" dc:"是否图片"`
}

// GetListRes 查询附件列表返回参数
type GetListRes struct {
	page.ResPage
	List []AttachmentInfo `json:"list" dc:"附件列表"`
}

// UploadReq 上传附件请求参数
type UploadReq struct {
	g.Meta   `path:"/attachment/upload" method:"post" tags:"附件管理" summary:"上传附件"`
	File     string  `json:"file" v:"required#请选择要上传的文件" dc:"文件内容（base64或multipart）"`
	Name     *string `json:"name,omitempty" dc:"自定义附件名称"`
	Category *string `json:"category,omitempty" dc:"文件分类"`
	Remark   *string `json:"remark,omitempty" dc:"备注"`
}

// UploadRes 上传附件返回参数
type UploadRes struct {
	Id  uint64 `json:"id" dc:"附件ID"`
	Url string `json:"url" dc:"访问URL"`
}

// UpdateReq 更新附件请求参数
type UpdateReq struct {
	g.Meta `path:"/attachment/{id}" method:"put" tags:"附件管理" summary:"更新附件"`
	Id     uint64 `json:"id" v:"required#请输入附件ID" dc:"附件ID"`
	AttachmentCommon
}

// UpdateRes 更新附件返回参数
type UpdateRes struct{}

// DeleteReq 删除附件请求参数
type DeleteReq struct {
	g.Meta `path:"/attachment/{id}" method:"delete" tags:"附件管理" summary:"删除附件"`
	Id     uint64 `json:"id" v:"required#请输入附件ID" dc:"附件ID"`
}

// DeleteRes 删除附件返回参数
type DeleteRes struct{}

// BatchDeleteReq 批量删除附件请求参数
type BatchDeleteReq struct {
	g.Meta `path:"/attachment/batch" method:"post" tags:"附件管理" summary:"批量删除附件"`
	Ids    []uint64 `json:"ids" v:"required#请选择要删除的附件" dc:"附件ID列表"`
}

// BatchDeleteRes 批量删除附件返回参数
type BatchDeleteRes struct{}

// DownloadReq 下载附件请求参数
type DownloadReq struct {
	g.Meta `path:"/attachment/{id}/download" method:"get" tags:"附件管理" summary:"下载附件"`
	Id     uint64 `json:"id" v:"required#请输入附件ID" dc:"附件ID"`
}

// DownloadRes 下载附件返回参数
type DownloadRes struct {
	FileName string `json:"fileName" dc:"文件名"`
	Content  []byte `json:"content" dc:"文件内容"`
}

// AttachmentInfo 附件信息
type AttachmentInfo struct {
	Id           uint64      `json:"id" dc:"主键ID"`
	FileName     string      `json:"fileName" dc:"文件名"`
	OriginalName string      `json:"originalName" dc:"原始文件名"`
	FileSize     uint64      `json:"fileSize" dc:"文件大小（字节）"`
	FileType     string      `json:"fileType" dc:"文件类型（MIME）"`
	FileExt      string      `json:"fileExt" dc:"文件扩展名"`
	FilePath     string      `json:"filePath" dc:"文件路径"`
	FileUrl      string      `json:"fileUrl" dc:"文件URL"`
	IsImage      int         `json:"isImage" dc:"是否为图片"`
	ThumbnailUrl string      `json:"thumbnailUrl" dc:"缩略图URL"`
	UploaderId   uint64      `json:"uploaderId" dc:"上传者ID"`
	UploaderName string      `json:"uploaderName" dc:"上传者名称"`
	Status       int         `json:"status" dc:"状态（1正常，0禁用/删除）"`
	Remark       string      `json:"remark" dc:"备注"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"上传时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
