const AuthController = require('../controllers/AuthController')
const validator = require('../modules/validator')

module.exports = (app, express) => {
  const apiRoutes = express.Router()

  apiRoutes.post('/register',validator.register, AuthController.register)
  apiRoutes.post('/login', validator.login, AuthController.login)

  app.use('/api', apiRoutes)
}