# An average Calculator Service

This application provides a calculator service by using a gRPC client streaming API which takes inputs from a user, sends the inputs as message requests(in form of a data stream) to the server  and returns the average of the inputs to the user(client).Through the use of gRPC protocolbuffer, it defines and sends several message request as a stream and returns a respons ethat represent the average of the numbers collected from the commandline.
## Setup

This project requires a `gcc` compiler installed and the `protobuf` code generation tools.

### Install protobuf compiler

Install the `protoc` tool using the instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/). 

## Generating gRPC Code 

To generate the gRPC code, run the following command:
 
```bash
protoc -Iaveragecalculator/proto --go_out=. --go_opt=module=github.com/AdekunleDally/grpc-client-streaming-api --go-grpc_out=. --go-grpc_opt=module=github.com/AdekunleDally/grpc-client-streaming-api averagecalculator/proto/avr_calculator.proto 

