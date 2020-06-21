package helpers

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "backend-app/go/models"
  "strconv"
)

func GetUSDValue() float32{
	response, err := http.Get("https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=115c2a56202d33c99dc0")
  if err != nil {
    fmt.Printf("The HTTP request failed with error %s\n", err)
    return 0
  } else {
      data, _ := ioutil.ReadAll(response.Body)
      var usdValue models.USDValue
      json.Unmarshal([]byte(data), &usdValue)

      return usdValue.Value
  }
}

func ChangeIDRtoUSDCurrency(fetchData []models.Fetch, usdValue float32) models.FetchCollection {
  result := models.FetchCollection{}
  for i := 0; i < len(fetchData); i++ {
    fetch := fetchData[i]

    value, _ := strconv.ParseFloat(fetch.Price, 32)
    floatPrice := float32(value)
    fetch.Price = fmt.Sprintf("%f", floatPrice * usdValue)

		result.Fetchs = append(result.Fetchs, fetch)
  }

  return result
}