package API

type JsonRespone struct {
	JMeta struct {
		     Status string `json:"status"`
		     Message string `json:"message"`
		     Errors interface{} `json:"errors"`
	     } `json:"meta"`
	JData interface{} `json:"data"`
}

func (r JsonRespone) Ok(data ...interface{}) *JsonRespone {
	if len(data) > 0 {
		r.Data(data[0])
	}
	
	r.JMeta.Status = string(STATUS.Ok)
	return &r
}

func (r JsonRespone) Err(err ...interface{}) *JsonRespone {
	if len(err) > 0 {
		r.Errors(err[0])
	}
	r.JMeta.Status = string(STATUS.Error)
	return &r
}

func (r JsonRespone) Errors(err interface{}) *JsonRespone {
	r.JMeta.Errors = err
	return &r
}

func (r JsonRespone) Data(data interface{}) *JsonRespone {
	r.JData = data
	return &r
}

func (r JsonRespone) Msg(msg string) *JsonRespone {
	r.JMeta.Message = msg
	return &r
}

func JSON(data ...interface{}) *JsonRespone  {
	res := JsonRespone{}
	res.Ok(data)
	
	return &res
}

