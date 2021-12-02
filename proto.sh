function gen_go_grpc(){
  echo   "Generating go grpc"
  protoc \
     --proto_path=proto \
     --go_out=proto \
     --go-grpc_out=paths=source_relative:proto \
     --go_opt=paths=source_relative \
      proto/note.proto
}

gen_go_grpc
