# Wrappers

> example usage...

```go
service := micro.NewService(
  micro.Name("com.example.service.foo"),
  micro.Version("v0.1.0"),
  micro.WrapSubscriber(wrapper.NewSubscriberWrapper()),
)

publisher := micro.NewEvent("topic", service.Client())

service.Init(
  micro.WrapClient(wrapper.NewClientWrapper(service)),
  micro.WrapHandler(wrapper.NewHandlerWrapper(service)),
  micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)),
)
```
