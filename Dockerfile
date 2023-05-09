# #build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go build -o /go/bin/app -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=project-name Version=0.0.1
# EXPOSE 5000

FROM golang:1.19-alpine3.16

RUN mkdir /project-name

COPY . /project-name

WORKDIR /project-name

LABEL Name=project-name Version=0.0.1

RUN go build -o project-name-api

EXPOSE  8080

CMD [ "./project-name-api", "--migrate=true" ]