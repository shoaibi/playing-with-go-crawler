FROM golang:alpine as debug

ENV CGO_ENABLED 0
ENV APP_NAME="crawler"
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates \
        git \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

RUN go get github.com/go-delve/delve/cmd/dlv

COPY ./src $GOPATH/src/$APP_NAME/

COPY bin/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

WORKDIR $GOPATH/src/$APP_NAME
RUN go build -gcflags "-N -l" -o /crawler.app

ENTRYPOINT ["/entrypoint.sh"]
CMD ["dlv", "--headless=true", "--listen=:2345", "--api-version=2", "--accept-multiclient", "exec", "/crawler.app"]

###########START NEW IMAGE###################
FROM alpine as prod
COPY --from=debug /crawler.app /
CMD /crawler.app