#!/usr/bin/ruby
#-------
#@Author: Yiğit GÜMÜŞ <github.com/yigitoo>
#@Date: May 13 2023 - still continues.
#-------
require "sinatra/activerecord"
require 'sinatra/base'
require 'sinatra'
require 'json'
require 'dotenv'

#@brief: database actions, models and schemas files
require_relative 'db/models'
require_relative 'api'

#@brief: middlewares
require_relative 'middlewares/login'

#@brief: other files for spesific pages
require_relative 'components/map.rb'

#@brief: Main App!
class PosidoniaServer < Sinatra::Base
    #@brief: datbase migration
    register Sinatra::ActiveRecordExtension

    #@brief: app settings and middlewares.
    use Middleware::Login

    Dotenv.load()
    set :port, 8080
    set :bind, '0.0.0.0'
    set :public_folder, __dir__ + '/static'

    #GET REQUESTS
    get '/' do
        session_user = get_session_user()
        if session_user == nil
            redirect "/login"
        else
            erb :index
        end
    end
    
    get ('/map') { erb :map }
    get ('/login') { erb :login }

    get '/logout' do
        session.delete(:user_id)
        redirect '/login'
    end

    #POST REQUESTS
    post '/login' do
        
    end

    #@description: error code situations 
    error 403 do
        content_type :json
        {
            :status => 403,
            :message => "Error: Access forbidden."
        }.to_json
    end

    error 404 do
        content_type :json
        {
            :status => 404,
            :message => "Error: Page/User does not exist."
        }.to_json
    end

    # own functions for utilities
    def get_session_user
        if session[:user_id] == nil
            return nil
        else
            return @database_users.find({:_id => session[:user_id]}).first
        end
    end

    run! if app_file == $0
end