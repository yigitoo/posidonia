require 'sinatra'
require 'json'

require_relative 'database'

users = create_client('users')

set :port, 8080
set :host, '0.0.0.0'

get '/' do
    
    content_type :json
    {
        :name => "Hello World"
    }.to_json
end