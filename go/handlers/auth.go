package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"backend-app/go/models" 
  "backend-app/go/helpers"
	"math/rand"
  "strconv"
  "github.com/dgrijalva/jwt-go"
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

func Login() echo.HandlerFunc {
  return func (c echo.Context) error {
    form := new(models.Login)
		c.Bind(form)
		if err := c.Validate(form); err != nil {
			return err 
    }
    
    //user phone exist
    user := helpers.GetUserByPhone(form.Phone)
    if user.ID == 0 {
      return c.JSON(400, H{
        "message": "User not found",
      })
    }

    //password check
    if form.Password != user.Password {
      return c.JSON(400, H{
        "message": "Password doesn't match",
      })
    }

    //generate jwt token
    token, err := helpers.GenerateToken(user.Name, user.Phone, user.Role)
    if err != nil {
      panic(err)
    }

    return c.JSON(http.StatusOK, H{
      "data": token,
    })
  }
}

func UserAuth() echo.HandlerFunc {
	return func (c echo.Context) error {
		user := c.Get("user").(*jwt.Token)

    return c.JSON(http.StatusOK, H{
      "data": user.Claims,
    })
	}
}
