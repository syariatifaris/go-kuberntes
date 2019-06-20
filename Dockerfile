FROM golang:latest
COPY ./ /
RUN go get github.com/gorilla/mux
WORKDIR /
EXPOSE 9090
RUN go build -o main .
CMD ["./main"]

