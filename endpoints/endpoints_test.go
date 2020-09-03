package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// helper to test a handler function
func testEndpoint(t *testing.T, handlerFuncName string, endpoint http.HandlerFunc, vars map[string]string, reqBody io.Reader, expectedStatus int, expectedBody string) {
	t.Helper()

	req, _ := http.NewRequest("", "", reqBody)
	if req.ContentLength > 0 {
		req.Header.Add("content-type", "application/json")
	}
	rr := httptest.NewRecorder()
	if vars != nil {
		req = mux.SetURLVars(req, vars)
	}
	endpoint.ServeHTTP(rr, req)

	if status := rr.Code; status != expectedStatus {
		t.Errorf("%v returned status %v, expected %v", handlerFuncName, status, expectedStatus)
	}

	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("%v returned body\n%v\nexpected\n%v", handlerFuncName, body, expectedBody)
	}

}

// calls json.Marshall and panics on error
func marshallJSONWithoutError(value interface{}) []byte {
	j, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return j
}

func TestNewEndpoints(t *testing.T) {
	t.Run("sets some routes", func(t *testing.T) {
		got := NewEndpoints(&MockRepository{}, mux.NewRouter())

		routeCount := 0
		got.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			routeCount++
			return nil
		})
		if routeCount == 0 {
			t.Errorf("NewEndpoints() set 0 routes, want more")
		}
	})
}

func Test_respondWithError(t *testing.T) {
	type args struct {
		w       *httptest.ResponseRecorder
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{"writes error object", args{w: httptest.NewRecorder(), code: 500, message: "internal server error"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithError(tt.args.w, tt.args.code, tt.args.message)
			if status := tt.args.w.Code; status != tt.args.code {
				t.Errorf("respondWithError() set code %v, want %v", status, tt.args.code)
			}
			if body := tt.args.w.Body.String(); body != fmt.Sprintf(`{"error":"%s"}`, tt.args.message) {
				t.Errorf("respondWithError() set body\n%v\nwant\n%v", body, fmt.Sprintf(`{"error":"%s"}`, tt.args.message))
			}
		})
	}
}

func Test_respondWithJSON(t *testing.T) {
	type args struct {
		w       *httptest.ResponseRecorder
		code    int
		payload interface{}
		want    string
	}
	tests := []struct {
		name string
		args args
	}{
		{"writes correct response", args{w: httptest.NewRecorder(), code: 200, payload: map[string]string{"hello": "world"}, want: `{"hello":"world"}`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithJSON(tt.args.w, tt.args.code, tt.args.payload)
			if status := tt.args.w.Code; status != tt.args.code {
				t.Errorf("respondWithJSON() set code %v, want %v", status, tt.args.code)
			}
			if body := tt.args.w.Body.String(); body != tt.args.want {
				t.Errorf("respondWithJSON() set body\n%v\nwant\n%v", body, tt.args.want)
			}
		})
	}
}
