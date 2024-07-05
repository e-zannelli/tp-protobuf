Protocol Buffers use a proto definition file to define the schema of the data that will be serialized and deserialized,
and the `protoc` compiler to generate the code that will be used to serialize and deserialize the data.

see https://protobuf.dev/

We need the following messages:
- A `Task` message with the following fields:
   - `title` (string)
   - `done` (bool)
- A `Todo` message which is a list of `Task` messages.

Define the protocol buffer schema in the file [todo.proto](todo.proto).


`make` should generate [todo.pb.go](pb/todo.pb.go) file with the generated Go code.