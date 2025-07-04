package utility

import (
	"context"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadConfig 上传配置
type UploadConfig struct {
	UploadDir    string   // 上传目录
	AllowedTypes []string // 允许的文件类型
	FilePrefix   string   // 文件名前缀
}

// DefaultImageConfig 默认图片上传配置
var DefaultImageConfig = UploadConfig{
	UploadDir:    "uploads/images",
	AllowedTypes: []string{"image/jpeg", "image/png", "image/gif", "image/webp"},
	FilePrefix:   "img",
}

// AvatarConfig 头像上传配置
var AvatarConfig = UploadConfig{
	UploadDir:    "uploads/avatars",
	AllowedTypes: []string{"image/jpeg", "image/png", "image/gif", "image/webp"},
	FilePrefix:   "avatar",
}

// SaveBase64Image 保存base64图片到文件系统
func SaveBase64Image(ctx context.Context, base64Data string, userId uint64, config UploadConfig) (string, error) {
	// 检查 context 是否已取消
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	// 解析base64数据
	parts := strings.Split(base64Data, ",")
	if len(parts) != 2 {
		return "", gerror.New("无效的base64图片数据")
	}

	// 获取文件类型
	headerParts := strings.Split(parts[0], ";")
	if len(headerParts) < 1 {
		return "", gerror.New("无法识别图片类型")
	}
	mimeType := strings.TrimPrefix(headerParts[0], "data:")

	// 检查是否为允许的文件类型
	allowed := false
	for _, allowedType := range config.AllowedTypes {
		if mimeType == allowedType {
			allowed = true
			break
		}
	}
	if !allowed {
		return "", gerror.New("不支持的图片格式")
	}

	// 获取文件扩展名
	var ext string
	switch mimeType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	case "image/webp":
		ext = ".webp"
	default:
		return "", gerror.New("不支持的图片格式")
	}

	// 解码base64数据
	imageData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", gerror.Wrap(err, "解码base64数据失败")
	}

	// 再次检查 context
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	// 生成文件名
	fileName := fmt.Sprintf("%s_%d_%d%s", config.FilePrefix, userId, gtime.Now().Timestamp(), ext)

	// 创建上传目录
	if err := gfile.Mkdir(config.UploadDir); err != nil {
		return "", gerror.Wrap(err, "创建上传目录失败")
	}

	// 保存文件
	filePath := filepath.Join(config.UploadDir, fileName)
	if err := gfile.PutBytes(filePath, imageData); err != nil {
		return "", gerror.Wrap(err, "保存文件失败")
	}

	// 返回访问URL
	return "/" + strings.ReplaceAll(filePath, "\\", "/"), nil
}

// SaveBase64Avatar 保存base64头像图片（便捷方法）
func SaveBase64Avatar(ctx context.Context, base64Data string, userId uint64) (string, error) {
	return SaveBase64Image(ctx, base64Data, userId, AvatarConfig)
}

// SaveBase64GeneralImage 保存base64通用图片（便捷方法）
func SaveBase64GeneralImage(ctx context.Context, base64Data string, userId uint64) (string, error) {
	return SaveBase64Image(ctx, base64Data, userId, DefaultImageConfig)
}