FROM golang:1.25

WORKDIR /usr/src/app

RUN mkdir -p /usr/src/app/images

COPY . .

RUN go build -v -o /usr/local/bin/goray

CMD ["goray"]