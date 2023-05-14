require 'sqlite3'

database = SQLite3::Database.new "db/posidonia.sqlite3"
result = database.execute <<-SQL
  insert into users values (1, "yigitgumus", "adminsifresi123")
SQL
puts result
