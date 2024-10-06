package goware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TestHttpRecords struct {
	name               string
	method             string
	path               string
	expectedStatusCode int
	expectedBody       string
}

var SimpleTestData = TestHttpRecords{
	name:               "SimpleTestData",
	method:             "GET",
	path:               "/",
	expectedStatusCode: http.StatusOK,
	expectedBody:       "Hello World",
}

func SetupTest(t *testing.T, input string) (GoWare, *http.Request, *httptest.ResponseRecorder, http.Handler) {
	sut := GoWare{}

	req, err := http.NewRequest(SimpleTestData.method, SimpleTestData.path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(input))
	})

	middleware := sut.Use(handler)

	return sut, req, rr, middleware
}
func TestMiddlewareWithBasicHttpRequest(t *testing.T) {
	sut, req, rr, middleware := SetupTest(t, "Hello World")
	middleware.ServeHTTP(rr, req)
	file, err := os.OpenFile("Info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	sut.SetupLogger(InfoLogger)

	if rr.Code != SimpleTestData.expectedStatusCode {
		t.Errorf("Expected %d, got %d", SimpleTestData.expectedStatusCode, rr.Code)
	}

	if body := rr.Body.String(); body != SimpleTestData.expectedBody {
		t.Errorf("Expected %s, got %s", SimpleTestData.expectedBody, body)
	}

}

func TestMiddlewareLogging(t *testing.T) {
	sut := GoWare{}

	file, err := os.OpenFile("Info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	sut.SetupLogger(InfoLogger)

	if sut.logger == nil {
		t.Errorf("Logger not set")
	}

}
