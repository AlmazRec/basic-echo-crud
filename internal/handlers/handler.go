package handlers

type Handler struct {
	PostHandler PostHandler
}

func NewHandler(postHandler PostHandler) *Handler {
	return &Handler{PostHandler: postHandler}
}
