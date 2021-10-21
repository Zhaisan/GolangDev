package main

import (
	"fmt"
	"log"
	"net/http"
)


//ResponseWriter сам по себе является интерфейсом, обеспечивающим доступ к методам, необходимым для возврата ответа клиенту
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}