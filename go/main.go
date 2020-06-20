package main

import (
  "gopkg.in/go-playground/validator.v9"
  "github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main()  {
	e := echo.New()

	e.Logger.Fatal(e.Start(":8000"))
}
