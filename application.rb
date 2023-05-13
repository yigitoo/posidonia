#!/usr/bin/ruby
#-------
#@Author: Yiğit GÜMÜŞ <github.com/yigitoo>
#@Date: May 13 2023 - still continues.
#-------
require "sinatra/activerecord"
require 'sinatra/base'
require 'sinatra'
require 'json'

#@brief: database actions, models and schemas files
require_relative 'db/models'
require_relative 'api'

#@brief: middlewares
require_relative 'middlewares/login'
require_relative 'middlewares/map.rb'

#@brief: Main App!
class PosidoniaServer < Sinatra::Base
    #@brief: datbase migration
    register Sinatra::ActiveRecordExtension

    #@brief: app settings and middlewares/controllers.
    use Middleware::Login # for session registeration and login actions
    use Middleware::Map   # for mapping technologies.

    set :port, 8080
    set :bind, '0.0.0.0'
    set :public_folder, __dir__ + '/static'

    #@brief: Routes and requests!
    before do
        @session_user = get_session_user()
    end
    get ('/') { erb :index }
    get ('/home') { erb :index }
    get ('/map') { erb :map }
    get ('/login') { erb :login }
    get '/logout' do
        session.delete(:user_id)
        redirect '/login'
    end
    #@brief: static files
    get ('/logo.png') { send_file File.expand_path('logo.png', settings.public_folder) }
    get ('/logo.svg') { send_file File.expand_path('logo.svg', settings.public_folder) }
    get ('/Posidonia.png') { send_file File.expand_path('Posidonia.png', settings.public_folder) }

    #@description: error code situations
    error 403 do
        content_type :json
        {
            :status => 403,
            :message => "Error: Access forbidden."
        }.to_json
    end

    error 404 do
        redirect '/', 301
    end

    # own functions for utilities
    def get_session_user
        if session['user_name'] == nil
            return nil
        else
            return true
        end
    end

    run! if app_file == $0
end
