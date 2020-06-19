const AuthController = require('../controllers/AuthController')
const validator = require('../modules/validator')

module.exports = (app, express) => {
  const apiRoutes = express.Router()
  const authRoutes = express.Router()

  //middleware
  require('../middlewares/Auth')(authRoutes)

  apiRoutes.post('/register',validator.register, AuthController.register)
  apiRoutes.post('/login', validator.login, AuthController.login)

  app.use('/api', apiRoutes)
  app.use('/api', authRoutes)
}