package binds

type GetMessagesBind struct {
	PaginationBind
	Search  string `query:"search"`
	OrderBy string `query:"orderBy"`
}

type CreateMessageBind struct {
	SourceId int32  `json:"sourceId"`
	Raw      string `json:"raw"`
}
