require 'sinatra/base'
require 'sqlite3'

require_relative '../api.rb'

module Controllers
    class Map < Sinatra::Base
        post '/admin/map' do
            
        end
    end

end