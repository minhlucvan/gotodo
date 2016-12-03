package API

import "fmt"

type JsonRespone struct {
	JMeta struct {
		     Status string `json:"status"`
		     Message string `json:"message"`
		     Errors interface{} `json:"errors"`
	     } `json:"meta"`
	JData interface{} `json:"data"`
}

func (r JsonRespone) Ok() *JsonRespone {
	r.JMeta.Status = string(STATUS.Ok)
	return &r
}

func (r JsonRespone) Err() *JsonRespone {
	r.JMeta.Status = string(STATUS.Error)
	return &r
}

func (r JsonRespone) Throw(err interface{}) *JsonRespone {
	r.JMeta.Errors = err
	return &r
}

func (r JsonRespone) Data(data interface{}) *JsonRespone {
	r.JData = data
	fmt.Println(r)
	return &r
}

func (r JsonRespone) Msg(msg string) *JsonRespone {
	r.JMeta.Message = msg
	fmt.Println(r)
	return &r
}

func JSON(data ...interface{}) *JsonRespone  {
	res := JsonRespone{}
	res.Ok()
	
	return &res
}

