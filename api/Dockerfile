FROM golang:alpine AS build

ENV PROJECT_DIR="/go/src/github.com/deltabrot/clere-coding-challenge-api"
WORKDIR $PROJECT_DIR

COPY . $PROJECT_DIR

RUN go get -d -v ./... && \
    go install -v ./...

RUN go build -o /bin/clere-coding-challenge-api

FROM alpine

WORKDIR /root

COPY --from=build /bin/clere-coding-challenge-api /bin/clere-coding-challenge-api

EXPOSE 8085

CMD /bin/clere-coding-challenge-api
