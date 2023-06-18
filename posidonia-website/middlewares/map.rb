require 'sinatra/base'
require 'json'

#@brief: for communicate with golang backend
require 'net/http'
require 'uri'
require 'dotenv'
require 'date'


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
            request.body.rewind
            @request_payload_addItem = JSON.parse request.body.read

            if session[:id] and session[:username] then

                polygon_list = @request_payload_addItem['locations']
                pp polygon_list
                polygon_architecture = Array.new()
                for i in polygon_list do
                    string_of_one_coord = i[0].to_s + ':' + i[1].to_s
                    polygon_architecture.append(string_of_one_coord)
                end

                uri = URI("http://localhost:#{ENV['GO_PORT']}/addPolygon")
                http = Net::HTTP.new(uri.host, uri.port)
                request = Net::HTTP::Post.new(uri.path, 'Content-Type' => 'application/json')
                request.body = {
                    polygon: polygon_architecture,
                    addedBy: session[:username],
                    addedTime: Time.now.strftime("%d/%m/%Y %H:%M")
                }.to_json
                response = http.request(request)
                body = JSON.parse response.body

                content_type :json

                if body["successful"] == true then
                    {
                        "message": "Place added succesfully by #{session[:username]}",
                        "status": 200,
                    }.to_json
                else
                    {
                        "message": "Server error",
                        "status": 500,
                    }.to_json
                end
            else
                redirect to('/'), 403
            end
        end

        get ('/dumpPolygon') {
            uri = URI("http://localhost:#{ENV['GO_PORT']}/dumpPolygon")
            response = Net::HTTP.get_response(uri)
            content_type :json

            if !(response.is_a?(Net::HTTPSuccess)) then
                return {
                    "locations": "",
                    "successful": false,
                }.to_json
            else
                parsed_response = JSON.parse(response.body)
                return {
                    "locations": parsed_response['all_locations'],
                    "successful": parsed_response['successful']
                }.to_json
            end
        }

        post '/getBbox' do

            latitude = @request_payload_getBbox["lat"]
            longitude = @request_payload_getBbox["lng"]

            uri = URI("http://localhost:#{ENV['GO_PORT']}/bbox/#{latitude}/#{longitude}")
            response = Net::HTTP.get_response(uri)

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
