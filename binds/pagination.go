package binds

type PaginationBind struct {
	Page  int32 `query:"page"`
	Limit int32 `query:"limit"`
}
