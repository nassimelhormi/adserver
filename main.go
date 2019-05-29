package main

import (
	"flag"
	"log"
	"net/http"
)

func parseFlags() (filename string, port string) {

	flag.StringVar(&filename, "file", "datas/campaigns.json", "The data file. By default we use our datas.")
	flag.StringVar(&port, "port", "4242", "The server port. By default it's 8080.")
	flag.Parse()

	return filename, port
}

func main() {

	filename, port := parseFlags()
	log.Println(filename, port)
	log.Println("<--- Adserver Started ! --->")
	http.Handle("/ping", handlers.adServerHandler(jsonData))
	http.ListenAndServe(":"+port, nil)
}
