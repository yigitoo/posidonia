require 'dotenv/load'
require 'mongo'

def create_client(name)
    db_uri = ENV['DB_URI']
    puts db_uri
    client = Mongo::Client.new(db_uri)
    database = client[name.to_sym]
    return database
end