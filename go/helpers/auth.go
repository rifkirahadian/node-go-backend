package helpers

import (
  "backend-app/go/models"
  "backend-app/go/configs"
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