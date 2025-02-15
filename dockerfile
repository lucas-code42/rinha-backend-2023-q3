FROM golang:1.22.1

WORKDIR /application

COPY . .

RUN go build cmd/main.go

EXPOSE 80

ENTRYPOINT [ "./main" ]