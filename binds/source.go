package binds

type CreateSourceBind struct {
	Name string `json:"name"`
}

type GetSourceMessagesBind struct {
	GetMessagesBind
	Id int32 `param:"id"`
}

type GetSourceBind struct {
	Id int32 `param:"id"`
}
