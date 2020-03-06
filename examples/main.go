package main

import (
	"fmt"
	"net/http"

	"github.com/monirz/track"
)

func main() {

	router := track.New()

	// router.Get("/users/id", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello users!")
	// })

	// router.Get("/posts/:id", postHandler)

	router.Get("/api2/v2/:id/posts", postHandler)
	// router.X(fooHandler)
	// router.Get("/foo", fooHandler)
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

func postHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Println("id: ", id)
}
