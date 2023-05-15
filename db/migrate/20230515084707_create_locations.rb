class CreateLocations < ActiveRecord::Migration[7.0]
  def change
    create_table :locations do |t|
      t.text :polygon
      t.string :added_by
      t.string :added_date
    end
  end
end
