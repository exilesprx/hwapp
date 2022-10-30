FROM golang:1.18.3 as build

WORKDIR /go/src

COPY ./ /go/src/

COPY .docker/entrypoint.sh /go/src/

RUN chmod +x /go/src/entrypoint.sh

RUN make install && make build


FROM golang:1.18.3 as app

COPY --from=build /go/src/build/go-hwapp /go/app

COPY --from=build /go/src/entrypoint.sh /go/entrypoint.sh

COPY --from=build /go/src/Makefile /go/


ENTRYPOINT [ "/go/entrypoint.sh" ]