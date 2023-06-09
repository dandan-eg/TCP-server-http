package param

import "strings"

type Param struct {
	Method string
	URI    string
}

func New(req string) Param {
	//example req "GET /home HTTP/1.1"

	fs := strings.Fields(req)

	return Param{
		Method: fs[0],
		URI:    fs[1],
	}
}
