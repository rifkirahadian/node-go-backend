package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"backend-app/go/models" 
	"backend-app/go/helpers" 
	"math/rand"
	"strconv"
)

type H map[string]interface{}

func Register() echo.HandlerFunc  {
	return func (c echo.Context) error {
		u := new(models.User)
		c.Bind(u)
		if err := c.Validate(u); err != nil {
			return err 
		}

		password := ""
		//user name exist check
		currentPassword := helpers.UserNameExistCheck(u.Name)
		if currentPassword == "" {
      password = strconv.Itoa(rand.Intn(10000))

      //create user
      createUser := helpers.CreateUser(u.Name, u.Phone, u.Role, password)
      if createUser == false {
        return c.JSON(400, H{
          "message": "Phone number has been used by another user",
        })
      }
		}else{
      password = currentPassword
    }

		return c.JSON(http.StatusOK, H{
			"data": password,
		})
	}
}