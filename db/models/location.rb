class Location < ActiveRecord::Base
    validates_presence_of :lat
    validates_presence_of :long
    validates_presence_of :added_date
end