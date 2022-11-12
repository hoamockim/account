package errors

import "net/http"

type ClientInfo struct {
	IP     string
	Method string
	Uri    string
}

type ErrResponse struct {
}

func HandlerErr(requestId string, request *http.Request, errCode ErrorCode, clientInfo *ClientInfo) (code int, res *ErrResponse) {
	//TODO: log,
	return
}
