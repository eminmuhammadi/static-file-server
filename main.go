package main

import (
	"flag"
	"log"
	"path/filepath"

	app "github.com/eminmuhammadi/static-file-server/app"
)

var (
	path = flag.String("path", ".", "Path to directory to serve")
	port = flag.String("port", "8080", "Port to serve on")
)

func main() {
	flag.Parse()

	dirname, err := filepath.Abs(*path)
	if err != nil {
		log.Fatalf("Could not get absolute path to directory: %s: %s", dirname, err.Error())
	}

	log.Printf("Serving %s on port %s", dirname, *port)

	if err := app.Run(*port, dirname); err != nil {
		log.Fatalf("Could not run server: %s", err.Error())
	}
}
