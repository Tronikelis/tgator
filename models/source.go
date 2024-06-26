package models

type SourceModel struct {
	ID   int32  `db:"id" goqu:"defaultifempty"`
	Name string `db:"name"`
}
