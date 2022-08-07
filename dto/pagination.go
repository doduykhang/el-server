package dto

type PaginationRequest struct {
	PageSize uint `form:"pageSize"`
	PageNum  uint `form:"pageNum"`
}
