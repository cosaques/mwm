package main

import (
	"log"
	"net/http"

	"github.com/cosaques/mwm/business"
)

func main() {
	df := business.NewDataFeeder()

	http.HandleFunc("/admin/upload", df.AdminHandler)

	http.HandleFunc("/api/departments/", df.ApiHandler)

	log.Println("Starting the webserver on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
