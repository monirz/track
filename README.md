# track
Track is a simple HTTP router for Go. Yet another router for Go? I created this for learning and experiment.  


### Install

`go get github.com/monirz/track`

### Usage

```go
func main() {

	r := track.New()

	r.Get("/users", users)
	r.Get("/users/profile", users)
	r.Post("/users/:id", users)
	r.Patch("/users/:id", users)
	r.Delete("/users/:id", users)

	http.ListenAndServe(":8080", r)
}

//create your http handler 
func users(w http.ResponseWriter, r *http.Request, params map[string]string) {
	//get the route parameter
	id := params["id"]
	fmt.Println(id)
	w.Write([]byte("users"))
}
```
