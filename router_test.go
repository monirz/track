package track

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tests = []struct {
	ReqMethod string
	ReqPath   string
	Method    string
	Path      string
	Params    map[string]string
}{
	{
		"GET", "/users",
		"GET", "/users", nil,
	},

	{
		"GET", "/user/profile",
		"GET", "/user/profile", nil,
	},

	{
		"POST", "/user",
		"POST", "/user", nil,
	},

	{
		"POST", "/users/1234",
		"POST", "/users/:id", nil,
	},

	{
		"PATCH", "/users/1234",
		"PATCH", "/users/:id", nil,
	},

	{
		"DELETE", "/users/1234",
		"DELETE", "/users/:id", nil,
	},
	//with parameter
	{
		"POST", "/users/1234",
		"POST", "/users/:id", map[string]string{"id": "1234"},
	},
}

func TestTrack(t *testing.T) {
	for _, test := range tests {
		r := New()
		param := make(map[string]string)
		r.handle(test.Path, test.Method, Handle(func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			param = params
		}))

		req, err := http.NewRequest(test.ReqMethod, test.ReqPath, nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		fmt.Println(w.Code)
		if w.Code != http.StatusOK {
			t.Errorf("want %v, got %v ", http.StatusOK, w.Code)
		}

		for k, v := range test.Params {
			if v != param[k] {
				t.Errorf("want %v, got %v ", param[k], v)

			}
		}
	}
}
