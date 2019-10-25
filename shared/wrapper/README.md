# Wrappers

> example usage...

```go
service := micro.NewService(
  micro.Name("com.example.srv.foo"),
  micro.Version("v0.1.0"),
  micro.WrapSubscriber(wrapper.NewSubscriberWrapper()),
)

publisher := micro.NewPublisher("topic", service.Client())

service.Init(
  micro.WrapClient(wrapper.NewClientWrapper(service)),
  micro.WrapHandler(wrapper.NewHandlerWrapper(service)),
  micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)),
)
```
