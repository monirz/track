package track

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *Router

func TestAdd(t *testing.T) {

	if router != nil {
		t.Errorf("expected %v got %v ", nil, router)
	}

	router = New()
	v := []string{"hello", "world"}

	r := httptest.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
	})

	testHandler(w, r)

	router.add(v, "GET", testHandler)

	// if n.Children["hello"].Children["world"].Value !=  {
	// 	t.Errorf("expected %v  got %v", "handler", n.Children["hello"].Children["world"].Value)
	// }

	v = []string{"api", "v1", "users", "posts"}

	router.add(v, "GET", testHandler)
	v = []string{"api2", "v2", ":id", "posts"}
	router.add(v, "GET", testHandler)

}

func TestSearch(t *testing.T) {

	v := []string{"hello", "world"}

	result := router.search(v)

	v = []string{"api", "v1", "users", "posts"}
	result = router.search(v)

	if result == nil {
		t.Errorf("expected %v got %v ", nil, result)
	}

	v = []string{"api2", "v2", "12345", "post"}
	result = router.search(v)

	if result != nil {
		t.Errorf("expected %v got %v ", nil, result)
	}

	v = []string{"api2", "v2", "12345", "posts"}
	result = router.search(v)

	if result == nil {
		t.Errorf("expected %v got %v ", nil, result)
	}

}

func BenchmarkMux(b *testing.B) {
	router := New()
	handler := func(w http.ResponseWriter, r *http.Request) {}
	router.Get("/book/:id", handler)

	request, _ := http.NewRequest("GET", "/book/anything", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, request)
	}
}

func BenchmarkAdd(b *testing.B) {
	// run the Fib function b.N times
	v := []string{"v1", ":id"}
	handler := func(w http.ResponseWriter, r *http.Request) {}

	for i := 0; i < b.N; i++ {
		router.add(v, "GET", http.HandlerFunc(handler))
	}
}
