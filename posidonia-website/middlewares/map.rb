require 'sinatra/base'
require 'json'

#@brief: for communicate with golang backend
require 'net/http'
require 'uri'
require 'dotenv'

module Middleware
    class Map < Sinatra::Base
        Dotenv.load()
        enable :sessions

        before '/getAddr' do
            request.body.rewind
            @request_payload_getAddr = JSON.parse request.body.read
        end

        before '/getBbox' do
            request.body.rewind
            @request_payload_getBbox = JSON.parse request.body.read
        end

        post '/addItem' do
            if session[:id] and session[:username] then
                nil
            else
                redirect to('/'), 403
            end
        end

        post '/getBbox' do

            latitude = @request_payload_getBbox["lat"]
            longitude = @request_payload_getBbox["lng"]

            uri = URI("http://localhost:#{ENV['GO_PORT']}/bbox/#{latitude}/#{longitude}")
            response = Net::HTTP.get_response(uri)
            puts response.body if response.is_a?(Net::HTTPSuccess)

            content_type :json
            if !(response.is_a?(Net::HTTPSuccess))
                return {
                    "status": 404,
                    "bbox": "TESPİT EDİLEMEDİ",
                }.to_json
            else
                parsed_response = JSON.parse(response.body)
                return {
                    "status": 200,
                    "bbox_list": parsed_response['bbox_list'],
                    "x_min": parsed_response['x_min'],
                    "y_min": parsed_response['y_min'],
                    "x_max": parsed_response['x_max'],
                    "y_max": parsed_response['y_max'],
                }.to_json
            end
        end

        post '/getAddr' do

            latitude = @request_payload_getAddr["lat"]
            longitude = @request_payload_getAddr["lng"]

            uri = URI("http://localhost:#{ENV['GO_PORT']}/coordinates/#{latitude}/#{longitude}")
            #params = { :limit => 10, :page => 3 }
            #uri.query = URI.encode_www_form(params)
            response = Net::HTTP.get_response(uri)
            puts response.body if response.is_a?(Net::HTTPSuccess)

            content_type :json

            if !(response.is_a?(Net::HTTPSuccess))
                return {
                    "status": 404,
                    "addr": "TESPİT EDİLEMEDİ!",
                }.to_json
            else
                return {
                    "status": 200,
                    "addr": JSON.parse(response.body)["result"],
                }.to_json
            end
        end
    end

end
