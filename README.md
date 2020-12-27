# Comandos

    go mod init github.com/galileoluna/greet

    go get -u google.golang.org/grpc     

    go get github.com/golang/protobuf/protoc-gen-go

    protoc  greetpb/greet.proto --go_out=plugins=grpc:.