package builders

import (
	"strings"
	"tgator/db"
	"tgator/dtos"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type GetMessagesBuilder struct {
	builder *goqu.SelectDataset
	table   db.Table
}

func NewGetMessagesBuilder[T any](dw goqu.DialectWrapper, paginationDto dtos.PaginationDTO[T]) *GetMessagesBuilder {
	table := db.Table("messages")

	builder := dw.
		From(string(table)).
		Select(goqu.I(table.WithSuffix("*"))).
		Limit(uint(paginationDto.Limit)).
		Offset(uint(paginationDto.Offset)).
		Order(goqu.I(table.WithSuffix("id")).Desc())

	return &GetMessagesBuilder{builder, table}
}

func (b *GetMessagesBuilder) UnwrapSelectDataset() *goqu.SelectDataset {
	return b.builder
}

func (b *GetMessagesBuilder) ToSQL() (string, []interface{}, error) {
	return b.builder.ToSQL()
}

func (b *GetMessagesBuilder) JoinSources() *GetMessagesBuilder {
	b.builder = b.builder.
		Select(goqu.I("sources.*")).
		LeftJoin(
			goqu.T("sources"), goqu.On(goqu.I("sources.id").Eq(goqu.I(b.table.WithSuffix("source_id")))),
		)

	return b
}

func (b *GetMessagesBuilder) OrderBy(orderBy string) *GetMessagesBuilder {
	b.builder = b.builder.ClearOrder()

	var order exp.OrderedExpression

	col := b.table.WithSuffix("id")

	switch orderBy {
	case "asc":
		order = goqu.I(col).Asc()
	default:
		order = goqu.I(col).Desc()
	}

	b.builder = b.builder.Order(order)

	return b
}

func (b *GetMessagesBuilder) WhereSearch(search string) *GetMessagesBuilder {
	col := goqu.C("raw")

	// for now escape postgres search patterns
	search = strings.ReplaceAll(search, "_", "\\_")
	search = strings.ReplaceAll(search, "%", "\\%")

	not := search[0] == '!'
	if not {
		search = search[1:]
	}

	target := "%" + search + "%"

	expressions := []exp.Expression{}

	if not {
		expressions = append(expressions, col.NotLike(target))
	} else {
		expressions = append(expressions, col.Like(target))
	}

	where := goqu.Or(expressions...)

	b.builder = b.builder.Where(where)

	return b
}
