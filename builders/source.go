package builders

import (
	"tgator/dtos"

	"github.com/doug-martin/goqu/v9"
)

type GetSourceMessagesBuilder struct {
	*GetMessagesBuilder
}

func NewGetSourceMessagesBuilder[T any](dw goqu.DialectWrapper, paginationDto dtos.PaginationDTO[T], sourceId int32) *GetSourceMessagesBuilder {
	b := NewGetMessagesBuilder(dw, paginationDto)

	b.builder = b.builder.Where(goqu.C("source_id").Eq(sourceId))

	return &GetSourceMessagesBuilder{b}
}
