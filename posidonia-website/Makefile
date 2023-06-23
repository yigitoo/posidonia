all: run
install:
	gem install bundler
	bundle install
run:
	bundle exec rerun 'rackup -p 8008' > /dev/null 2>&1
.PHONY: all run install
