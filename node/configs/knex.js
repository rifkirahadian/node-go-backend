// knex configuration settting
require('dotenv').config();
const environment = process.env.NODE_ENV
const config = require('../knexfile')[environment]

module.exports = require('knex')(config)