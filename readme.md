# To generate code:
Run the following:

```shell script
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/iam/service.proto
protoc -I /usr/local/include \
        -I . \
        --gotag_out=xxx="bson+\"-\"":. ./rpc/iam/service.proto
```

[![codecov](https://codecov.io/gl/oceanbolt/iamserver/branch/master/graph/badge.svg?token=i8vFzG5tBo)](https://codecov.io/gl/oceanbolt/iamserver)