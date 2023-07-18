# protoc-gen-cdd
This is still in Alpha Version.

### What is protoc-gen-cdd
This is implementation of contract driven development. protoc-gen-cdd act as generator of the contracts. Right now it can generate cdd-grst frameworks. https://github.com/krisnasw/go-grst/tree/main/grst

### Requirement
```
- protoc-gen-go v1.25.0 (https://github.com/protocolbuffers/protobuf-go/tree/master/cmd/protoc-gen-go)
- protoc v3.13.0 (https://github.com/protocolbuffers/protobuf/tree/v3.13.0)
- protoc-gen-grpc-gateway v2.2.0 (https://github.com/grpc-ecosystem/grpc-gateway)
- protoc-gen-go-grpc 1.35.0 (https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### How To Use
```
TBD
Will update later with proper examples
```
