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

exports.fetchAggregate = async(req, res) => {
  try {
    let data = json.readJsonFileSync('../data/aggregate.json')
    const {type} = req.query
    fetching.aggregateTypeValidate(type, res)

    if (type === 'provincyArea') {
      data = fetching.groupByAreaProvincies(data)
    }else{
      data = fetching.groupByWeeklyDate(data)
    }
    
    let dataAggregated = fetching.getAggregates(data)

    return responser.successResponse(res, dataAggregated, null)
  } catch (error) {
    return responser.errorResponseHandle(error, res)
  }
}