const responser = require('../modules/responser')

module.exports = (adminRoutes) => {
  adminRoutes.use((req, res, next) => {
    let user = req.user
    if (user.role !== 'admin') {
      return responser.errorResponseStatus(res, 403, 'Forbidden request')
    }

    next()
  })
}