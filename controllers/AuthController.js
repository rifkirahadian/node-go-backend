const responser = require('../modules/responser')
const validator = require('../modules/validator')
const auth = require('../modules/auth')

exports.register = async(req, res) => {
  try {
    validator.formValidate(req, res)

    const {name, phone, role} = req.body
    let user = await auth.nameExistCheck(name)
    user = await auth.createUserIfNotExist(name, phone, role, user, res)
    
    return responser.successResponse(res,user, null)
  } catch (error) {
    return responser.errorResponseHandle(error, res)
  }
}

exports.login = async(req, res) => {
  try {
    validator.formValidate(req, res)
    const {phone, password} = req.body
    let user = await auth.getUserByPhone(phone, res)
    auth.passwordMatchValidate(user.password, password, res)

    const authToken = auth.authTokenGenerate(user)

    return responser.successResponse(res,authToken, null)
  } catch (error) {
    return responser.errorResponseHandle(error, res)
  }
}