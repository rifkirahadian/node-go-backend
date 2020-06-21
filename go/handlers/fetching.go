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
      usdFetchData := helpers.ChangeIDRtoUSDCurrency(fetchData, usdValue)

      return c.JSON(http.StatusOK, H{
        "data": usdFetchData.Fetchs,
      })
    }

    return c.JSON(http.StatusOK, H{
      "data": fetchData,
    })
  }
}