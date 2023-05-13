class CreateLocations < ActiveRecord::Migration[7.0]
  def change
    create_table :locations do |t|
      t.string :lat
      t.string :long
      t.string :added_date
    end
  end
end