package binds

type GetMessagesBind struct {
	PaginationBind
	Search string `query:"search"`
}
