package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"fmt"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(122)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h appHandler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNoPermission, 403, "Forbidden"},
	{errNotFound, 404, "Not Found"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper (t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}


func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response,
	expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode ||
		body != expectedMsg {
		t.Errorf("expect (%d, %s); " +
			"got (%d, %s)", expectedCode, expectedMsg, response.StatusCode, body)
	}
}