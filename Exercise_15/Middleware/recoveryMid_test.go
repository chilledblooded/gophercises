package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"strings"
	"testing"

	handler "github.com/chilledblooded/gophercises/Exercise_15/Handlers"
)

func GetTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test entered test handler, this should not happen")
	}
	return http.HandlerFunc(fn)
}
func TestRecoveryMid(t *testing.T) {
	handler := http.HandlerFunc(handler.PanicDemo)
	executeRequest("Get", "/panic", RecoveryMid(handler))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}

func TestErrLinks(t *testing.T) {
	stack := debug.Stack()
	result := errLinks(string(stack))
	if !strings.Contains(result, "<a href=") {
		t.Error("Expected result is Error links, actual result is nil")
	}
}
