package model

const SuccessCode = 0

const SuccessMsg = "ok"

type JsonResponse struct {
	ErrorCode   int         `json:"errcode"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

type BlockchainResponse struct {
	TxID     string `json:"tx_id"`
	Response string `json:"response"`
}
