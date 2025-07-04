package attachment

import (
	"context"
	"fmt"
	"mime/multipart"

	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "server/app/admin/api/attachment/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/middleware"
	"server/app/admin/internal/model/entity"
)

type sAttachment struct{}

func New() *sAttachment {
	return &sAttachment{}
}

// GetList 获取附件列表
func (s *sAttachment) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}

	// 构建查询条件
	m := dao.Attachment.Ctx(ctx)

	if req.FileName != nil {
		m = m.WhereLike(dao.Attachment.Columns().FileName, "%"+*req.FileName+"%")
	}
	if req.FileExt != nil {
		m = m.WhereLike(dao.Attachment.Columns().FileExt, "%"+*req.FileExt+"%")
	}
	if req.IsImage != nil {
		m = m.Where(dao.Attachment.Columns().IsImage, *req.IsImage)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询附件总数失败")
	}
	res.Total = int(total)

	// 分页查询
	var list []entity.Attachment
	err = m.Page(req.CurrentPage, req.PageSize).
		OrderDesc(dao.Attachment.Columns().CreatedAt).
		Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询附件列表失败")
	}

	// 转换数据格式
	if err := gconv.Scan(list, &res.List); err != nil {
		return nil, gerror.Wrap(err, "数据转换失败")
	}

	res.CurrentPage = req.CurrentPage
	return res, nil
}

// Upload 上传附件
func (s *sAttachment) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	res = &v1.UploadRes{}

	// 从请求中获取文件
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	if file == nil {
		return nil, gerror.New("请选择要上传的文件")
	}

	// 验证文件大小（限制为50MB）
	maxSize := int64(50 * 1024 * 1024)
	if file.Size > maxSize {
		return nil, gerror.New("文件大小不能超过50MB")
	}

	// 验证文件类型安全性
	if err = s.validateFileType(file.FileHeader); err != nil {
		return nil, err
	}

	// 获取文件信息
	originalName := file.Filename
	extension := strings.ToLower(filepath.Ext(originalName))
	mimeType := file.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = s.getMimeTypeByExtension(extension)
	}

	// 生成文件名
	fileName := s.generateFileName(originalName)
	if req.Name != nil && *req.Name != "" {
		fileName = *req.Name + extension
	}

	// 确定文件分类
	isImage := s.isImageFile(mimeType, extension)

	// 创建上传目录
	uploadDir := fmt.Sprintf("uploads/%s", time.Now().Format("2006/01/02"))
	fmt.Println("上传目录", uploadDir)
	if err = gfile.Mkdir(uploadDir); err != nil {
		return nil, gerror.Wrap(err, "创建上传目录失败")
	}

	// 保存文件
	filePath := filepath.Join(uploadDir)
	if _, err = file.Save(filePath); err != nil {
		return nil, gerror.Wrap(err, "保存文件失败")
	}

	// 生成访问URL
	baseURL := g.Cfg().MustGet(ctx, "server.baseUrl", "http://localhost:8000").String()
	fileURL := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), filePath)

	// 从上下文获取当前用户信息
	var uploaderId uint64
	var uploaderName string

	// 使用中间件提供的用户信息
	if userID, ok := ctx.Value(middleware.CtxUserID).(uint64); ok {
		uploaderId = userID
	}
	if username, ok := ctx.Value(middleware.CtxUsername).(string); ok {
		uploaderName = username
	}

	// 保存到数据库
	data := g.Map{
		dao.Attachment.Columns().FileName:     fileName,
		dao.Attachment.Columns().OriginalName: originalName,
		dao.Attachment.Columns().FilePath:     filePath,
		dao.Attachment.Columns().FileUrl:      fileURL + "/" + originalName,
		dao.Attachment.Columns().FileSize:     file.Size,
		dao.Attachment.Columns().FileType:     mimeType,
		dao.Attachment.Columns().FileExt:      extension,
		dao.Attachment.Columns().IsImage:      gconv.Int(isImage),
		dao.Attachment.Columns().UploaderId:   uploaderId,
		dao.Attachment.Columns().UploaderName: uploaderName,
		dao.Attachment.Columns().Status:       1, // 默认启用
	}

	if req.Remark != nil {
		data[dao.Attachment.Columns().Remark] = *req.Remark
	}

	// 如果是图片，生成缩略图URL（可选）
	if isImage {
		thumbnailURL := s.generateThumbnailURL(fileURL)
		data[dao.Attachment.Columns().ThumbnailUrl] = thumbnailURL
	}

	id, err := dao.Attachment.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		// 删除已上传的文件
		gfile.Remove(filePath)
		return nil, gerror.Wrap(err, "保存附件信息失败")
	}

	res.Id = uint64(id)
	res.Url = fileURL
	return res, nil
}

// Update 更新附件
func (s *sAttachment) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	res = &v1.UpdateRes{}

	// 检查附件是否存在
	count, err := dao.Attachment.Ctx(ctx).Where(dao.Attachment.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询附件失败")
	}
	if count == 0 {
		return nil, gerror.New("附件不存在")
	}

	// 动态构建更新数据
	updateData := g.Map{}

	if req.FileName != nil {
		// 更新文件名时需要保持扩展名
		var attachment entity.Attachment
		err = dao.Attachment.Ctx(ctx).Where(dao.Attachment.Columns().Id, req.Id).Scan(&attachment)
		if err != nil {
			return nil, gerror.Wrap(err, "查询附件信息失败")
		}
		newFileName := *req.FileName
		updateData[dao.Attachment.Columns().FileName] = newFileName
	}

	if req.Remark != nil {
		updateData[dao.Attachment.Columns().Remark] = *req.Remark
	}

	// 检查是否有字段需要更新
	if len(updateData) == 0 {
		return nil, gerror.New("没有需要更新的字段")
	}

	// 更新数据
	_, err = dao.Attachment.Ctx(ctx).Where(dao.Attachment.Columns().Id, req.Id).Data(updateData).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新附件失败")
	}

	return res, nil
}

