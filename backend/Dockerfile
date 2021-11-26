# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM golang:1.17 AS stage-compile

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/shortlink2
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/shortlink2

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM alpine:latest 

WORKDIR /root/
COPY --from=stage-compile /go/src/app/shortlink2 ./
COPY --from=stage-compile /go/src/app/settings.yaml ./

RUN apk --no-cache add ca-certificates

ENTRYPOINT [ "./shortlink2" ]  