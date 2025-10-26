twirp-todo — proto generation helper

This repository contains protocol buffers for a small Twirp-based Todo service.

Quick generation (PowerShell)

Ensure protoc, protoc-gen-go, and protoc-gen-twirp are installed and on your PATH.

Install generators:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/twitchtv/twirp/protoc-gen-twirp@latest

Generate the proto files:

From the project root, run:

protoc -I=proto --go_out=../proto_generated --go_opt=paths=source_relative --twirp_out=../proto_generated --twirp_opt=paths=source_relative proto/todo.proto
