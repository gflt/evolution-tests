FROM golang:alpine

WORKDIR /go/src/app

COPY . .

ENV PORT=8000
ENV PORT=5433

# RUN apk add --no-cache gcc musl-dev postgresql-dev

CMD ["go", "run", "main.go"]