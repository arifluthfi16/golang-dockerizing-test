FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/arifluthfi16/golang-dockerizing-test
RUN cd /build && git clone https://github.com/arifluthfi16/golang-dockerizing-test.git

RUN cd /build/golang-dockerizing-test/app && go build

EXPOSE 4016

ENTRYPOINT [ "/build/golang-dockerizing-test/app/golang-dockerizing-test"]