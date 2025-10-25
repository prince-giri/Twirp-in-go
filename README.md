# twirp-todo — proto generation helper

This repository contains protocol buffers for a small Twirp-based Todo service.

Quick generation (PowerShell):

1. Ensure `protoc`, `protoc-gen-go`, and `protoc-gen-twirp` are installed and on your PATH.
   - Install generators with:
     ```powershell
     go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
     go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
     ```

2. Run the helper script (from project root):
   ```powershell
   .\scripts\generate-proto.ps1
   ```

What I fixed
- The original error happens when `protoc` can't find the .proto file relative to any of its include (`-I`) paths. You need to either run `protoc` from the `proto/` directory or pass `-I=proto` and give the path `proto\todo.proto`.
- I ran `protoc` with `-I=proto` and `--*_opt=paths=source_relative` to produce the generated files in `proto/proto_generated`.

If you want, wire this script into your build/CI or adapt output directories to your preferred layout.
"# Twirp-in-go" 
