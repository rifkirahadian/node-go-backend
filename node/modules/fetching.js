const api = require('./api')
const servicesConfig = require('../configs/services')
const responser = require('./responser')
const NodeCache = require("node-cache")
const jsonQuery = require('json-query')
const moment = require('moment')

const myCache = new NodeCache();

exports.IDRtoUSD = async (res) => {
  try {
    let request = await api.send('GET', servicesConfig.currency.url, {})

    return request.IDR_USD
  } catch (error) {
    throw responser.errorResponse(res, error.message)
  }
}

exports.setUSDValueCache = (usdValue) => {
  myCache.set('usdValue', usdValue, servicesConfig.currency.cachingValueTime)
}

exports.getUSDValue = async(res) => {
  let usdValue = myCache.get('usdValue')
  if (!usdValue) {
    usdValue = await this.IDRtoUSD(res)
    this.setUSDValueCache(usdValue)
  }

  return usdValue
}

exports.handlePriceDataCurrency = async(data, currency, res) => {
  if (currency === 'USD') {
    let usdValue = await this.getUSDValue(res)

    data = data.map(item => {
      item.price = parseInt(item.price) * usdValue
      return item
    })
  }

  return data
}

exports.groupByAreaProvincies = (data) => {
  let areaProvincies = {}
  data.forEach(item => {
    if (item.area_provinsi) {
      if (!areaProvincies.hasOwnProperty(item.area_provinsi)) {
        areaProvincies[item.area_provinsi] = {
          items: []
        }
      }
      
      areaProvincies[item.area_provinsi].items.push(item)
    }
  })

  return areaProvincies
}

exports.groupByWeeklyDate = (data) => {
  let weeklyDates = {}
  data.forEach(item => {
    if (item.tgl_parsed) {
      let weeklyDate = moment(item.tgl_parsed).format('YYYY-WW')
      if (weeklyDate !== "Invalid date") {
        let parseWeeklyDate = weeklyDate.split('-')
        weeklyDate = `${parseWeeklyDate[0]} week ${parseWeeklyDate[1]}`

        if (!weeklyDates.hasOwnProperty(weeklyDate)) {
          weeklyDates[weeklyDate] = {
            items: []
          }
        }

        weeklyDates[weeklyDate].items.push(item)
      }
    }
  })

  return weeklyDates
}

exports.getMedian = (array) => {
  let len = array.length
  const arrSort = array.sort();
  const mid = Math.ceil(len / 2);

  return len % 2 == 0 ? (arrSort[mid] + arrSort[mid - 1]) / 2 : arrSort[mid - 1];
} 

exports.getAggregates = (data) => {
  let newData = {}
  for (const key in data) {
    const element = data[key]

    let [sum, min, max, uniqueValue] = [0,null,null, []]

    element.items.forEach(item => {
      let price = parseInt(item.price)
      if (price) {
        if ((price > max) || (max === null)) {
          max = price
        }

        if ((price < min) || (min === null)) {
          min = price
        }

        sum += price

        if (uniqueValue.indexOf(price) < 0) {
          uniqueValue.push(price)
        }
      }
    })

    let median = this.getMedian(uniqueValue)
    let avg = parseFloat((sum/element.items.length).toFixed(2))

    newData[key] = {min, max,median, avg}
  }

  return newData
}

exports.aggregateTypeValidate = (type, res) => {
  if (['provinceArea' , 'weeklyDate'].indexOf(type) < 0) {
    throw responser.errorResponse(res, 'Invalid type')
  }
}