package rpcr

type Input struct {
	Endpoint      string `json:"endpoint"`
	Input         string `json:"input"`
	CorrelationID string `json:"correlation_id"`
}
