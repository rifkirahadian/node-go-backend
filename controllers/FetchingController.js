const responser = require('../modules/responser')
const fetching = require('../modules/fetching')
const json = require('../modules/json')

exports.fetch = async (req, res) => {
  try {
    const {currency} = req.query
    let data = json.readJsonFileSync('../data/fetching.json')
    data = await fetching.handlePriceDataCurrency(data, currency, res)

    return responser.successResponse(res, data, null)
  } catch (error) {
    return responser.errorResponseHandle(error, res)
  }
}