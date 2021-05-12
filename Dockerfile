FROM alpine

WORKDIR /app

COPY . /app
EXPOSE 9001
CMD ["./albumservice"]