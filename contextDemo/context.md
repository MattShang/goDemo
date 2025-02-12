context 的结构：
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}

根上下文
context.Background()
context.TODO()

派生上下文：
context 包主要作用：管理请求的生命周期
1. 取消操作：用于终止 goroutine 或其他操作
`context.WithCancel` 创建可取消的上下文，调用 cancel 时，所有监听该上下文的操作都会收到取消信号。
ctx.Err() 返回 context.Canceled
2. 超时控制：设置操作的超时时间，超时自动取消
`context.WithTimeout` 可以设置超时时间，超时后自动取消
ctx.Err() 返回 context.DeadlineExceeded
3. 传递数据：在请求链中安全的传递键值对数据
`context.WithValue` 可以在上下文中传递键值对数据