package helpers

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "backend-app/go/models"
  "strconv"
  "time"
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

func ChangeIDRtoUSDCurrency(fetchData models.FetchCollection, usdValue float32) models.FetchCollection {
  result := models.FetchCollection{}
  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]

    value, _ := strconv.ParseFloat(fetch.Price, 32)
    floatPrice := float32(value)
    fetch.Price = fmt.Sprintf("%f", floatPrice * usdValue)

		result.Fetchs = append(result.Fetchs, fetch)
  }

  return result
}

func GetFetchByProvinceArea(fetchData models.FetchCollection, provinceArea string) models.FetchCollection {
  result := models.FetchCollection{}

  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]

    if fetch.AreaProvinsi == provinceArea {
      result.Fetchs = append(result.Fetchs, fetch)
    }
  }

  return result
}

func GetFetchByWeekNumber(fetchData models.FetchCollection, weekNumber string) models.FetchCollection {
  result := models.FetchCollection{}

  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]
    weekNumberSelected := GetWeekYearFromDate(fetch.TglParsed)

    if weekNumber == weekNumberSelected {
      result.Fetchs = append(result.Fetchs, fetch)
    }
  }

  return result
}

func GetProvinceArea(fetchData models.FetchCollection) map[string]models.FetchCollection {
  fetchDatas := make(map[string]models.FetchCollection)
  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]
    provinceArea := fetch.AreaProvinsi
    
    if _, ok := fetchDatas[provinceArea]; !ok {
      fetchDatas[provinceArea] = GetFetchByProvinceArea(fetchData, provinceArea)
    }
  }

  return fetchDatas
}



func GetWeekNumberGrouped(fetchData models.FetchCollection) map[string]models.FetchCollection {
  fetchDatas := make(map[string]models.FetchCollection)
  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]
    weekNumber := GetWeekYearFromDate(fetch.TglParsed)
    
    if _, ok := fetchDatas[weekNumber]; !ok {
      fetchDatas[weekNumber] = GetFetchByWeekNumber(fetchData, weekNumber)
    }
  }

  return fetchDatas
}

func GetWeekYearFromDate(str string) string {
  layout := "2006-01-02T15:04:05.000Z"
  t, _ := time.Parse(layout, str)

  year, week := t.ISOWeek()

  return strconv.Itoa(year) + " week " + strconv.Itoa(week)
}