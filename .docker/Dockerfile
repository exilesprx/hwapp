FROM golang:1.18.3 as build

WORKDIR /go/src

COPY ./ /go/src/

RUN go get -d

RUN go build -o ./build/


FROM golang:1.18.3 as app

COPY --from=build /go/src/build/go-hwapp /go/app

ENTRYPOINT [ "/go/app" ]