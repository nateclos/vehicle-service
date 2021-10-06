FROM golang:bullseye

WORKDIR /app

COPY * ./

COPY src/go.mod ./

COPY src/go.sum ./

RUN go mod download

COPY src/*.go ./

RUN go build -o /app/bin/vehicle-service

EXPOSE 8080

CMD [ "/app/bin/vehicle-service" ]