all: run
install:
	gem install bundler
	bundle install
run:
	rerun 'rackup -p 8080'
.PHONY: all run install
