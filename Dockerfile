FROM golang:1.17-alpine
WORKDIR /app
COPY . ./
RUN go build ./cmd/godb
ENV PORT=6342
EXPOSE $PORT
CMD ["./godb"]
