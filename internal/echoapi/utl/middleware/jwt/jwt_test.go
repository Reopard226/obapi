package jwt_test

import (
	"net/http"
	"net/http/httptest"
	"oceanbolt.com/obapi/internal/iam/iamclient"
	"testing"

	"oceanbolt.com/obapi/internal/echoapi/utl/middleware/jwt"
	"oceanbolt.com/obapi/internal/echoapi/utl/mock"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func echoHandler(mw ...echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()
	for _, v := range mw {
		e.Use(v)
	}
	e.GET("/hello", hwHandler)
	return e
}

func hwHandler(c echo.Context) error {
	return c.String(200, "Hello World")
}

func TestMWFunc(t *testing.T) {
	cases := []struct {
		name       string
		wantStatus int
		header     string
		signMethod string
	}{
		{
			name:       "Empty header",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Header not containing Bearer",
			header:     "notBearer",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Invalid header",
			header:     mock.HeaderInvalid(),
			wantStatus: http.StatusUnauthorized,
		},
		/*{
			name:       "Success",
			header:     mock.HeaderValid(),
			wantStatus: http.StatusOK,
		},*/
	}
	jwtMW := jwt.New(`MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA5pfKq7tMmsxrqWUWIUCL
ffoVRXwYRDtDsq71ntLYo1b75gdoQBjchFWTzX0AgSefpiY3Vu3abxhjgPjP0AU6
PQvHGQO7NMB2uthV/Y29NGY8JJJaYcrqxZEZNRwL8tQ4PybXfXP+AuJIQIIyNxey
FvOI8v3nyVtEEMpI2JrW+ExrCPMHG00CcIRjDVxof2rT1B5UvZpc6il0PcPydYJn
xbHz0A4vS5GDfsL74MlO4Hj4+ADb+8fjV4B+dbLP+W4GthIAWN2d70Ls1NjLJDRT
CejTG7n5H20D5uBoOkZuZL7N89a6idbg+K9AJlycrQpQdUwXOTzZbKjbmOz8q6Sh
2jhLIoBTH9c4q/YGYm7lc37LdENmNGjjX/6OlZ8JEMHJg5mRCJ+soslB3kS4SpYP
GQimw1PmipdrKKX374O3bdCzxirdo1+zxfO2luALGUV9RSzYeMYOp0PXeFV7ku+3
i2p5GbnySFfR2Zz39bFyFpCcwuS8AugRKfzMDbc313zY+isM86mdf9oYnGrMKXcY
PKUcJdGCHJ4gZDNKBe1KxisK9iLfUA0AQ7bIvuuDxxjxIGhdNIPPqeJyrXhPbuB0
QAF+SzUDK10tq+v2Zq39wYGI0vB6I9eb3Im+WeP8ztDDTiqirsmzxdvomKzabORZ
yNA6/ROkItfoSaMv8/QpnpsCAwEAAQ==`, "RS256", 60, iamclient.GetDefaultIamClient())
	ts := httptest.NewServer(echoHandler(jwtMW.MWFunc()))
	defer ts.Close()
	path := ts.URL + "/hello"
	client := &http.Client{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", path, nil)
			req.Header.Set("Authorization", tt.header)
			res, err := client.Do(req)
			if err != nil {
				t.Fatal("Cannot create http request")
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}
