
# WebServ

A lightweight and user-friendly HTTP server written in Go that allows you to serve files from the current directory. With customizable ports and simple installation, WebServ is the perfect tool for quick and efficient file serving.

---

## Features

- **Simple to Use**: Serve files from the current directory with a single command.
- **Customizable Port**: Change the port easily using the `-port` flag.
- **Port Persistence**: Remembers your last used port for convenience.
- **Graceful Shutdown**: Close the server with `CTRL+C`.
- **Cross-Platform**: Works on Linux, macOS, and Windows.

---

## Installation

Install WebServ directly from GitHub using `go install`:

```bash
go install github.com/fernstedt/webserv@latest
```

> Ensure your `GOPATH/bin` is in your `PATH` to use `webserv` globally. If not, you can add it with:
> ```bash
> export PATH=$PATH:$(go env GOPATH)/bin
> ```

---

## Usage

### Start the Server

Run WebServ with the default port (8080):

```bash
webserv
```

### Specify a Custom Port

Set a custom port (e.g., 9090):

```bash
webserv -port=9090
```

The server will remember this port for future runs.

### Access the Server

Once running, you can access the server in your browser:

```
http://<your-ip>:<port>/
```

For example:

```
http://192.168.1.100:8080/
```

### Stop the Server

Press `CTRL+C` in the terminal to stop the server gracefully.

---

## Example Output

When running the server, you will see:

```
Serving HTTP on 192.168.1.100 port 8080 (http://192.168.1.100:8080/)

Run webserv with -port <port> to select any port.
(Close with CTRL+C)
```

---

## Development

### Clone the Repository

To contribute or modify the project, clone the repository:

```bash
git clone https://github.com/fernstedt/webserv.git
cd webserv
```

### Build the Binary

Build the project locally:

```bash
go build -o webserv
```

Run the binary:

```bash
./webserv
```

---

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for bug fixes, improvements, or new features.

---

## License

WebServ is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute this software.

---

## Contact

For questions or feedback, please contact [fernstedt](https://github.com/fernstedt).

---

### Happy Serving! 🚀
