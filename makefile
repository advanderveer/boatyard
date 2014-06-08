# build for production
build:
	cd res && gulp build:prod
	go build -tags release -o bin/boatyard

# build and start production
start: build
	@ ./bin/boatyard

# development interation
dev:
	# iterating..
	cd res && gulp build:dev
	go build -tags debug -o bin/boatyard
	./bin/boatyard