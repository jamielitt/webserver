package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Go Web Server Initialising ...")
	fmt.Println("Registering / as a route")
	http.HandleFunc("/", get_write_headers)

	fmt.Println("Starting web server, listening on port 1111")
	err := http.ListenAndServe(":1111", nil)

	// Check for errors
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func get_write_headers(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request")
	fmt.Println("All Headers:")

	// Look over all key,value pairs sent in the header
	// range acts like a foreach
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
