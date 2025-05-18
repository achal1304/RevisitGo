import grpc
import service_pb2
import service_pb2_grpc

def send_message(stub):
    response = stub.SendMessage(service_pb2.MessageRequest(message="World"))
    print("Unary response:", response.response)

def stream_messages(stub):
    print("Server Streaming...")
    request = service_pb2.StreamRequest(prefix="Message")
    for response in stub.StreamMessages(request):
        print(response.message)

def stream_client_messages(stub):
    print("Client Streaming...")
    messages = ["Hello", "from", "Python", "client"]
    def message_generator():
        for msg in messages:
            yield service_pb2.MessageRequest(message=msg)
    response = stub.StreamClientMessages(message_generator())
    print("Client Stream response:", response.response)

def bidirectional_stream(stub):
    print("Bidirectional Streaming...")
    messages = ["Hello", "this", "is", "bidirectional", "streaming"]
    def message_generator():
        for msg in messages:
            yield service_pb2.MessageRequest(message=msg)
    for response in stub.BidirectionalStream(message_generator()):
        print(response.response)

def run():
    with grpc.insecure_channel('localhost:9000') as channel:
        stub = service_pb2_grpc.MessagingServiceStub(channel)
        send_message(stub)
        stream_messages(stub)
        stream_client_messages(stub)
        bidirectional_stream(stub)

if __name__ == '__main__':
    run()
