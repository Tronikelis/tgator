package models

type MessageModel struct {
	ID        int64  `db:"id" goqu:"defaultifempty"`
	CreatedAt string `db:"created_at" goqu:"defaultifempty"`

	Raw      string `db:"raw"`
	SourceId int32  `db:"source_id"`
}
