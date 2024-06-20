package dtos

import "tgator/binds"

type PaginationDTO[T any] struct {
	page   int32
	offset int32
	limit  int32
	data   []T
}

func (p PaginationDTO[T]) FromBind(b binds.PaginationBind) PaginationDTO[T] {
	return PaginationDTO[T]{
		page:   b.Page(),
		offset: b.Offset(),
		limit:  b.Limit(),
	}
}

func (p *PaginationDTO[T]) SetData(data []T) *PaginationDTO[T] {
	p.data = data
	return p
}

func (p *PaginationDTO[T]) Offset() uint {
	return uint(p.offset)
}

func (p *PaginationDTO[T]) Limit() uint {
	return uint(p.limit)
}
