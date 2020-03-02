# ProtoSmart

This is a simple codec for Go Protocol Buffers that allows you to offload the marshaling and unmarshaling to protobufs to different time in your program.

## Setup
To setup the ProtoSmart codec, just call: `protosmart.OverrideCodec("proto")` and it will override the default protobuf codec with ProtoSmart

## Send Usage
To use ProtoSmart, all you need to do is pre-encode whatever you want to send, turning it into `[]byte`. Once you do this
you can call the server or client `SendMsg(data)` function.

```
b, err := proto.Marshal(MyProtoBufType)
err = server.ServerStream.SendMsg(b)
```

This will auto detect that it has already received bytes and will not try to turn it into bytes again from a protobuf. If you pass it a protobuf, it will convert it to bytes first before sending

## Recv Usage

Rather than pass a protobuf, just pass it a pointer to bytes

```
var b []byte
err := server.ServerStream.RecvMsg(&b)
```

It will not unmarshal from protobuf, it will leave it as is



