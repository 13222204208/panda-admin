package page

// ReqPage 分页请求参数
type ReqPage struct {
	CurrentPage int `json:"currentPage" v:"required#请输入页码" d:"1" dc:"当前页码"`
	PageSize    int `json:"pageSize" v:"required#请输入每页数量" d:"10" dc:"每页数量"`
}

// ResPage 分页返回参数
type ResPage struct {
	Total       int `json:"total" dc:"总数"`
	CurrentPage int `json:"currentPage" dc:"当前页码"`
}
