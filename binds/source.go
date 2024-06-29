package binds

type CreateSourceBind struct {
	Name string `json:"name"`
}

type GetSourceMessagesBind struct {
	GetMessagesBind
	GetSourceBind
}

type GetSourceBind struct {
	Id int32 `param:"id"`
}

type CreateSourceMessageBind struct {
	GetSourceBind
}
