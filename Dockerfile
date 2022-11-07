FROM golang:1.19.2-alpine

RUN  mkdir /app

ADD . /app
COPY  . /app
WORKDIR /app

RUN go build -o main .

EXPOSE 8000

CMD [“/app/main”]