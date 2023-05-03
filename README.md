# Clean Driven Development

How To Run on Linux
===========================================================================
1. cd cdd && env GOOS=linux GOARCH=amd64 go build -o ../cdd && cd ..
2. cd protoc-gen-cdd && env GOOS=linux GOARCH=amd64 go build -o ../protoc-gen-cdd && cd ..

How To Run Default
===========================================================================
1. cd cdd && go build -o ../cdd && cd ..
2. cd protoc-gen-cdd && go build -o ../protoc-gen-cdd && cd ..
