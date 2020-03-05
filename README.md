# track 
[![Coverage](https://gocover.io/_badge/github.com/monirz/track)](https://gocover.io/github.com/monirz/track) [![Actions Status](https://github.com/monirz/track/workflows/build/badge.svg)](https://github.com/monirz/gotri/actions) 

Track is a fast and lightweight HTTP router for Go built using Trie/Prefix data structure, which doesn't break the standard net/http handler.  


### Install

`go get -u github.com/monirz/track`

### Usage

```go
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

	http.ListenAndServe(":8090", router)
}

```

### Using middleware

```go
func main() {

	r := track.New()

	r.Use(track.CORSMethodMiddleware(r))
	r.Get("/users", userHandler)
	r.Options("/users", userHandler)

	http.ListenAndServe(":8090", r)

}
```

### Test 

```
$ go test -v . 
```
### Benchmark 

```
$ go test -bench=. 
```