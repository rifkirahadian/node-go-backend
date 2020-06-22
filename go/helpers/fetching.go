package helpers

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "backend-app/go/models"
  "strconv"
  "time"
  "reflect"
  "sort"
)

type MeanMedian struct {
	numbers []float32
}

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

func GetProvinceArea(fetchData models.FetchCollection) map[string]map[string]float32 {
  fetchDatas := make(map[string]map[string]float32)
  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]
    provinceArea := fetch.AreaProvinsi
    
    if _, ok := fetchDatas[provinceArea]; !ok {
      fetchDataGrouped := GetFetchByProvinceArea(fetchData, provinceArea)
      fetchDatas[provinceArea] = GetAggregatesData(fetchDataGrouped)
    }
  }

  return fetchDatas
}

func GetWeekNumberGrouped(fetchData models.FetchCollection) map[string]map[string]float32 {
  fetchDatas := make(map[string]map[string]float32)
  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]
    weekNumber := GetWeekYearFromDate(fetch.TglParsed)
    
    if _, ok := fetchDatas[weekNumber]; !ok {
      fetchDataGrouped := GetFetchByWeekNumber(fetchData, weekNumber)
      fetchDatas[weekNumber] = GetAggregatesData(fetchDataGrouped)
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

func GetAggregatesData(fetchData models.FetchCollection) map[string]float32 {
  aggregate := make(map[string]float32)

  var min float32
  var max float32 = 0
  var sum float32
  uniquePrices := MeanMedian{}

  for i := 0; i < len(fetchData.Fetchs); i++ {
    fetch := fetchData.Fetchs[i]
    
    value, _ := strconv.ParseFloat(fetch.Price, 32)
    floatPrice := float32(value)

    if min == 0 {
      min = floatPrice
    }else if floatPrice < min {
      min = floatPrice
    }

    if floatPrice > max {
      max = floatPrice
    }

    sum += floatPrice

    if fetch.Price != "" {
      if !ItemExists(uniquePrices.numbers, floatPrice) {
        uniquePrices.numbers = append(uniquePrices.numbers, floatPrice)
      }
    }
  }

  aggregate["min"] = min
  aggregate["max"] = max
  aggregate["avg"] = sum/float32(len(fetchData.Fetchs))
  aggregate["median"] = GetMedianFromArray(SortFloat32Array(uniquePrices))

  return aggregate
}

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func SortFloat32Array(float32Values MeanMedian) MeanMedian {
  float32AsFloat64Values := make([]float64, len(float32Values.numbers))

  for i, val := range float32Values.numbers {
    float32AsFloat64Values[i] = float64(val)
  }
  sort.Float64s(float32AsFloat64Values)

  for i, val := range float32AsFloat64Values {
    float32Values.numbers[i] = float32(val)
  }

  return float32Values
}

func GetMedianFromArray(datas MeanMedian) (float32) {
  if len(datas.numbers) > 1 {
    mNumber := len(datas.numbers) / 2
    return ((datas.numbers[mNumber-1] + datas.numbers[mNumber]) / 2)
  }else{
    return datas.numbers[0]
  }
}

