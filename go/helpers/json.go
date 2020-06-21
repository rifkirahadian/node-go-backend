package helpers

import (
  "backend-app/go/models"
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
)

func GetJsonFetchData(fileName string) []models.Fetch {
	jsonFile, err := os.Open("data/" + fileName)
  // if we os.Open returns an error then handle it
  if err != nil {
      fmt.Println(err)
  }

  // defer the closing of our jsonFile so that we can parse it later on
  defer jsonFile.Close()

  // read our opened jsonFile as a byte array.
  byteValue, _ := ioutil.ReadAll(jsonFile)

  // we initialize our Users array
  var fetchs []models.Fetch

  // we unmarshal our byteArray which contains our
  // jsonFile's content into 'users' which we defined above
  json.Unmarshal(byteValue, &fetchs)

  return fetchs
}