package request

import "strings"

type Request struct {
	Method string
	URI    string
}

func New(reqLn string) Request {

	fs := strings.Fields(reqLn)
	//example req = "GET /home HTTP/1.1" -> ["GET", "/home", "HTTP/1.1"]

	return Request{
		Method: fs[0],
		URI:    fs[1],
	}
}
