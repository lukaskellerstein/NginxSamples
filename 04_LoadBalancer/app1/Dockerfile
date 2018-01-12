# -----------------------
# Stage 1 - BUILD
# -----------------------
FROM golang
ADD . $GOPATH/src/github.com/cellarstone/nginxtest
WORKDIR $GOPATH/src/github.com/cellarstone/nginxtest

RUN go-wrapper download
RUN go-wrapper install
RUN go build -o testapp1

# -----------------------
# Stage 2 - RUNTIME 
# -----------------------
FROM golang

# WORKDIR /app
COPY --from=0 /go/src/github.com/cellarstone/nginxtest/testapp1 .

CMD ["./testapp1"]

