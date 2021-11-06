# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM golang:1.16 AS stage-compile

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/sigma
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/sigma

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM alpine:latest 

WORKDIR /root/
COPY --from=stage-compile /go/src/app/sigma ./
COPY --from=stage-compile /go/src/app/settings.yaml ./

RUN apk --no-cache add ca-certificates

ENTRYPOINT [ "./sigma" ]  