require 'sinatra/base'
require 'sqlite3'

module Middleware
    class Map < Sinatra::Base
        enable :sessions

        post '/map/addItem' do
            if session[:id] and session[:username] and session_user[:username] then
                database = SQLite3::Database.new 'db/posidonia.sqlite3'

                database.execute <<-SQL
                    INSERT INTO locations VALUES (
                        JSON
                    )
                SQL
            else
                redirect to('/'), 403
            end
        end
    end

end
