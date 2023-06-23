require 'sinatra/base'
require 'dotenv'
require 'json'

module Middleware
    class Login < Sinatra::Base
        Dotenv.load
        enable :sessions

        post ('/login') do
            username = params["username"].to_s
            password = params["password"].to_s
            begin
                uri = URI("http://localhost:#{ENV['GO_PORT']}/login")
                http = Net::HTTP.new(uri.host, uri.port)
                request = Net::HTTP::Post.new(uri.path, 'Content-Type' => 'application/json')
                request.body = {username: username, password: password}.to_json
                response = http.request(request)
                body = JSON.parse(response.body)

                session[:id] = body["user_id"]
                session[:username] = body["username"]

                to_main = true
            rescue => error
                to_main = false
            end

            if to_main == true
                redirect to('/addItem'), 301
            else
                redirect to('/login'), 301
            end
        end
    end
end
