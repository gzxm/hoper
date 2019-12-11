package xyz.hoper.hello.client;

import io.grpc.ManagedChannel;
import lombok.extern.log4j.Log4j2;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import xyz.hoper.protobuf.GreeterGrpc;
import xyz.hoper.protobuf.HelloReply;
import xyz.hoper.protobuf.HelloRequest;

import javax.annotation.PostConstruct;

@Component
@Log4j2
public class HelloWorldClient {
    @Value("${grpc.client.host}")
    private String host;
    @Value("${grpc.client.port}")
    private Integer port;

    @Autowired
    private GrpcClientMananer grpcClientMananer;

    @PostConstruct
    public void init() {
        call();
    }

    public void call(){
        ManagedChannel channel = grpcClientMananer.getChannel(host,port);
        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(channel);
        HelloRequest request = HelloRequest.newBuilder().setName("world").build();
        HelloReply helloReply = stub.sayHello(request);
        log.info("time: "+ helloReply.getTime());
    }
}