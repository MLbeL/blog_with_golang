FROM golang:1.25

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . /usr/src/app/
RUN go build -v -o /usr/local/bin/app ./cmd \
	&& go build -v -o /usr/local/bin/migrate ./migrations

EXPOSE 8080

CMD ["app"]
