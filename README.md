# track
Track is a fast and lightweight HTTP router for Go built using Trie/Prefix data structure, which doesn't break the standard net/http handler.  


### Install

`go get -u github.com/monirz/track`

### Usage

```go
func main() {

	router := track.New()

	router.Get("/users", users)
	router.Get("/users/profile", users)
	router.Post("/users/:id", users)
	router.Patch("/users/:id", users)
	router.Delete("/users/:id", users)

	http.ListenAndServe(":8090", router)
}

//create your http handler 
func users(w http.ResponseWriter, r *http.Request) {
	//get the route parameter
	id := r.Context().Value("id")
	log.Println("user id", id)
	fmt.Fprintf(w, "user id: %v", id)
}
```

### Test 

```
$ go test -v . 
```