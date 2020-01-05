package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
