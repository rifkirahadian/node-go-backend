package main

import (
	"backend-app/go/handlers"
	"backend-app/go/helpers"
	"backend-app/go/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helpers.ValidationResponse

	apiRoutes := e.Group("/api")

	//auth middleware
	authRoutes := apiRoutes.Group("")
	authRoutes.Use(middleware.JWT([]byte("secret")))

	//admin middleware
	adminRoutes := authRoutes.Group("/admin")
	adminRoutes.Use(middlewares.AdminMiddleware)

	//no auth routes
	apiRoutes.POST("/register", handlers.Register())
	apiRoutes.POST("/login", handlers.Login())

	//auth routes
	authRoutes.GET("/user", handlers.UserAuth())
	authRoutes.GET("/fetching/fetch", handlers.FetchingFetch())

	adminRoutes.GET("/fetching/aggregate", handlers.FetchingAggregate())

	e.Logger.Fatal(e.Start(":8000"))
}
