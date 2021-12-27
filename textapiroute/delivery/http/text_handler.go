package http

import (
	"encoding/json"
	"github.com/afzalabbasi/demo-code-invozone/service"
	"github.com/afzalabbasi/demo-code-invozone/system/messages"
	"net/http"

	httpresponse "github.com/afzalabbasi/demo-code-invozone/network/response"
	"github.com/afzalabbasi/demo-code-invozone/textapiroute"
	"github.com/labstack/echo"
)

func NewTextHandler(pub *echo.Group) {

	pub.POST("/text", TextCounter)
}

func TextCounter(c echo.Context) error {

	req := textapiroute.InputRequest{}
	err := c.Bind(&req)
	if err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, messages.PleaseTryAgain, err.Error())
	}
	// validate input request body
	if err := c.Validate(req); err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, messages.PleaseTryAgain, err.Error())
	}
	data, err := service.WordCount(req.Text)
	if err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, messages.PleaseTryAgain, err.Error())
	}
	b, _ := json.Marshal(data)
	return httpresponse.CreateSuccessResponse(&c, http.StatusOK, "success", "success", b)

}
