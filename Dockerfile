FROM golang:alpine3.19 

ENV CGO_ENABLED=1

WORKDIR /app

COPY . .
 
# EXPOSE 50051

RUN apk update && apk add --no-cache build-base && \
apk add --no-cache gcc && \
apk add --no-cache wget && go mod tidy

ENTRYPOINT ["go",  "run", "cmd/grpcServer/main.go"]