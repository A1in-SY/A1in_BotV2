package model

type QQCallResp struct {
	Status  string      `json:"status"`
	RetCode int64       `json:"retcode"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Wording string      `json:"wording"`
	Echo    string      `json:"echo"`
}

const (
	QQCallRespStatusOK = "ok"
)
