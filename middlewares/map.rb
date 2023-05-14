require 'sinatra/base'
require 'sqlite3'

module Middleware
    class Map < Sinatra::Base
        enable :sessions

        post '/map/addItem' do
            if session[:id] and session[:username] and session_user[:username]
                redirect to('/map'), 200
            else
                redirect to('/'), 301
            end
        end
    end

end
