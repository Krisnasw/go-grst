```
#protobuf protoc-gen-go V2
#in protoc-gen-cdd/ext/cddapis folder
protoc \
    --proto_path=$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/ \
    --go_out $GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis --go_opt paths=source_relative \
    --go-grpc_out $GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis --go-grpc_opt paths=source_relative \
    cdd/api/cddext.proto  
```
