GATEWAY_BINARY=gatewayApp

up: 
	@echo "Starting images"
	docker-compose up -d
	@echo "Images started"

up_build: build_gateway
	docker-compose down
	docker-compose up --build -d

down:
	docker-compose down

build_gateway:
	cd ./gateway-service && env GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} ./cmd/api


