package main

import (
  "fmt"
  "testing"
  "net/http"
  "net/http/httptest"
  . "github.com/smartystreets/goconvey/convey"
)

func TestUsersResource(t *testing.T) {
  r := setupRouter()
  setupDB()

  Convey("Shortening a link returns json", t, func() {
    req, _ := http.NewRequest("GET", "/v1/shorten/http://goerlitz.tech", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    fmt.Println(resp.Body.String())
    So(resp.Code, ShouldEqual, http.StatusOK)
    So(resp.Body.String(), ShouldEqual, "{\"id\":0,\"payload\":\"http://goerlitz.tech\"}")
  })

  Convey("Opening that links sends a redirect", t, func() {
    req, _ := http.NewRequest("GET", "/v1/resolve/1", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    fmt.Println(resp.Body.String())
    So(resp.Code, ShouldEqual, http.StatusTemporaryRedirect)
  })
}
