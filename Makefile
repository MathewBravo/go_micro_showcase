GATEWAY_BINARY=gatewayApp
SEARCH_BINARY=searchApp
LIBRARY_BINARY=libraryApp

up: 
	@echo "Starting images"
	docker-compose up -d
	@echo "Images started"

up_build: build_gateway build_search build_library
	docker-compose down
	docker-compose up --build -d

down:
	docker-compose down

build_gateway:
	cd ./gateway-service && env GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} ./cmd/api

build_search:
	cd ./search-service && env GOOS=linux CGO_ENABLED=0 go build -o ${SEARCH_BINARY} ./cmd/api

build_library:
	cd ./library-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LIBRARY_BINARY} ./cmd/api



