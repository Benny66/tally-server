protoc-gen-go := $(shell which protoc-gen-go)
protoc-gen-go-grpc := $(shell which protoc-gen-go-grpc)
protoc := $(shell which protoc)

proto_dirs := $(shell find ./proto -type d)
proto_files := $(foreach dir,$(proto_dirs),$(wildcard $(dir)/*.proto))

proto_go_files := $(patsubst %.proto,%.pb.go,$(proto_files))
proto_grpc_go_files := $(patsubst %.proto,%_grpc.pb.go,$(proto_files))

all: $(proto_go_files) $(proto_grpc_go_files)

%_grpc.pb.go: %.proto
	$(protoc) -I ./proto \
	--go_out=plugins=grpc:. \
	--go_opt=module=github.com/Benny66/tally-server \
	--grpc_out=. \
	--grpc_opt=module=github.com/Benny66/tally-server \
	$<

%.pb.go: %.proto
	$(protoc) -I ./proto \
	--go_out=. \
	--go_opt=module=github.com/Benny66/tally-server \
	$<