// Delete 删除附件
func (s *sAttachment) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	res = &v1.DeleteRes{}

	// 查询附件信息
	var attachment entity.Attachment
	err = dao.Attachment.Ctx(ctx).Where(dao.Attachment.Columns().Id, req.Id).Scan(&attachment)
	if err != nil {
		return nil, gerror.Wrap(err, "查询附件失败")
	}
	if attachment.Id == 0 {
		return nil, gerror.New("附件不存在")
	}

	// 删除数据库记录
	_, err = dao.Attachment.Ctx(ctx).Where(dao.Attachment.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除附件记录失败")
	}

	// 删除物理文件
	if attachment.FilePath != "" && gfile.Exists(attachment.FilePath) {
		if err := gfile.Remove(attachment.FilePath); err != nil {
			g.Log().Warning(ctx, "删除物理文件失败", g.Map{
				"path":  attachment.FilePath,
				"error": err,
			})
		}
	}

	return res, nil
}

// BatchDelete 批量删除附件
func (s *sAttachment) BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error) {
	res = &v1.BatchDeleteRes{}

	if len(req.Ids) == 0 {
		return nil, gerror.New("请选择要删除的附件")
	}
	fmt.Println(req.Ids, "批量删除附件")
	// 查询要删除的附件信息
	var attachments []entity.Attachment
	err = dao.Attachment.Ctx(ctx).WhereIn(dao.Attachment.Columns().Id, req.Ids).Scan(&attachments)
	if err != nil {
		return nil, gerror.Wrap(err, "查询附件信息失败")
	}

	// 删除数据库记录
	_, err = dao.Attachment.Ctx(ctx).WhereIn(dao.Attachment.Columns().Id, req.Ids).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除附件记录失败")
	}

	// 删除物理文件
	for _, attachment := range attachments {
		if attachment.FilePath != "" && gfile.Exists(attachment.FilePath) {
			if err := gfile.Remove(attachment.FilePath); err != nil {
				g.Log().Warning(ctx, "删除物理文件失败", g.Map{
					"path":  attachment.FilePath,
					"error": err,
				})
			}
		}
	}

	return res, nil
}

// Download 下载附件
func (s *sAttachment) Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error) {
	res = &v1.DownloadRes{}

	// 查询附件信息
	var attachment entity.Attachment
	err = dao.Attachment.Ctx(ctx).Where(dao.Attachment.Columns().Id, req.Id).Scan(&attachment)
	if err != nil {
		return nil, gerror.Wrap(err, "查询附件失败")
	}
	if attachment.Id == 0 {
		return nil, gerror.New("附件不存在")
	}

	// 检查文件状态
	if attachment.Status != 1 {
		return nil, gerror.New("文件已被禁用")
	}

	// 检查文件是否存在
	if !gfile.Exists(attachment.FilePath) {
		return nil, gerror.New("文件不存在")
	}

	// 读取文件内容
	content := gfile.GetBytes(attachment.FilePath)
	if content == nil {
		return nil, gerror.New("读取文件失败")
	}

	res.FileName = attachment.OriginalName
	res.Content = content
	return res, nil
}

// validateFileType 验证文件类型安全性
func (s *sAttachment) validateFileType(file *multipart.FileHeader) error {
	// 危险文件扩展名黑名单
	dangerousExts := []string{".exe", ".bat", ".cmd", ".com", ".pif", ".scr", ".vbs", ".js", ".jar", ".php", ".asp", ".jsp"}
	ext := strings.ToLower(filepath.Ext(file.Filename))

	for _, dangerousExt := range dangerousExts {
		if ext == dangerousExt {
			return gerror.New("不允许上传此类型的文件")
		}
	}
	return nil
}

// isImageFile 判断是否为图片文件
func (s *sAttachment) isImageFile(mimeType, extension string) bool {
	if strings.HasPrefix(mimeType, "image/") {
		return true
	}
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg"}
	for _, ext := range imageExts {
		if extension == ext {
			return true
		}
	}
	return false
}

// getMimeTypeByExtension 根据扩展名获取MIME类型
func (s *sAttachment) getMimeTypeByExtension(extension string) string {
	mimeTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".txt":  "text/plain",
	}
	if mimeType, exists := mimeTypes[extension]; exists {
		return mimeType
	}
	return "application/octet-stream"
}

// generateFileName 生成唯一文件名
func (s *sAttachment) generateFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, ext)
	timestamp := time.Now().UnixNano()
	hash := gmd5.MustEncryptString(fmt.Sprintf("%s_%d", name, timestamp))
	return fmt.Sprintf("%s_%s%s", name, hash[:8], ext)
}

// generateThumbnailURL 生成缩略图URL（简单实现）
func (s *sAttachment) generateThumbnailURL(originalURL string) string {
	// 这里可以实现缩略图生成逻辑
	// 简单实现：返回原图URL，实际项目中可以集成图片处理服务
	return originalURL
}
