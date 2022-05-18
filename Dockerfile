FROM golang:1.18-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download
RUN go mod download github.com/jackc/chunkreader 
RUN go mod download github.com/jackc/pgproto3
RUN go build cmd/main.go
EXPOSE 8880

CMD [ "./main" ]
