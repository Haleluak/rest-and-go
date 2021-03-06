package api

import (
	"github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  09-Apr-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func TestIndexHandler(t *testing.T) {
	convey.Convey("Given HTTP request to /", t, func() {
		routes := GetRoutes()

		req := httptest.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()

		convey.Convey("When request is handled by the Router", func() {
			routes.ServeHTTP(resp, req)
			convey.Convey("Response code should be 200", func() {

				convey.So(resp.Code, convey.ShouldEqual, 200)
			})
		})
	})
}
