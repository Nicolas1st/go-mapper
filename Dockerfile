FROM golang:1.18-alpine

RUN mkdir /app
COPY go.mod /app
COPY go.sum /app
COPY package.json /app
COPY package-lock.json /app
WORKDIR /app

# go deps
RUN go mod download
RUN go mod download github.com/jackc/chunkreader
RUN go mod download github.com/jackc/pgproto3

# js deps
RUN apk add --update nodejs npm
RUN npm install webpack

COPY . /app
RUN go build cmd/main.go
RUN npm run build

EXPOSE 8880

CMD [ "./main" ]