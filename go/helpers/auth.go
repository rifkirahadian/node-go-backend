package helpers

import (
	"time"

	"backend-app/go/configs"
	"backend-app/go/models"

	"github.com/dgrijalva/jwt-go"
)

func UserNameExistCheck(name string) string {
	user := new(models.User)

	db := configs.InitGormDB()
	if db.First(&user, "name=?", name).RecordNotFound() {
		return ""
	}

	return user.Password
}

func CreateUser(name string, phone string, role string, password string) bool {
	user := models.User{Name: name, Phone: phone, Role: role, Password: password}
	db := configs.InitGormDB()

	if err := db.Create(&user).Error; err != nil {
		return false
	}

	return true
}

func GetUserByPhone(phone string) models.User {
	db := configs.InitGormDB()
	var user models.User
	db.First(&user, "phone=?", phone)

	return user
}

func GenerateToken(name string, phone string, role string) (string, error) {
	//create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["phone"] = phone
	claims["name"] = name
	claims["role"] = role
	claims["timestamp"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//generate encode token
	t, err := token.SignedString([]byte("secret"))

	return string(t), err
}
