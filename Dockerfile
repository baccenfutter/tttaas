FROM golang:1.13-alpine as builder

RUN apk update
RUN apk add bash make protobuf
RUN mkdir -p /go/src/github.com/baccenfutter/tttaas
WORKDIR /go/src/github.com/baccenfutter/tttaas
COPY . .
RUN make all

FROM scratch
LABEL maintainer Brian Wiborg <baccenfutter@c-base.org>
COPY --from=builder /go/bin/tictactoe .

EXPOSE 8000 8080

CMD [ "./tictactoe" ]
