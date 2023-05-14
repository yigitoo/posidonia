require 'sinatra/base'
require 'sqlite3'

module Middleware
    class Map < Sinatra::Base
        enable :sessions

        post '/admin/map' do

        end
    end

end
