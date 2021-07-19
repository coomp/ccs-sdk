#!/bin/bash
protoc -I. --go_out=plugins=grpc:. ./grpc/message.proto