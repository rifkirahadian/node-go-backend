package handlers

import (
  "backend-app/go/helpers"
  "github.com/labstack/echo"
  "net/http"
)

func FetchingFetch() echo.HandlerFunc {
  return func (c echo.Context) error {
    fetchData := helpers.GetJsonFetchData("fetching.json")
    currency := c.QueryParam("currency")
    if currency == "USD" {
      usdValue := helpers.GetUSDValue()
      fetchData = helpers.ChangeIDRtoUSDCurrency(fetchData, usdValue)
    }

    return c.JSON(http.StatusOK, H{
      "data": fetchData.Fetchs,
    })
  }
}

func FetchingAggregate() echo.HandlerFunc {
  return func (c echo.Context) error {
    types := c.QueryParam("type")
    fetchData := helpers.GetJsonFetchData("aggregate.json")

    if types == "provinceArea" {
      provinceAreaGrouped := helpers.GetProvinceArea(fetchData)

      return c.JSON(http.StatusOK, H{
        "data": provinceAreaGrouped,
      })
    } else if types == "weeklyDate" {
      weekNumberGrouped := helpers.GetWeekNumberGrouped(fetchData)

      return c.JSON(http.StatusOK, H{
        "data": weekNumberGrouped,
      })
    }
    
    return c.JSON(http.StatusOK, H{
      "data": fetchData,
    })
  }
}
