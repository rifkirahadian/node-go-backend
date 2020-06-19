const knex = require('../configs/knex')
const responser = require('./responser')
const randomstring = require('randomstring')
const jwt = require('jsonwebtoken')
const moment = require('moment')

exports.nameExistCheck = async (name) => {
  return await knex('users').where({ name }).first('password')
}

exports.createUserIfNotExist = async (name, phone, role, user, res) => {
  try {
    if (!user) {
      const password = randomstring.generate({
        length: 4,
        charset: 'numeric'
      })

      let insertUser = await knex('users').insert({ name, phone, role, password })
      return await knex('users').whereIn('id', insertUser).first('password')
    } else {
      return user
    }
  } catch (error) {
    throw responser.errorResponse(res, 'Phone number has been used by another user')
  }
}

exports.getUserByPhone = async (phone, res) => {
  const user =  await knex('users').where({ phone }).first('name', 'phone', 'role', 'password')
  if (!user) {
    throw responser.errorResponseStatus(res, 401, 'Phone number not found')
  }

  return user
}

exports.passwordMatchValidate = (userPassword, inputPassword, res) => {
  if (userPassword !== inputPassword) {
    throw responser.errorResponseStatus(res, 401, 'Password doesn`t match')
  }
}

exports.authTokenGenerate = (user) => {
  const {name, phone, role} = user
  const timestamp = moment().format('YYYY-MM-DD HH:mm:ss')
  return jwt.sign({name, phone, role, timestamp}, process.env.JWT_SECRET, {
    expiresIn: 1000000 * 1440
  })
}