package utility

import (
	"time"
)

// FormatTimeToChineseStyle 将时间格式化为中文样式：年月日 时分秒
func FormatTimeToChineseStyle(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006年01月02日 15:04:05")
}

// ParseAndFormatTime 解析时间字符串并格式化为中文样式
func ParseAndFormatTime(timeStr string) string {
	if timeStr == "" {
		return ""
	}
	
	// 尝试多种时间格式解析
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05",
		time.RFC3339,
		time.DateTime,
	}
	
	for _, layout := range layouts {
		if t, err := time.Parse(layout, timeStr); err == nil {
			return FormatTimeToChineseStyle(t)
		}
	}
	
	// 如果解析失败，返回原字符串
	return timeStr
}