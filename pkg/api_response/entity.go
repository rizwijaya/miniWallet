package api

const (
	HttpStatusSuccess = "success"
	HttpStatusFailed  = "fail"
	HttpStatusError   = "error"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}
