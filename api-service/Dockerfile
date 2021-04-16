FROM golang:1.14-alpine AS build


WORKDIR /api-service/
COPY . .
RUN CGO_ENABLED=0 go mod tidy
RUN CGO_ENABLED=0 go mod download
RUN CGO_ENABLED=0 go build -tags=musl -o /bin/app src/main.go

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]