// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/attachment/v1"
)

type (
	IAttachment interface {
		// GetList 获取附件列表
		GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
		// Upload 上传附件
		Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
		// Update 更新附件
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		// Delete 删除附件
		Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
		// BatchDelete 批量删除附件
		BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error)
		// Download 下载附件
		Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error)
	}
)

var (
	localAttachment IAttachment
)

func Attachment() IAttachment {
	if localAttachment == nil {
		panic("implement not found for interface IAttachment, forgot register?")
	}
	return localAttachment
}

func RegisterAttachment(i IAttachment) {
	localAttachment = i
}
