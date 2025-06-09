package api

type GetHelloRequest struct {
	Id uint64 `json:"id"`
}

type GetHelloResponse struct {
	Text string
}
