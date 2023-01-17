package apimodels

type CommentInput struct {
	ThreadID string `json:"threadId"`
	Body     string `json:"body"`
}

type CommentEditInput struct {
	Body string `json:"body"`
}
