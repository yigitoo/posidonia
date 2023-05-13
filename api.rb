require 'sqlite3'
require 'yaml'

module API
    class Database
        attr_accessor :database, :database_config, :last_query
        attr_reader :version

        def initialize
            @version = "v1"
=begin
            configuration of database (example):
            adapter: sqlite3
            database: db/posidonia.sqlite3
            pool: 5
            timeout: 5000
=end
            @database_config = YAML.load_file('config/database.yml').inspect
            
            @database = SQLite3::Database.new(@database_file_content["database"])
        end

        def execute(query)
            @last_query = query
            result = @database.execute query
            return result
        end

        def just_execute(query)
            @last_query = query
            @database.execute query 
        end
    end
end