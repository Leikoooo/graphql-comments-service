FROM golang:1.19-alpine

WORKDIR /app

RUN apk add --no-cache bash gcc musl-dev

RUN apk add --no-cache wget && \
    wget -O migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    rm migrate.tar.gz

RUN apk add --no-cache dos2unix

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./src/cmd

COPY run.sh .

RUN dos2unix run.sh

RUN chmod +x run.sh

ENTRYPOINT ["./run.sh"]
