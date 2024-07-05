Use grpc to implement a server that will do the operations to the dbPath file, the todo command will now be a client of this server.

You will eventually need go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

The first step is to define the service in the [proto file](todo.proto) and generate the go code with make.
Update the code of the [server](cmd/server/main.go) to implement the service defined in the proto file.
Then update the code of the [todo command](cmd/todo/main.go) to use the grpc client to communicate with the server.
