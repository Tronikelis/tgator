package dtos

import (
	"tgator/binds"
)

type PaginationDTO[T any] struct {
	Page   int32
	Offset int32
	Limit  int32
	Data   []T
}

func (p *PaginationDTO[T]) SetFromBind(b binds.PaginationBind) *PaginationDTO[T] {
	p.Page = b.GetPage()
	p.Limit = b.GetLimit()
	p.Offset = b.GetOffset()

	return p
}
