package rpcr

func responseKey(correlationID, endpoint string) string {
	return "rpcr:" + endpoint + ":" + correlationID
}
