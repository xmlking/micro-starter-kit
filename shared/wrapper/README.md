# Wrappers

> example usage...

```go
srv := micro.NewService(
  micro.Name("com.example.srv.foo"),
  micro.Version("v0.1.0"),
  micro.WrapSubscriber(wrapper.NewSubscriberWrapper()),
)

srv.Init(
  micro.WrapClient(wrapper.NewClientWrapper(srv)),
  micro.WrapHandler(wrapper.NewHandlerWrapper(srv)),
)
```
