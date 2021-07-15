FROM golang:alpine AS build

ENV PROJECT_DIR="/go/src/github.com/deltabrot/salon-booking-guru"
WORKDIR $PROJECT_DIR

COPY ./api $PROJECT_DIR

RUN go get -d -v ./... && \
    go install -v ./...

RUN go build -o /bin/salon-booking-guru

FROM alpine

WORKDIR /root

COPY --from=build /bin/salon-booking-guru /bin/salon-booking-guru

EXPOSE 8085

CMD /bin/salon-booking-guru
