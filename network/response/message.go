package response

import (
	"fmt"
	"github.com/labstack/echo"
)

func CreateBadResponse(c *echo.Context, requestCode int, message string, subMessage string) error {
	localC := *c
	response := fmt.Sprintf("{\"data\":{},\"message\":%q,\"submessage\":%q}", message, subMessage)
	return localC.JSONBlob(requestCode, []byte(response))
}

func CreateSuccessResponse(c *echo.Context, requestCode int, message string, subMessage string, data []byte) error {

	localC := *c
	response := fmt.Sprintf("{\"data\":%s,\"message\":%q,\"submessage\":%q}", data, message, subMessage)
	return localC.JSONBlob(requestCode, []byte(response))
}

func CreateRawResponse(c *echo.Context, requestCode int, response []byte) error {
	localC := *c
	return localC.JSONBlob(requestCode, response)
}
