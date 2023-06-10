package main

import (
	"TCP-server-http/request"
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer func() { log.Println(li.Close()) }()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	read(conn)
	write(conn)
}

const requestLn = 0

func read(conn net.Conn) {
	ln := 0
	sc := bufio.NewScanner(conn)

	for sc.Scan() {
		txt := sc.Text()

		if ln == requestLn {
			p := request.New(txt)
			log.Printf("[%s] %s", p.URI, p.Method)
		}

		if txt == "" {
			break
		}

		ln++
	}

}

func write(conn net.Conn) {
	body := fmt.Sprintf(
		`
<!DOCTYPE html>
<html lang="en">
	<body>
		<h1>%s</h1>
	</body>
</html>`, "Hello world")

	//headers
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type : text/html\r\n")

	fmt.Fprintf(conn, "\r\n")

	//body
	fmt.Fprintf(conn, body)
}
