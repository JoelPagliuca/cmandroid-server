package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	sw "github.com/JoelPagliuca/cmandroid-server/go"
)

// VERSION program version
var VERSION string

// GITCOMMIT commit hash that we built from
var GITCOMMIT string

func usage() {
	fmt.Println("cmandroid server")
	fmt.Println("version:", VERSION, "commit:", GITCOMMIT)
	fmt.Println("")
	flag.PrintDefaults()
}

func main() {
	helpFlag := flag.Bool("h", false, "Display this help text")
	port := flag.Int("p", 8080, "Port to run the server on")
	flag.Parse()

	if *helpFlag {
		usage()
		os.Exit(0)
	}

	log.Printf("Server started on port " + fmt.Sprint(*port))

	router := sw.NewRouter()

	portString := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(portString, router))
}
