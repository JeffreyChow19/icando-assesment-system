package dao

type MetaDao struct {
	Page      int   `json:"page"`
	TotalPage int   `json:"totalPage"`
	Limit     int   `json:"limit"`
	TotalItem int64 `json:"totalItem"`
}
