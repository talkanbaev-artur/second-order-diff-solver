FROM golang:alpine as build 
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY src ./src
RUN go build -o /main src/backend/cmd/main.go

FROM alpine:latest
WORKDIR /
ARG version
COPY --from=build /main /
COPY build/$version/frontend /frontend
ENTRYPOINT [ "/main" ]