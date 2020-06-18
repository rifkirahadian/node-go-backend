
exports.up = function(knex) {
  return knex.schema
    .createTable('users', function (table) {
      table.increments('id')
      table.string('phone', 20).notNullable()
      table.string('name', 150).notNullable().unique()
      table.string('password', 5).notNullable()
      table.timestamps()
    })
};

exports.down = function(knex) {
  return knex.schema.dropTable('users')
};
