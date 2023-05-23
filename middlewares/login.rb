require 'sinatra/base'
require 'sqlite3'

module Middleware
    class Login < Sinatra::Base
        enable :sessions

        post ('/login') do
            database = SQLite3::Database.new "db/posidonia.sqlite3"

            username = params["username"].to_s
            password = params["password"].to_s

            database.execute ("select * from users") do |row|
                if row[1] == username and row[2] == password then
                    session[:id] = row[0]
                    session[:username] = row[1]
                    session[:password] = row[2]
                    redirect to('/addItem'), 301
                end
            end
            redirect to('/login'), 301
        end
    end
end
