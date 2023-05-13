require 'sinatra/base'
require 'sqlite3'

require_relative '../api.rb'

module Controllers
    class Map < Sinatra::Base
        self.database_users = SQLite3::Database.new 'db/posidonia.sqlite3'
        
        
        get ('/map') {
            erb :login 
        }

        post ('/login') {
            
        }
    end

end