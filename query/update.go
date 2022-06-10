package query

type UpdateReq map[string]any

type UpdateParam interface {
	DoUpdate(UpdateReq)
}
