package common

func ServerErrResponse(err error) *Response {
	resp := new(Response)
	resp.Code = 500
	resp.Msg = err.Error()
	return resp
}

func SuccessResponse() *Response {
	resp := new(Response)
	resp.Code = 200
	resp.Msg = "成功"
	return resp
}
