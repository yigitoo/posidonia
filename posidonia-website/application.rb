#!/usr/bin/ruby
#-------
#@Author: Yiğit GÜMÜŞ <github.com/yigitoo>
#@Date: May 13 2023 - still continues.
#-------
require 'sinatra/base'
require 'sinatra'
require 'json'

#@brief: middlewares
require_relative 'middlewares/login'
require_relative 'middlewares/map'

#@brief: Main App!
class PosidoniaServer < Sinatra::Base
    # set root folder of the project
    set :root, File.dirname(__FILE__)

    #@brief: app settings and middlewares/controllers.
    use Middleware::Login # for session registeration and login actions
    use Middleware::Map   # for mapping technologies.

    set :port, 1234
    set :bind, '0.0.0.0'
    set :public_folder, __dir__ + '/static'

    #@brief: Routes and requests!
    before do
        @session_user = get_session_user()
    end
    get ('/') { erb :index }
    get ('/home') { erb :index }
    get ('/privacy_policy') { erb :privacy_policy }
    get ('/login') {
        if session[:id] and session[:username] then
            redirect to('/'), 301
        else
            erb :login
        end
    }
    get ('/congratulations') { erb :congratulations}
    get ('/map') { erb :map, layout: false}
    get ('/addItem') {
        if session[:id] and session[:username] then
            erb :add_item, layout: false
        else
            redirect to('/'), 301
        end
    }
    get ('/logout') {
        session.delete(:id)
        session.delete(:username)
        redirect '/login'
    }
    #@brief: static files
    get ('/logo.png') { send_file File.expand_path('logo.png', settings.public_folder) }
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
        if session[:username] == nil
            return nil
        else
            return true
        end
    end

    run! if app_file == $0
end
