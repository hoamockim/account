package errors

import "net/http"

type ClientInfo struct {
	IP     string
	Method string
	Uri    string
}

type ErrResponse struct {
	Message string `json: "message"`
}

func HandlerErr(requestId string, request *http.Request, errCode ErrorCode, clientInfo *ClientInfo) (res *ErrResponse) {
	//TODO: log,
	res = &ErrResponse{
		Message: string(errCode),
	}
	return
}
