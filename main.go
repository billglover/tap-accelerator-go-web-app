package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Tanzu does Go too!"))
	})

	err := http.ListenAndServe(":8080", mux)
	return err
}