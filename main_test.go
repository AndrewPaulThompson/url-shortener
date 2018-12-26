package main_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "url-shortener/router"
    "os"
)

var a router.App

func TestMain(m *testing.M) {
    a = router.App{}
    a.Initialize()
    code := m.Run()
    os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)
    return rr
}

func TestCreateUrl(t *testing.T) {
    payload := []byte(`{"longUrl":"https://www.google.com"}`)
    req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
