# Static File Server

The fastest way to serve static files.

## Usage
```bash
go run main.go --port <http-port> --path <directory>
```

## Clients
Static File Server uses Minio as a backend (on docker environmet). So if you want to upload files to the server, you can use the following clients:

- https://docs.min.io/docs/java-client-quickstart-guide.html
- https://docs.min.io/docs/golang-client-quickstart-guide.html
- https://docs.min.io/docs/python-client-quickstart-guide.html
- https://docs.min.io/docs/javascript-client-quickstart-guide.html
- https://docs.min.io/docs/dotnet-client-quickstart-guide.html
- https://docs.min.io/docs/haskell-client-quickstart-guide.html