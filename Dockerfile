FROM golang:1.18-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

# build go
RUN go mod download
RUN go mod download github.com/jackc/chunkreader
RUN go mod download github.com/jackc/pgproto3
RUN go build cmd/main.go

# build js
RUN apk add --update nodejs npm
RUN npm install webpack
RUN npm run build

EXPOSE 8880

CMD [ "./main" ]