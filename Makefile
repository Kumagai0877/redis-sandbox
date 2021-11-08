build:
	cd ./api && \
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.artifacts/api-linux-amd64 ./main.go

up: build
	docker-compose -f docker-compose.yaml up

down:
	docker-compose -f docker-compose.yaml down

logs:
	docker-compose -f docker-compose.yaml logs -f