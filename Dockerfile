FROM golang:1.17 as builder

WORKDIR /app
COPY . /app

RUN go build -o /bin/recipes

FROM ubuntu:18.04
COPY --from=builder /bin/recipes /bin/recipes
COPY .env /.env
COPY migrations /migrations

# RUN apt install ca-certificates && update-ca-certificates

ENTRYPOINT [ "/bin/recipes" ]