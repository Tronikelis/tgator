package binds

type GetMessagesBind struct {
	PaginationBind
	Search  string `query:"search"`
	OrderBy string `query:"orderBy"`
}
