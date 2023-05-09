migrate_up:
	migrate -path db/migration -database "" -verbose up

migrate_down:
	migrate -path db/migration -database "-verbose down

migrate_force:
	migrate -path db/migration -database " force $(version)

run:
	@if [ "$(migrate)" == "true" ]; \
	then \
		go build project-name-api.go && ./project-name-api --migrate=true; \
    else \
		go build project-name-api.go && ./project-name-api --migrate=false; \
    fi

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