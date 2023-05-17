FROM golang:1.19-alpine3.16

RUN mkdir /project-name

COPY . /project-name

WORKDIR /project-name

LABEL Name=project-name Version=0.0.1

RUN go build -o project-name-api

EXPOSE  8080

CMD [ "./project-name-api", "--migrate=true" ]