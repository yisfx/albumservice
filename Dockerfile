FROM alpine

WORKDIR /app

COPY . /app

CMD ["./albumservice"]