package service

type basePage struct {
	PageNum  int `form:"pageNum"`  // 当前页码
	PageSize int `form:"pageSize"` // 每页数
}
