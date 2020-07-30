package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"backend-app/go/models"
)

func GetJsonFetchData(fileName string) models.FetchCollection {
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
	fetchs := models.FetchCollection{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &fetchs.Fetchs)

	return fetchs
}
