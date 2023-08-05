all: run
ifneq ($(OS), Windows_NT)
run:
	bundle exec rerun 'rackup -p 8008' > /dev/null 2>&1
else
run:
	bundle exec rerun 'rackup -p 8008'
endif
install:
	gem install bundler
	bundle install

.PHONY: all run install
