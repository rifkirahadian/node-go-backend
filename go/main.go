package main

import (
  "gopkg.in/go-playground/validator.v9"
  "github.com/labstack/echo"
  "backend-app/go/handlers"
  "backend-app/go/helpers"
  "github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main()  {
  e := echo.New()

  e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helpers.ValidationResponse

  apiRoutes := e.Group("/api")
  
  //no auth routes
  apiRoutes.POST("/register", handlers.Register())
  apiRoutes.POST("/login", handlers.Login())

  //auth middleware
	authRoutes := apiRoutes.Group("/auth")
  authRoutes.Use(middleware.JWT([]byte("secret")))
  
  //auth routes
  authRoutes.GET("/user", handlers.UserAuth())
  authRoutes.GET("/fetching/fetch", handlers.FetchingFetch())

	e.Logger.Fatal(e.Start(":8000"))
}
