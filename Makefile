go:
	protoc --proto_path=proto --go_out=gen/go --go_opt=Mpostapp.proto=./postapp --go-grpc_out=gen/go --go-grpc_opt=Mpostapp.proto=./postapp proto/*.proto
	
# Also: protoc --proto_path=proto --go_out=. --go_opt=Mpostapp.proto=../gen/go --go-grpc_out=. --go-grpc_opt=Mpostapp.proto=../gen/go postapp.proto
# --proto_path is to specify which directory to look for the .proto files
# --go_out and --go-grpc_out means where to put the generated go files, with regards to the go_package 
# --go_opt and go-grpc_opt with this pattern: --go_opt=M{PROTOFILE_NAME}={GO_PACKAGE} will overwrite go_package specific in the .proto file
# this way in the --go_out destination (e.g. gen/go) we won't see annoying nested folders like github.com/homayoonalimohammadi/go-post-app/postapp and then
# the generated files inside this.

run:
	go run cmd/postapp/*.go