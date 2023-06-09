package param

import "strings"

type Param struct {
	Method string
	URI    string
}

func New(req string) Param {

	fs := strings.Fields(req)
	//example req = "GET /home HTTP/1.1" -> ["GET", "/home", "HTTP/1.1"]

	return Param{
		Method: fs[0],
		URI:    fs[1],
	}
}
