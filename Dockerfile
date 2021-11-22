# Start by building the application.
FROM golang:1.17 as build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./...

RUN go build

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/src/app/webhook-recorder /
USER nonroot
EXPOSE 1323/tcp
CMD ["/webhook-recorder"]