package api

type GetHelloRequest struct {
}

type GetHelloResponse struct {
	Text1 string `json:"text1"`
	Text2 string `json:"text2"`
}
