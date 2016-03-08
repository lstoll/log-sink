package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleLogShuttle(w http.ResponseWriter, r *http.Request) {
	for {
		var buff = make([]byte, 100, 100)
		_, err := r.Body.Read(buff)
		if err != nil && err != io.EOF {
			log.Fatal("Error:", err)
		}
		fmt.Print(string(buff))
		if err == io.EOF {
			break
		}
	}
}

func main() {
	http.HandleFunc("/", handleLogShuttle)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
