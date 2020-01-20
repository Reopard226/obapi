# To generate code:
Run the following:

```shell script
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/iam/service.proto
protoc -I /usr/local/include \
        -I . \
        --gotag_out=xxx="bson+\"-\"":. ./rpc/iam/service.proto
```