# Oceanbolt Data API monorepo

## Introduction

This repo is a monorepo, containing all the microservices used to deploy the Oceanbolt Data API.

### Commit Instructions

This repo used a trunk-based development - meaning all commits should be done to master branch.
It is acceptable to have very short lived feature branches, but ideally everyone should commit their work to master branch at least once per day.

Git workflow:
- Work on **SHORT LIVED** feature branch
- Work locally
- Commit regularly locally
- Test locally
- If tests pass - Push to origin to run the tests on the feature branch
- Squash commits (force push to remote branch if you have pushed to origin already)
- Merge into master (rebase + fast forward merge) - This should be done in Gitlab


### Microservices in this repo:
- iamserver: IAM and apikey management (internal only)
- echoapi: main accesspoint for customers (external)

Layout:

```
root
├── cmd
│   ├── echoapi
│   │   └── main.go                         //entrypoint for echoapi
|   |
│   └── iamserver
│       └── main.go                         //entrypoint for iamserver

├── internal
│   ├── echoapi
│       └── ...                             //internal echoapi packages
|
│   └── iam
│       └── ...                            //internal iamserver packages
|
├── rpc 
│   └── iam                                 //rpc definitions for iamserver
│       ├── service.pb.go
│       ├── service.proto
│       └── service.twirp.go
|
├── tagger
│   └── tagger.proto                        //proto build files
|
└── test.sh                                 //test script

```

## iamserver
The **iamserver** microservice manages the apikey validation and creation. It is TWIRP based.
It is based on a protobuf file (service.proto), and skeleton code needs to be generated with the following command:
 
##### To generate code:
Run the following:

```shell script
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/iam/service.proto
protoc -I /usr/local/include \
        -I . \
        --gotag_out=xxx="bson+\"-\" firestore+\"-\"":. ./rpc/iam/service.proto
```

[![codecov](https://codecov.io/gl/oceanbolt/iamserver/branch/master/graph/badge.svg?token=i8vFzG5tBo)](https://codecov.io/gl/oceanbolt/iamserver)

## echoapi

The main entrypoint for customers to access the Oceanbolt Data Api. This is echo based.


## JWT 

API Parses JWT keys from Auth0 - fetches keyset if it is not buffered locally.

API also signs own JWT tokens, and parses them using a private and public key.

## Databases

MongoDB and Postgres
Most data is being migrated to MongoDB


# Credentials

All credentials are passed out using ENVKEY variables. This works by setting an ENVKEY environment variable, and then the config is automatically fetched.

