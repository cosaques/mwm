package main

import (
	"log"
	"net/http"

	"github.com/cosaques/mwm/business"
)

func main() {
	df := &business.DataFeeder{}

	http.HandleFunc("/admin/upload", df.AdminHandler)

	log.Println("Starting the webserver on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
