# XCTX

这里的类型用于在golang中使用类似 python 的
```python
with resouce:
    ...
```
或者csharp中的
```csharp
using (resource):
    ...
```

xctx (xContext) 用于控制资源的生命周期和控制权

## ReleaseLocker

```go
var mut syncx.Mutex

xctx.NewLocker(&mut).Run(func() {
    ...
})
```

## ReleaseCloser

```go
var closer io.Closer = ...

xctx.NewCloser(&closer).Run(func() {
    ...
})
```