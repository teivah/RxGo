package disposable

type Disposable interface {
	Dispose()
	IsDisposed() bool
}
