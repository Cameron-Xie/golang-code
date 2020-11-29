FROM golang:1.15-buster

ENV GolangCILint_Version=1.33.0

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
        sh -s -- -b $(go env GOPATH)/bin v${GolangCILint_Version} \
        && go get github.com/smartystreets/goconvey

