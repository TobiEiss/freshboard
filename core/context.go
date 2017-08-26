package core

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CContext struct {
	echo.Context
	Config Config
}

// JSONF returns data as formatted JSON
// {"status": <http-status>, "data": <data>}
func (context *CContext) JSONF(code int, data interface{}, msg string) error {
	payload := map[string]interface{}{"status": strconv.Itoa(code), "msg": msg}
	if data != nil {
		payload["data"] = data
	}
	return context.JSON(code, payload)
}

// JSONFFail wraps JSONF and logs
func (context *CContext) JSONFFail(status int, err error) error {
	return context.JSONF(status, nil, err.Error())
}

// JSONFOK wraps JSONF with msg: OK and status: 200
func (context *CContext) JSONFOK(data interface{}) error {
	return context.JSONF(http.StatusOK, data, "OK")
}
