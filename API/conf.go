package API

type APIStatus string

var (
	STATUS = struct{
		Ok APIStatus
		Error APIStatus
	}{
		Ok: APIStatus("ok"),
		Error: APIStatus("error"),
	}
)