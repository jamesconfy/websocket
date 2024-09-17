migrate_up:
	migrate -path db/migration -database "" -verbose up

migrate_down:
	migrate -path db/migration -database "-verbose down

migrate_force:
	migrate -path db/migration -database " force $(version)

run:	
	go build websocket.go && ./websocket --migrate=false

run_migrate:
	go build websocket.go && ./websocket --migrate=true

gotidy:
	go mod tidy

goinit:
	go mod init

swag:
	swag init -g websocket.go -ot go,yaml 

migrate_init:
	migrate create -ext sql -dir db/migration -seq init_schema

launch:
	flyctl launch

docker_init:
	docker build -t everybody8/websocket:v$(version) .

docker_push:
	docker push everybody8/websocket:v$(version)

deploy:
	flyctl deploy

test:
	go test ./tests/repo_test && go test ./tests/service_test && go test ./tests/handler_test

repo_test:
	go test ./tests/repo_test

service_test:
	go test ./tests/service_test

handler_test:
	go test ./tests/handler_test