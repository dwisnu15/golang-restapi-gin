package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//basically a helper function
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	return w
}

func TestGetAllItemsRequest(t *testing.T) {

}
