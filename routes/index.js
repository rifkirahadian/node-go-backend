const AuthController = require('../controllers/AuthController')
const FetchingController = require('../controllers/FetchingController')
const validator = require('../modules/validator')

module.exports = (app, express) => {
  const apiRoutes = express.Router()
  const authRoutes = express.Router()
  const adminRoutes = express.Router()

  //middleware
  require('../middlewares/Auth')(authRoutes)
  require('../middlewares/Admin')(adminRoutes)

  apiRoutes.post('/register',validator.register, AuthController.register)
  apiRoutes.post('/login', validator.login, AuthController.login)

  authRoutes.get('/auth/user', AuthController.authUser)
  authRoutes.get('/fetching/fetch', FetchingController.fetch)
  adminRoutes.get('/fetching/aggregate', FetchingController.fetchAggregate)

  app.use('/api', apiRoutes)
  app.use('/api', authRoutes)
  app.use('/api/admin', adminRoutes)
}