migrate_up:
	migrate -path db/migration -database "postgres://localhost:password@localhost:5432/project_name?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "" -verbose down

run:
	go build project-name-api.go && ./project-name-api

gotidy:
	go mod tidy

goinit:
	go mod init

swag:
	swag init -g project-name-api.go -ot go,yaml 

migrate_init:
	migrate create -ext sql -dir db/migration -seq init_schema

launch:
	flyctl launch

docker_init:
	docker build -t everybody8/project-name-api:v$(version) .

docker_push:
	docker push everybody8/project-name-api:v$(version)

deploy:
	flyctl deploy