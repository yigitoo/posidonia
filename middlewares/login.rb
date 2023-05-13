require 'sinatra/base'
require 'sqlite3'

require_relative '../api'

module Middleware
    class Login < Sinatra::Base
        enable :sessions

        post ('/login') do
            pp request.POST.inspect
        end
    end
end
