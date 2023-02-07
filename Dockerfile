FROM golang:1.19.5 as builder

WORKDIR /app
COPY . /app
RUN make build


FROM ubuntu:20.04
COPY --from=builder /app/bin/server /app/bin/server
COPY --from=builder /app/static /app/static
COPY --from=builder /app/migrations /app/migrations

RUN apt-get update && apt-get install ca-certificates -y
WORKDIR /app

ENTRYPOINT [ "./bin/server" ]
