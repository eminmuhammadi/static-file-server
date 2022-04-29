# Static File Server

The fastest way to serve static files.

## Usage
```bash
go run main.go --port <http-port> --path <directory>
```

## Production Environment
### Requirements
- docker
- docker-compose

### Exposed ports
- 80: static files http directory (production)
- 9000: Minio API endpoints
- 9001: Minio Console

### Start service 
```bash
bash docker-up.sh
```

### Configuration
By default you need to create static-files bucket. To do so, open minio console then create new bucket named as `static-files`.

### Clients
Static File Server uses Minio as a backend (on docker environmet). So if you want to upload files to the server, you can use the following clients:

- https://docs.min.io/docs/java-client-quickstart-guide.html
- https://docs.min.io/docs/golang-client-quickstart-guide.html
- https://docs.min.io/docs/python-client-quickstart-guide.html
- https://docs.min.io/docs/javascript-client-quickstart-guide.html
- https://docs.min.io/docs/dotnet-client-quickstart-guide.html
- https://docs.min.io/docs/haskell-client-quickstart-guide.html
