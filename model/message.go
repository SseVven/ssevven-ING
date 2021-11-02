package model

type RegistRes struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
type LoginRes struct {
	Code    int32  `json:"code"`
	Msg     string `json:"msg"`
	UserUrl string `json:"userUrl"`
}
