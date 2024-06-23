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

func (p PaginationDTO[T]) FromBind(b binds.PaginationBind) PaginationDTO[T] {
	return PaginationDTO[T]{
		Page:   b.GetPage(),
		Offset: b.GetOffset(),
		Limit:  b.GetLimit(),
	}
}
