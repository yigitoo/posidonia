require 'sinatra/base'
require 'sqlite3'
require 'json'

# for api requests.
require 'faraday'
require 'faraday/retry'
require 'faraday_middleware'

module Middleware
    class Map < Sinatra::Base
        enable :sessions

        before '/getAddr' do
            request.body.rewind
            @request_payload = JSON.parse request.body.read
        end

        post '/addItem' do


            if session[:id] and session[:username] and session_user[:username] then
                database = SQLite3::Database.new 'db/posidonia.sqlite3'


                database.execute ("select * from locations") do |row|
                    row_polygon_json = row[1].to_json

                    if @request_payload[:polygon] == row_polygon_json
                        is_exist = true
                    end
                end

                if is_exist then
                    redirect to('/addItem'), 301
                else
                    database.execute <<-SQL
                        INSERT INTO locations VALUES (

                        )
                    SQL
                end
            else
                redirect to('/'), 403
            end
        end

        post '/getAddr' do

            latitude = @request_payload[:lat]
            longtitude = @request_payload[:lng]

            @api_site_url = "https://api.geoapify.com"
            @request_url = "/v1/geocode/reverse?lat=#{latitude}&lon=#{longtitude}&apiKey=39b51f681a6345929728a75e57f5e32a"
            connection = Faraday.new(@api_site_url) do |f|
                f.request :json
                f.request :retry
                f.response :json
                f.adapter :net_http
            end

            response = connection.get(@request_url)

            content_type :json
            return {
                'addr': response.body['features'][0]['properties']['formatted'],
                'full': response.body
            }.to_json

        end
    end

end
