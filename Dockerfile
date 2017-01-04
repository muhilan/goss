FROM golang

RUN mkdir /app

ADD ./src/ /app/

WORKDIR /app

RUN go build -o goss .

CMD ["/app/goss"]