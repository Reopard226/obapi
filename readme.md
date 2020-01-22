# To generate code:
Run the following:

```shell script
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/iam/service.proto
protoc -I /usr/local/include \
        -I . \
        --gotag_out=xxx="bson+\"-\" firestore+\"-\"":. ./rpc/iam/service.proto
```

[![codecov](https://codecov.io/gl/oceanbolt/iamserver/branch/master/graph/badge.svg?token=i8vFzG5tBo)](https://codecov.io/gl/oceanbolt/iamserver)

# Api

Oceanbolt Data Api

## JWT 

API Parses JWT keys from Auth0 - fetches keyset if it is not buffered locally.

API also signs own JWT tokens, and parses them using a private and public key.

## Databases

MongoDB and Postgres
Most data is being migrated to MongoDB

