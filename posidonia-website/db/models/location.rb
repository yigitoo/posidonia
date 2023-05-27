class Location < ActiveRecord::Base
    validates_presence_of :polygons
    validates_presence_of :added_by
    validates_presence_of :added_date
end
