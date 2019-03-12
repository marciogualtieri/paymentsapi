package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func sendRequest(handler http.Handler, request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	return response
}

/*
GetResource sends a GET request to the resource path.
*/
func GetResource(handler http.Handler, path string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		panic(err)
	}
	return sendRequest(handler, request)
}

/*
PostResource sends a POST request to the resource path.
*/
func PostResource(handler http.Handler, path string, body string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("POST", path, bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}
	return sendRequest(handler, request)
}

/*
PutResource sends a PUT request to the resource path.
*/
func PutResource(handler http.Handler, path string, body string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("PUT", path, bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}
	return sendRequest(handler, request)
}

/*
DeleteResource sends a DELETE request to the resource path.
*/
func DeleteResource(handler http.Handler, path string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		panic(err)
	}
	return sendRequest(handler, request)
}
