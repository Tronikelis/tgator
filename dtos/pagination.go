package dtos

import (
	"math"
	"tgator/binds"
)

type PaginationDTO[T any] struct {
	Page   int32
	Offset int32
	Limit  int32
	Pages  int32
	Count  int32
	Data   []T
}

func (p *PaginationDTO[T]) SetFromBind(b binds.PaginationBind) *PaginationDTO[T] {
	p.Page = b.SafePage()
	p.Limit = b.SafeLimit()
	p.Offset = b.SafeOffset()

	return p
}

func (p *PaginationDTO[T]) SetPages(rowCount int32) *PaginationDTO[T] {
	p.Count = rowCount
	p.Pages = int32(math.Ceil(float64(rowCount) / float64(p.Limit)))

	return p
}
