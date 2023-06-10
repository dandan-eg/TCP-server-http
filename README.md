# TCP Server HTTP

This is a simple TCP server that handles HTTP requests and sends back a response. It listens on port 8080 and responds
with a basic HTML page displaying "Hello world".

## Purpose

The purpose of this project is to provide a practical implementation of a TCP server handling HTTP requests. By working
with this code, I aim to deepen my understanding of the TCP protocol and the HTTP protocol.

Throughout this project, I will explore various aspects of TCP, such as establishing connections, accepting and managing
multiple connections concurrently, and reading and writing data over TCP connections. Additionally, I will gain insights
into the basics of the HTTP protocol, including parsing HTTP requests and constructing HTTP responses.

## Usage

To use this server, follow the steps below:


1. Clone the repository:
   ```shell
   git clone https://github.com/dandan-eg/TCP-server-http.git
   ```
2. Navigate to the project directory:
   ```shell
   cd TCP-server-http
   ```
3. Build the project:
   ```shell
   go build
   ```
4. Run the executable:
   ```shell
   ./TCP-server-http
   ```

The server will start listening on port 8080. You can access it by opening a web browser and
entering `http://localhost:8080` in the address bar. You should see the "Hello world" message displayed on the page.

## Dependencies

This server use the following Go packages:

- `net`: Provides basic networking operations, including TCP/IP socket communication.

The package is part of the Go standard library and is available by default.

## Server Implementation

The server is implemented in the `main` function. It follows the typical TCP server pattern of accepting connections and
spawning a goroutine to handle each connection.

### `main` Function

The `main` function is responsible for setting up the server and accepting connections. Here's an overview of its
functionality:

1. It listens for incoming connections on port 8080 using the `net.Listen` function.
2. If an error occurs during listening, it logs the error and exits.
3. It defers the closure of the listener using `li.Close()` to ensure proper cleanup when the server shuts down.
4. It enters an infinite loop, accepting connections using `li.Accept()`.
5. For each accepted connection, it spawns a goroutine to handle the connection by calling the `handle` function.

### `handle` Function

The `handle` function is responsible for handling an individual connection. Here's an overview of its functionality:


1. It defers the closure of the connection using `conn.Close()` to ensure the connection is closed when the function
   returns.
2. It calls the `read` function to read the incoming request from the client.
3. It calls the `write` function to write the response back to the client.

### `read` Function

The `read` function reads the incoming request from the client. Here's an overview of its functionality:

1. It uses a scanner (`bufio.Scanner`) to read the request line by line.
2. It checks if the line number is equal to `requestLn` (which is currently set to 0). If so, it creates a `param`
   object using the line and logs the URI and method of the request.
3. It checks if the line number is equal to `requestLn` (which is currently set to 0). If so, it creates a `param`
   object using the line and logs the URI and method of the request.



### `write` Function

The `write` function writes the response back to the client. Here's an overview of its functionality:

1. It creates the HTML body of the response with a simple "Hello world" message.
2. It writes the HTTP response headers to the client, including the status code, content length, and content type.
3. It sends a blank line (`\r\n`) to indicate the end of the headers.
4. It writes the HTML body to the client.

## Conclusion

In conclusion, this project serves as a personal learning endeavor aimed at gaining a better understanding of TCP and
HTTP protocols. It provides a TCP server that handles HTTP requests and demonstrates how routing and template rendering
can be implemented.

It is important to note that this project is not intended for actual usage in production
environments.