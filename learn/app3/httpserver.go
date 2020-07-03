package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloer)
	err := http.ListenAndServe(":6161", nil)
	if err != nil {
		log.Println("hata var", err)
		return
	}
}

func helloer(rw http.ResponseWriter, r *http.Request) {
	isim := r.URL.Query().Get("ad")
	rw.Write(([]byte("Merhaba " + isim)))

	//	http://localhost:6161/?ad=murat
}
