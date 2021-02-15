#!/bin/bash

protoc --proto_path=greet --go_out=greet --go-grpc_out=greet --go_opt=paths=source_relative greet/greetpb/greet.proto