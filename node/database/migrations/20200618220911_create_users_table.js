
exports.up = function(knex) {
  return knex.schema
    .createTable('users', function (table) {
      table.increments('id')
      table.string('phone', 20).notNullable().unique()
      table.string('name', 150).notNullable().unique()
      table.string('password', 5).notNullable()
      table.string('role', 20).notNullable()
      table.datetime('deleted_at')
      table.timestamps(true, true)
    })
};

exports.down = function(knex) {
  return knex.schema.dropTable('users')
};
