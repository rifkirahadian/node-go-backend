const knex = require('../configs/knex')
const responser = require('./responser')
const randomstring = require('randomstring')

exports.nameExistCheck = async(name) => {
  return await knex('users').where({name}).first('password')
}

exports.createUserIfNotExist = async(name, phone, role, user, res) => {
  try {
    if (!user) {
      const password = randomstring.generate({
        length: 4,
        charset: 'numeric'
      })

      let insertUser = await knex('users').insert({name, phone, role, password})
      return await knex('users').whereIn('id', insertUser).first('password')
    }else{
      return user
    }
  } catch (error) {
    return responser.errorResponse(res, 'Phone number has been used by another user')
  }
}