require 'sinatra/base'
require 'sqlite3'

module Middleware

    class Login < Sinatra::Base
        self.database_users = SQLite3::Database.new 'db/posidonia.sqlite3'
        rows = db.execute <<-SQL
            create table users(
                username varchar(30),
                password varchar(30),
            );
        SQL
        enable :sessions
        get ('/login') { erb :login }

        post ('/login') {
            if params['name'] == '' && params['password'] == ''
        }
    end

end