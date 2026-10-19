package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func TestContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(recorder)
	return ctx
}

func MakeRequest(
	c *gin.Context,
	params gin.Params,
	urlValues url.Values,
	method string,
	body *bytes.Buffer,
) {
	var req *http.Request
	var err error

	if body == nil {
		req, err = http.NewRequest(method, "/?"+urlValues.Encode(), nil)
	} else {
		req, err = http.NewRequest(method, "/?"+urlValues.Encode(), body)
	}

	if err != nil {
		panic(err)
	}

	c.Request = req
	c.Params = params
}
