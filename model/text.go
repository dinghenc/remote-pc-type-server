package model

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	RetCode int    `json:"ret_code"`
	ErrInfo string `json:"err_info"`
}
