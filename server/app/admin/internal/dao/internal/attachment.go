// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttachmentDao is the data access object for the table attachment.
type AttachmentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AttachmentColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AttachmentColumns defines and stores column names for the table attachment.
type AttachmentColumns struct {
	Id           string // 主键ID
	FileName     string // 文件名
	OriginalName string // 原始文件名
	FileSize     string // 文件大小（字节）
	FileType     string // 文件类型（MIME）
	FileExt      string // 文件扩展名
	FilePath     string // 文件路径
	FileUrl      string // 文件URL
	IsImage      string // 是否为图片
	ThumbnailUrl string // 缩略图URL
	UploaderId   string // 上传者ID
	UploaderName string // 上传者名称
	Status       string // 状态（1正常，0禁用/删除）
	Remark       string // 备注
	CreatedAt    string // 上传时间
	UpdatedAt    string // 更新时间
}

// attachmentColumns holds the columns for the table attachment.
var attachmentColumns = AttachmentColumns{
	Id:           "id",
	FileName:     "file_name",
	OriginalName: "original_name",
	FileSize:     "file_size",
	FileType:     "file_type",
	FileExt:      "file_ext",
	FilePath:     "file_path",
	FileUrl:      "file_url",
	IsImage:      "is_image",
	ThumbnailUrl: "thumbnail_url",
	UploaderId:   "uploader_id",
	UploaderName: "uploader_name",
	Status:       "status",
	Remark:       "remark",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewAttachmentDao creates and returns a new DAO object for table data access.
func NewAttachmentDao(handlers ...gdb.ModelHandler) *AttachmentDao {
	return &AttachmentDao{
		group:    "default",
		table:    "attachment",
		columns:  attachmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AttachmentDao) Columns() AttachmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AttachmentDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
