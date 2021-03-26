FROM alpine:3.12 as build

RUN apk --no-cache add go

ADD . /go/src/github.com/safecornerscoffee/echo-twitter
WORKDIR /go/src/github.com/safecornerscoffee/echo-twitter

RUN go build

FROM alpine:3.12

COPY --from=build /go/src/github.com/safecornerscoffee/echo-twitter/echo-twitter /usr/bin/echo-twitter

CMD [ "/usr/bin/echo-twitter" ]