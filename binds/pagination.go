package binds

type PaginationBind struct {
	Page  int32 `query:"page"`
	Limit int32 `query:"limit"`
}

func (p *PaginationBind) SafePage() int32 {
	if p.Page <= 0 {
		return p.Page
	}

	return p.Page
}

func (p *PaginationBind) SafeLimit() int32 {
	if p.Limit <= 0 {
		return 50
	}

	if p.Limit >= 100 {
		return 100
	}

	return p.Limit
}

func (p *PaginationBind) SafeOffset() int32 {
	return p.SafePage() * p.SafeLimit()
}
