package binds

type CreateSourceBind struct {
	Ip string `json:"ip"`
}

type GetSourceMessagesBind struct {
	GetMessagesBind
	Id int32 `param:"id"`
}

type GetSourceBind struct {
	Id int32 `param:"id"`
}
