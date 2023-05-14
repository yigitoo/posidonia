require 'sinatra/base'
require 'sqlite3'

module Middleware
    class Login < Sinatra::Base
        enable :sessions

        post ('/login') do
            database = SQLite3::Database.new "db/posidonia.sqlite3"

            username = params["username"]
            password = params["password"]
=begin
            query = <<-SQL
                PRAGMA encoding="UTF-8";
                select * from users
                where email = ? and password
            SQL
=end
            database.execute( "select * from users" ) do |row|
                if row[1] == username && row[2] == password
                    session[:id] = row[0]
                    session[:username] = row[1]
                    session[:password] = row[2]
                    redirect to('/'), 301
                else
                    redirect to('/login'), 301
                end
            end
        end
    end
end
