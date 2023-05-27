require 'sinatra/base'
require 'sqlite3'
require 'json'

#@brief: for communicate with golang backend
require 'net/http'
require 'dotenv'

module Middleware
    class Map < Sinatra::Base
        Dotenv.load
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


            PORT_NO = ENV['GO_PORT']
            response = Net::HTTP.get_response("localhost:#{PORT_NO}")
            content_type :json
            return {
                'addr': response.body['features'][0]['properties']['formatted'],
                'full': response.body
            }.to_json

        end
    end

end
