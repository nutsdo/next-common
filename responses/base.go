package responses

type ResponseJson struct {
	StatusCode int
	Msg interface{}
	Data interface{}
}

type ErrResponseJson struct {
	StatusCode int
	SubCode int
	ErrorMsg interface{}
}


func SuccessResponse(code int,msg string,data interface{}) (sr *ResponseJson) {
	sr = &ResponseJson{StatusCode: code, Data: data, Msg: msg}
	return
}

func ErrResponse(code,subcode int,errmsg string) (er *ErrResponseJson) {
	er = &ErrResponseJson{StatusCode: code, SubCode: subcode, ErrorMsg: errmsg}
	return
}