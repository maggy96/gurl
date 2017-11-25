package main

import (
  "testing"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "net/http"
  "net/http/httptest"
  . "github.com/smartystreets/goconvey/convey"
)

func TestUrlShortener(t *testing.T) {
  r := setupRouter()
  db, _ = gorm.Open("sqlite3", "./test.db")
  defer db.Close()
  db.AutoMigrate(&Url{})

  Convey("Shortening a link returns json", t, func() {
    req, _ := http.NewRequest("GET", "/v1/shorten/http://goerlitz.tech", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    So(resp.Code, ShouldEqual, http.StatusOK)
    So(resp.Body.String(), ShouldEqual, "{\"id\":1,\"payload\":\"http://goerlitz.tech\"}")
  })

  Convey("Opening that links sends a redirect", t, func() {
    req, _ := http.NewRequest("GET", "/v1/resolve/1", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    So(resp.Code, ShouldEqual, http.StatusTemporaryRedirect)
    So(resp.Body.String(), ShouldEqual, "<a href=\"http://goerlitz.tech\">Temporary Redirect</a>.\n\n")
  })
  db.DropTable(&Url{})
}
