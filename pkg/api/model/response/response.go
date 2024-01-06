package response

type ResponseDto struct {
	Error error            `json:"error"`
	Data  []map[string]any `json:"data"`
}
