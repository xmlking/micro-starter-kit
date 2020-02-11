micro --client=grpc call --metadata trans-id=1234 greetersrv Greeter.Hello  '{"name": "John"}'
expect '{"name": "Hello John"}'
