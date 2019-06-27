add-submodule:
	git submodule add git@github.com:mf-sakura/bh_proto.git proto

proto-gen-go:
	protoc --proto_path=proto/hotel/v1/ --go_out=plugins=grpc:app/proto proto/hotel/v1/hotel.proto

update-submodule:
	git submodule foreach 'git fetch;git checkout master; git pull'

db-test:
	go test github.com/mf-sakura/bh_hotel/app/db

build:
	go build -o bh_hotel github.com/mf-sakura/bh_hotel/app/cmd