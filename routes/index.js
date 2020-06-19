const AuthController = require('../controllers/AuthController')
const FetchingController = require('../controllers/FetchingController')
const validator = require('../modules/validator')

module.exports = (app, express) => {
  const apiRoutes = express.Router()
  const authRoutes = express.Router()

  //middleware
  require('../middlewares/Auth')(authRoutes)

  apiRoutes.post('/register',validator.register, AuthController.register)
  apiRoutes.post('/login', validator.login, AuthController.login)

  authRoutes.get('/auth/user', AuthController.authUser)
  authRoutes.get('/fetching/fetch', FetchingController.fetch)

  app.use('/api', apiRoutes)
  app.use('/api', authRoutes)
}