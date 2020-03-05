package track

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	tests = []struct {
		ReqMethod string
		ReqPath   string
		Method    string
		Path      string
		Params    map[string]string
		handler   http.HandlerFunc
	}{
		{
			"GET", "/users",
			"GET", "/users", nil,
			testHandler,
		},

		{
			"GET", "/user/profile",
			"GET", "/user/profile", nil,
			testHandler,
		},

		{
			"POST", "/user",
			"POST", "/user", nil,
			testHandler,
		},

		{
			"POST", "/users/1234",
			"POST", "/users/:id", nil,
			testHandler,
		},

		{
			"PATCH", "/users/1234",
			"PATCH", "/users/:id", nil,
			testHandler,
		},

		{
			"DELETE", "/users/1234",
			"DELETE", "/users/:id", nil,
			testHandler,
		},
		{
			"POST", "/users/1234",
			"POST", "/users/:id", map[string]string{"id": "1234"},
			testHandler,
		},
	}
)

func TestTrack(t *testing.T) {

	r := New()

	for _, route := range tests {

		switch route.Method {
		case "GET":
			r.Get(route.Path, route.handler)
		case "POST":
			r.Post(route.Path, route.handler)
		case "PUT":
			r.Put(route.Path, route.handler)
		case "PATCH":
			r.Patch(route.Path, route.handler)
		case "DELETE":
			r.Delete(route.Path, route.handler)
		}
		req := httptest.NewRequest(route.ReqMethod, route.ReqPath, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("want %v, got %v ", http.StatusOK, w.Code)
		}

		resp := w.Result()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if string(respBody) != "hello" {
			t.Errorf("Expected %v got %v", "hello", string(respBody))
		}

	}
}

func TestMux(t *testing.T) {

}

func TestHTTPMethods(t *testing.T) {
	router := New()

	//test HTTP HEAD
	head := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Topic", "1")
		w.WriteHeader(200)
	}

	router.Head("/topic", head)
	testServer := httptest.NewServer(router)

	defer testServer.Close()

	resp, err := http.Head(testServer.URL + "/topic")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("head request failed expected %v got %v", 200, resp.StatusCode)
	}
	if resp.Header.Get("X-Topic") == "" {
		t.Errorf("expected %v got %v", "X-Topic", resp.Header.Get("X-Topic"))
	}

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	if _, body := testRequest(t, testServer, "GET", "/users", nil); body != "hello" {
		t.Fatalf(body)
	}

	if _, body := testRequest(t, testServer, "POST", "/users", nil); body != "hello" {
		t.Fatalf(body)
	}

	if resp, body := testRequest(t, testServer, "PATCH", "/users", nil); body != "hello" {

		if resp.StatusCode != 405 {
			t.Errorf("expected %v got %v", http.StatusMethodNotAllowed, resp.StatusCode)
		}

	}

}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}
