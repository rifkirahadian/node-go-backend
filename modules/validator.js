const { validationResult } = require('express-validator')
const responser = require('./responser')

//validation form handle
exports.formValidate = (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    throw responser.formErrorValidationResponse(errors.array(), res)
  }
}