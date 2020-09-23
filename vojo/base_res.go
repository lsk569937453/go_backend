package vojo

const (
	NORMAL_RESPONSE_STATUS   = 0
	ERROR_RESPONSE_STATUS    = -1
	ERROR_STATUS_PARAM_WRONG = -2
)

type BaseRes struct {
	Rescode int         `json:"resCode"`
	Message interface{} `json:"message"`
}
