# statsviz
## 使用

下载最新版本:

    go get github.com/hellobchain/statsviz@latest


样例如下：

```go
mux := http.NewServeMux()
statsviz.Register(mux)
```

或者 这样方式注册 `http.DefaultServeMux`:

```go
statsviz.RegisterDefault()
```

默认URL： `/debug/statsviz/`.

```go
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

然后浏览器打开：http://localhost:6060/debug/statsviz/.