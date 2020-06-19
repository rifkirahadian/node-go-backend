const { validationResult, body } = require('express-validator')
const responser = require('./responser')

//validation form handle
exports.formValidate = (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    throw responser.formErrorValidationResponse(errors.array(), res)
  }
}

//register validation rule
exports.register = [
  body(['name']).notEmpty().withMessage('Name required'),
  body(['phone']).notEmpty().withMessage('Phone required'),
  body(['role']).notEmpty().withMessage('Role required')
]