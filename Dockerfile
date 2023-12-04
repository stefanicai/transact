FROM golang:1.21-bullseye AS builder

WORKDIR /src

COPY go.mod go.sum ./

#sources
COPY internal internal
COPY main.go main.go
COPY Makefile Makefile
COPY config/ config/

RUN mkdir ./bin && make build

RUN cd ./bin && ls -LR

FROM gcr.io/distroless/base-debian12
COPY --from=builder /src/bin/transact /bin/entry
COPY --from=builder /src/config /config

EXPOSE 8080
ENTRYPOINT ["/bin/entry"]