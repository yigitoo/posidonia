require 'sinatra/base'
require 'sqlite3'

module Middleware

    class Login < Sinatra::Base
        enable :sessions

        post ('/login') do
            
        end
    end
end