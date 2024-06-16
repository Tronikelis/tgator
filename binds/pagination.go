package binds

type PaginationBind struct {
	page  int32 `query:"page"`
	limit int32 `query:"limit"`
}

func (p *PaginationBind) Page() int32 {
	if p.page <= 0 {
		return p.page
	}

	return p.page
}

func (p *PaginationBind) Limit() int32 {
	if p.limit <= 0 {
		return 50
	}

	if p.limit >= 100 {
		return 100
	}

	return p.limit
}

func (p *PaginationBind) Offset() int32 {
	return p.Page() * p.Limit()
}
