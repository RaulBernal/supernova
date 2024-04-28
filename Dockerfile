#===============
# Stage 1: Build
#===============

FROM golang:1.21-alpine as builder

COPY . /app

WORKDIR /app

RUN go build -o supernova ./cmd

#===============
# Stage 2: Run
#===============

FROM alpine

COPY --from=builder /app/supernova /usr/local/bin/supernova
COPY --from=builder /app  /
