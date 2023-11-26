FROM golang:1.21
WORKDIR /godb
COPY . ./
RUN go build ./cmd/server
ENV PORT=6342
EXPOSE $PORT
CMD ["./server"]
