package main

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
	"os"

	_textHttpDelivery "github.com/afzalabbasi/demo-code-invozone/textapiroute/delivery/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {

	// check if required env variables are provided
	hasValidEnvVariables()
	//server setup
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	pub := e.Group("/v1")
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST},
	}))

	_textHttpDelivery.NewTextHandler(pub)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
func hasValidEnvVariables() {
	if os.Getenv("PORT") == "" {
		panic(errors.New("Please Provide Valid PORT"))
	}

}
