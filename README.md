# grpcgen-go
Generating Client Code and Server Boilerplate using gRPC

[![Build Status](https://travis-ci.org/marcusljx/grpcgen-go.svg?branch=master)](https://travis-ci.org/marcusljx/grpcgen-go)
<!-- [![codecov](https://codecov.io/gh/marcusljx/grpcgen-go/branch/master/graph/badge.svg)](https://codecov.io/gh/marcusljx/grpcgen-go) -->
[![Coverage Status](https://coveralls.io/repos/github/marcusljx/grpcgen-go/badge.svg?branch=master)](https://coveralls.io/github/marcusljx/grpcgen-go?branch=master)
[![Issue Count](https://codeclimate.com/github/marcusljx/grpcgen-go/badges/issue_count.svg)](https://codeclimate.com/github/marcusljx/grpcgen-go)

This project is still in pre-alpha stage. Do not use in production!

## Goal
The goal is to be able to create a simple working microservice using the following steps:

1. Create a protobuf definition file (`.proto`)
2. Invoke the protobuf compiler to generate golang interfaces for client and server
3. Invoke `grpcgen` to generate Client code, and Server boilerplate code.
4. Fill in the `logic` within all the Server boilerplate.
5. Run `go run start_server.go`.

## Future work
- [ ] unit test boilerplate for server methods (using httptest)
- [ ] mocking boilerplate (for use either as-is, or with additional logic)
