package dtos

type PaginationDTO[T any] struct {
	Offset int32
	Limit  int32
	Data   []T
}
