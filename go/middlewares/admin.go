package middlewares

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

type H map[string]interface{}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)

    if claims["role"] != "admin" {
      return c.JSON(403, H{
        "message": "Forbidden access",
      })
    }
    return next(c)
  }
}