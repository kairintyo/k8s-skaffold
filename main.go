package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// log.Printf("TELEPRESENCE_ROOT: %v, TELEPRESENCE_MOUNTS: %v", os.Getenv("TELEPRESENCE_ROOT"), os.Getenv("TELEPRESENCE_MOUNTS"))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
	})

	http.HandleFunc("/happy", func(w http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadFile("/opt/app/assets/message.txt")
		if err != nil {
			log.Printf("Error: %v", err)
		}
		w.Write(msg)
	})

	log.Printf("Start server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
