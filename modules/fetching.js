const api = require('./api')
const servicesConfig = require('../configs/services')
const responser = require('./responser')
const NodeCache = require("node-cache");

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