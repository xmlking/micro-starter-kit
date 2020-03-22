micro --client=grpc call --metadata trans-id=1234 mkit.service.greeter Greeter.Hello  '{"name": "John"}'
expect '{"name": "Hello John"}'
