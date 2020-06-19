const knex = require('../configs/knex')
const responser = require('./responser')
const randomstring = require('randomstring')

exports.nameExistCheck = async(name) => {
  return await knex('users').where({name}).first('password')
}

exports.createUserIfNotExist = async(name, phone, role, user) => {
  if (!user) {
    const password = randomstring.generate({
      length: 4,
      charset: 'numeric'
    })

    return await knex('users').insert({name, phone, role, password}).returning('password')
  }else{
    return user
  }
}