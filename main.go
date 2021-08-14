package main

import (
	"fmt"
	"net/http"
)

func hello(
	writer http.ResponseWriter, // Response Writing Object
	req *http.Request) { // Pointer to Request object

	// print message
	fmt.Fprintf(writer, "<h1>Hello World!</h1>")

}

func headers(
	writer http.ResponseWriter, // Response Writing Object
	req *http.Request) { // Pointer to Request object)

	// Iterate over the headers in the request
	for name, headers := range req.Header {

		// Iterate over each subheader
		for _, h := range headers {

			// print message
			fmt.Fprintf(writer, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
