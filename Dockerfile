FROM golang:1.18-alpine AS build
WORKDIR /

RUN apk --update add bash protoc protobuf-dev && rm -rf /var/cache/apk/*

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go build -o bin/protoc-gen-doc ./cmd/protoc-gen-doc
RUN cp bin/protoc-gen-doc /usr/bin/protoc-gen-doc

VOLUME ["/out"]
VOLUME ["/protos"]

ENTRYPOINT ["./script/entrypoint.sh"]
CMD ["--doc_opt=html,index.html"]
