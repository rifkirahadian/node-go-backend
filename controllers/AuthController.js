const responser = require('../modules/responser')
const validator = require('../modules/validator')
const auth = require('../modules/auth')

exports.register = async(req, res) => {
  try {
    validator.formValidate(req, res)

    const {name, phone, role} = req.body
    let user = await auth.nameExistCheck(name)
    user = await auth.createUserIfNotExist(name, phone, role, user)

    return responser.successResponse(res,user, 'Register Success')
  } catch (error) {
    return responser.errorResponseHandle(error, res)
  }
}