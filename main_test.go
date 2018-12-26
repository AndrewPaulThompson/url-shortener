package main_test

import (
    "bytes"
    "fmt"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "time"
    "url-shortener/router"

    "github.com/speps/go-hashids"
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
    hd          := hashids.NewData()
    h, _        := hashids.NewWithData(hd)
    now         := time.Now()
    id, _       := h.Encode([]int{int(now.Unix())})

    url := fmt.Sprintf(`{"longUrl":"https://www.%s.com"}`, id)
    payload := []byte(url)
    req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))

    // First attempt at creating
    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
