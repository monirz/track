package main

import (
	"fmt"
	"net/http"

	"github.com/monirz/track"
)

func main() {

	router := track.New()

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello users!")
	})

	router.Use(track.CORSMethodMiddleware(router))

	// router.X(fooHandler)
	router.Get("/foo", fooHandler)
	// router.Options("/foo", fooHandler)

	http.ListenAndServe(":8090", router)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}
